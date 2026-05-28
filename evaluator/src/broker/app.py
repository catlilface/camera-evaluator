from faststream import FastStream
from faststream.rabbit import RabbitBroker, RabbitQueue
from loguru import logger

from src.models import ProcessingQueueItem
from src.settings import settings

broker = RabbitBroker(settings.rabbitmq_url)
app = FastStream(broker)


@broker.subscriber(RabbitQueue(settings.rabbitmq_processing_queue, durable=True))
async def handle(message: ProcessingQueueItem):
    logger.success(f"Got a message from processing queue: {message}!")

    await broker.publish(message, RabbitQueue(settings.rabbitmq_done_queue, durable=True))
