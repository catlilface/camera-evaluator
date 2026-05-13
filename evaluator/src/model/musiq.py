import os
from io import BytesIO

import tensorflow as tf
from PIL import Image

PP_CONFIG = {
    "patch_size": 32,
    "patch_stride": 32,
    "hse_grid_size": 10,
    "longer_side_lengths": [224, 384],  # default 3-scale MUSIQ
    "max_seq_len_from_original_res": -1,  # use all original-resolution patches
}


class MUSIQ:
    model_location = os.path.abspath(".model")

    def __init__(self):
        self.model = tf.saved_model.load(self.model_location)
        self.predict_fn = self.model.signatures["serving_default"]

    def _image_to_rgb_jpeg_bytes(self, image_path):
        img = Image.open(image_path).convert("RGB")
        buf = BytesIO()
        img.save(buf, format="JPEG", quality=95)
        return buf.getvalue()

    def predict_quality(self, image_path):
        image_bytes = self._image_to_rgb_jpeg_bytes(image_path)
        result = self.predict_fn(tf.constant(image_bytes))

        score = list(result.values())[0].numpy()
        return score
