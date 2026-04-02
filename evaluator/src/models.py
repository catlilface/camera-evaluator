from pydantic import BaseModel


class ProcessingQueueItem(BaseModel):
    """Model for item from processing queue"""

    id: str
    """UUID for the given request"""
    image_path: str
    """Path to sample image in shared directory"""


class DoneQueueMessage(BaseModel):
    id: str
    """UUID for the given request"""
    # TODO: return a structured score values model
    score: float
    """Generic computed score for the given camera"""
