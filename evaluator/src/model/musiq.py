import os
import pathlib
from dataclasses import dataclass
from typing import Any, Dict, Iterable, List, Optional, Union

import numpy as np
import pyiqa
import torch
from PIL import Image

PathLike = Union[str, os.PathLike]


@dataclass
class MusiqResult:
    score: float
    model_name: str
    source: Optional[str] = None
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
    ) -> None:
        self.model_name = model_name.lower()

        if device is None:
            device = "cuda" if torch.cuda.is_available() else "cpu"
        self._device = torch.device(device)

        pyiqa_name = self._resolve_pyiqa_name(self.model_name)
        self.metric = pyiqa.create_metric(pyiqa_name, device=self._device)

        if warmup:
            self.warmup()

    def evaluate(
        self,
        image: Union[PathLike, bytes, bytearray, Image.Image, np.ndarray],
        source: str | None = None,
    ) -> float:
        image_tensor = self._to_tensor(image)
        score = self.metric(image_tensor).item()

        return score

    def evaluate_file(
        self,
        path: PathLike,
        return_raw: bool = False,
    ) -> Union[float, MusiqResult]:
        return self.evaluate(
            image=path,
            source=str(path),
            return_raw=return_raw,
        )

    def evaluate_many(
        self,
        images: Iterable[Union[PathLike, bytes, bytearray, Image.Image, np.ndarray]],
        return_raw: bool = False,
    ) -> Union[List[float], List[MusiqResult]]:
        results = []
        for image in images:
            source = str(image) if isinstance(image, (str, os.PathLike)) else None
            result = self.evaluate(image=image, source=source, return_raw=return_raw)
            results.append(result)
        return results

    def evaluate_directory(
        self,
        directory: PathLike,
        extensions: Optional[Iterable[str]] = None,
        recursive: bool = False,
        return_raw: bool = True,
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

        results = []
        for path in paths:
            result = self.evaluate_file(path, return_raw=return_raw)
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
        _ = self.metric(tensor)

    def _to_tensor(
        self,
        image: Union[PathLike, bytes, bytearray, Image.Image, np.ndarray],
    ) -> torch.Tensor:
        if isinstance(image, (str, os.PathLike)):
            return self._load_image_tensor(str(image))

        if isinstance(image, (bytes, bytearray)):
            import io

            pil = Image.open(io.BytesIO(image)).convert("RGB")
            return self._pil_to_tensor(pil)

        if isinstance(image, Image.Image):
            if image.mode != "RGB":
                image = image.convert("RGB")
            return self._pil_to_tensor(image)

        if isinstance(image, np.ndarray):
            pil = Image.fromarray(image.astype(np.uint8))
            if pil.mode != "RGB":
                pil = pil.convert("RGB")
            return self._pil_to_tensor(pil)

    def _load_image_tensor(self, path: str) -> torch.Tensor:
        img = Image.open(path).convert("RGB")
        return self._pil_to_tensor(img)

    def _pil_to_tensor(self, image: Image.Image) -> torch.Tensor:
        from torchvision import transforms

        transform = transforms.ToTensor()
        tensor = transform(image).unsqueeze(0).to(self._device)
        return tensor

    def _resolve_pyiqa_name(self, model_name: str) -> str:
        model_name = model_name.lower()
        if model_name not in self.MODEL_VARIANTS:
            raise ValueError(
                f"Unknown MUSIQ model: {model_name!r}. "
                f"Available: {sorted(self.MODEL_VARIANTS)}"
            )
        return self.MODEL_VARIANTS[model_name]
