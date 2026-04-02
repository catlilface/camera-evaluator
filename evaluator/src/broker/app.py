from faststream import FastStream
from faststream.rabbit import RabbitBroker
from loguru import logger

from src.models import ProcessingQueueItem
from src.settings import settings

broker = RabbitBroker()
app = FastStream(broker)


@broker.subscriber(settings.rabbitmq_processing_queue)
async def handle(message: ProcessingQueueItem):
    logger.success("Got a message from processing queue: {msg}!")

    await broker.publish(message, settings.rabbitmq_done_queue)
