import os
import pathlib
from contextlib import contextmanager
from dataclasses import dataclass
from typing import Any, Dict, Iterable, List, Optional, Tuple, Union
from uuid import uuid4

import numpy as np
import pyiqa
import torch
from PIL import Image, ImageEnhance, ImageFilter, ImageOps

PathLike = Union[str, os.PathLike]


@dataclass
class MusiqResult:
    score: float
    model_name: str
    source: Optional[str] = None
    attention_image_path: Optional[str] = None
    raw_output: Optional[Dict[str, Any]] = None


class MusiqInference:
    MODEL_VARIANTS: Dict[str, str] = {
        "spaq": "musiq-spaq",
        "koniq": "musiq",
        "koniq-10k": "musiq",
        "paq2piq": "musiq-paq2piq",
        "ava": "musiq-ava",
    }

    def __init__(
        self,
        model_name: str = "koniq",
        device: str | None = None,
        warmup: bool = True,
        attn_img_dir: PathLike = "musiq_attention_images",
    ) -> None:
        self.model_name = model_name.lower()

        if device is None:
            device = "cuda" if torch.cuda.is_available() else "cpu"
        self._device = torch.device(device)

        self.attn_img_dir = pathlib.Path(attn_img_dir)
        self.attn_img_dir.mkdir(parents=True, exist_ok=True)

        pyiqa_name = self._resolve_pyiqa_name(self.model_name)
        self.metric = pyiqa.create_metric(pyiqa_name, device=self._device)
        self.metric.eval()

        if warmup:
            self.warmup()

    def evaluate(
        self,
        image: Union[PathLike, bytes, bytearray, Image.Image, np.ndarray],
        source: str | None = None,
    ) -> MusiqResult:
        image_tensor, pil_image = self._to_tensor_and_pil(image)

        with torch.inference_mode():
            score_tensor = self._extract_score_tensor(self.metric(image_tensor))
            score = float(score_tensor.reshape(-1)[0].item())

        attention_image_path = None
        attention_image_path = self._save_attention_image(
            image_tensor=image_tensor,
            pil_image=pil_image,
        )

        raw_output = None

        return MusiqResult(
            score=score,
            model_name=self.model_name,
            source=source,
            attention_image_path=attention_image_path,
            raw_output=raw_output,
        )

    def evaluate_file(
        self,
        path: PathLike,
        return_raw: bool = False,
        save_attention: bool = False,
    ) -> Union[float, MusiqResult]:
        return self.evaluate(
            image=path,
            source=str(path),
        )

    def evaluate_many(
        self,
        images: Iterable[Union[PathLike, bytes, bytearray, Image.Image, np.ndarray]],
        return_raw: bool = False,
        save_attention: bool = False,
    ) -> Union[List[float], List[MusiqResult]]:
        results = []

        for image in images:
            source = str(image) if isinstance(image, (str, os.PathLike)) else None
            result = self.evaluate(
                image=image,
                source=source,
            )
            results.append(result)

        return results

    def evaluate_directory(
        self,
        directory: PathLike,
        extensions: Optional[Iterable[str]] = None,
        recursive: bool = False,
        return_raw: bool = True,
        save_attention: bool = True,
    ) -> List[MusiqResult]:
        if extensions is None:
            extensions = {".jpg", ".jpeg", ".png", ".webp", ".bmp"}

        extensions = {ext.lower() for ext in extensions}
        directory = pathlib.Path(directory)
        pattern = "**/*" if recursive else "*"

        paths = [
            path
            for path in directory.glob(pattern)
            if path.is_file() and path.suffix.lower() in extensions
        ]

        results: List[MusiqResult] = []

        for path in paths:
            result = self.evaluate_file(
                path=path,
                return_raw=return_raw,
                save_attention=save_attention,
            )

            if isinstance(result, MusiqResult):
                results.append(result)
            else:
                results.append(
                    MusiqResult(
                        score=float(result),
                        model_name=self.model_name,
                        source=str(path),
                    )
                )

        return results

    def warmup(self) -> None:
        img = Image.new("RGB", (64, 64), color=(127, 127, 127))
        tensor = self._pil_to_tensor(img)

        with torch.inference_mode():
            _ = self.metric(tensor)

    def _save_attention_image(
        self,
        image_tensor: torch.Tensor,
        pil_image: Image.Image,
    ) -> str:
        attention = self._compute_input_attention(image_tensor)
        attention_image = self._render_attention_on_image(
            pil_image=pil_image,
            attention=attention,
        )

        output_path = self.attn_img_dir / f"{uuid4()}.png"
        attention_image.save(output_path)

        return str(output_path)

    def _compute_input_attention(self, image_tensor: torch.Tensor) -> np.ndarray:
        tensor = image_tensor.detach().clone().requires_grad_(True)

        self.metric.zero_grad(set_to_none=True)

        with self._temporarily_disable_parameter_grads():
            with torch.enable_grad():
                score_tensor = self._forward_metric_for_grad(tensor)
                score = score_tensor.reshape(-1).mean()
                score.backward()

        if tensor.grad is None:
            raise RuntimeError(
                "Could not compute attention map because no input gradient was produced."
            )

        grad = tensor.grad.detach().abs()

        # Shape: $1 \times C \times H \times W$ -> $H \times W$
        attention = grad.mean(dim=1).squeeze(0)
        attention = attention.float().cpu().numpy()

        attention = self._normalize_attention(attention)

        return attention

    def _forward_metric_for_grad(self, tensor: torch.Tensor) -> torch.Tensor:
        """
        First tries the normal PyIQA metric forward pass.

        Some PyIQA wrappers may internally disable gradients for inference.
        If that happens, this falls back to calling `metric.net` directly.
        """
        output = self.metric(tensor)
        score_tensor = self._extract_score_tensor(output)

        if score_tensor.requires_grad:
            return score_tensor

        if hasattr(self.metric, "net"):
            output = self.metric.net(tensor)
            score_tensor = self._extract_score_tensor(output)

            if score_tensor.requires_grad:
                return score_tensor

        raise RuntimeError(
            "Could not get a differentiable MUSIQ score from PyIQA. "
            "Your PyIQA version may wrap MUSIQ inference in `torch.no_grad()` "
            "or may not expose a differentiable internal network."
        )

    def _render_attention_on_image(
        self,
        pil_image: Image.Image,
        attention: np.ndarray,
    ) -> Image.Image:
        pil_image = pil_image.convert("RGB")

        resampling = getattr(Image, "Resampling", Image).BILINEAR

        mask = Image.fromarray((attention * 255).astype(np.uint8), mode="L")
        mask = mask.resize(pil_image.size, resampling)

        blur_radius = max(1, int(min(pil_image.size) * 0.02))
        mask = mask.filter(ImageFilter.GaussianBlur(radius=blur_radius))
        mask = ImageEnhance.Contrast(mask).enhance(1.8)

        # Dark image where attention is low, keep original where attention is high.
        dark_image = ImageEnhance.Brightness(pil_image).enhance(0.25)
        attention_image = Image.composite(pil_image, dark_image, mask)

        return attention_image

    def _normalize_attention(self, attention: np.ndarray) -> np.ndarray:
        attention = attention.astype(np.float32)
        attention = attention - float(attention.min())

        max_value = float(attention.max())
        if max_value > 1e-8:
            attention = attention / max_value

        # Slight gamma correction to make visible regions clearer.
        attention = np.sqrt(attention)

        return np.clip(attention, 0.0, 1.0)

    @contextmanager
    def _temporarily_disable_parameter_grads(self):
        params = list(self.metric.parameters())
        old_requires_grad = [param.requires_grad for param in params]

        try:
            for param in params:
                param.requires_grad_(False)
            yield
        finally:
            for param, requires_grad in zip(params, old_requires_grad):
                param.requires_grad_(requires_grad)

    def _to_tensor(
        self,
        image: Union[PathLike, bytes, bytearray, Image.Image, np.ndarray],
    ) -> torch.Tensor:
        tensor, _ = self._to_tensor_and_pil(image)
        return tensor

    def _to_tensor_and_pil(
        self,
        image: Union[PathLike, bytes, bytearray, Image.Image, np.ndarray],
    ) -> Tuple[torch.Tensor, Image.Image]:
        if isinstance(image, (str, os.PathLike)):
            pil = self._load_image(str(image))
            return self._pil_to_tensor(pil), pil

        if isinstance(image, (bytes, bytearray)):
            import io

            pil = Image.open(io.BytesIO(image))
            pil = ImageOps.exif_transpose(pil).convert("RGB")
            return self._pil_to_tensor(pil), pil

        if isinstance(image, Image.Image):
            pil = ImageOps.exif_transpose(image).convert("RGB")
            return self._pil_to_tensor(pil), pil

        if isinstance(image, np.ndarray):
            pil = self._ndarray_to_pil(image)
            return self._pil_to_tensor(pil), pil

        raise TypeError(f"Unsupported image type: {type(image)!r}")

    def _load_image(self, path: str) -> Image.Image:
        img = Image.open(path)
        img = ImageOps.exif_transpose(img).convert("RGB")
        return img

    def _ndarray_to_pil(self, array: np.ndarray) -> Image.Image:
        array = np.asarray(array)

        if array.dtype != np.uint8:
            if np.issubdtype(array.dtype, np.floating):
                if float(np.nanmax(array)) <= 1.0:
                    array = array * 255.0

            array = np.nan_to_num(array)
            array = np.clip(array, 0, 255).astype(np.uint8)

        if array.ndim == 2:
            return Image.fromarray(array, mode="L").convert("RGB")

        if array.ndim == 3:
            if array.shape[2] == 1:
                return Image.fromarray(array[:, :, 0], mode="L").convert("RGB")

            if array.shape[2] == 3:
                return Image.fromarray(array, mode="RGB")

            if array.shape[2] == 4:
                return Image.fromarray(array, mode="RGBA").convert("RGB")

        raise ValueError(f"Unsupported ndarray shape for image: {array.shape}")

    def _pil_to_tensor(self, image: Image.Image) -> torch.Tensor:
        image = image.convert("RGB")
        array = np.asarray(image).astype(np.float32) / 255.0

        tensor = torch.from_numpy(array)
        tensor = tensor.permute(2, 0, 1).unsqueeze(0)
        tensor = tensor.to(self._device)

        return tensor

    def _extract_score_tensor(self, output: Any) -> torch.Tensor:
        if torch.is_tensor(output):
            return output

        if isinstance(output, dict):
            preferred_keys = (
                "score",
                "scores",
                "pred",
                "prediction",
                "output",
                "quality",
            )

            for key in preferred_keys:
                value = output.get(key)
                if torch.is_tensor(value):
                    return value

            for value in output.values():
                if torch.is_tensor(value):
                    return value

        if isinstance(output, (list, tuple)):
            for value in output:
                if torch.is_tensor(value):
                    return value

        raise TypeError(
            f"Could not extract score tensor from output type: {type(output)!r}"
        )

    def _resolve_pyiqa_name(self, model_name: str) -> str:
        model_name = model_name.lower()

        if model_name not in self.MODEL_VARIANTS:
            raise ValueError(
                f"Unknown MUSIQ model: {model_name!r}. "
                f"Available: {sorted(self.MODEL_VARIANTS)}"
            )

        return self.MODEL_VARIANTS[model_name]
