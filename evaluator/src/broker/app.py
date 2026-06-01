from faststream import FastStream
from faststream.rabbit import RabbitBroker, RabbitQueue
from loguru import logger

from src.model import MusiqInference
from src.models import ProcessingQueueItem
from src.settings import settings

broker = RabbitBroker(settings.rabbitmq_url)
app = FastStream(broker)
model = MusiqInference()


@broker.subscriber(RabbitQueue(settings.rabbitmq_processing_queue, durable=True))
async def handle(msg: ProcessingQueueItem):
    logger.debug(f"Got a message from processing queue: {msg}!")

    score = model.evaluate(msg.image_path)
    logger.success(f"Image: {msg.image_path}\n{score=}")

    await broker.publish({"id": msg.id, "score": score}, settings.rabbitmq_done_queue)
