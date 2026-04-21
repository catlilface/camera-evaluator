package consumer

import (
	"context"
	"encoding/json"
	"log"
	"photo-upload-service/internal/config"
	"photo-upload-service/internal/models"
	"photo-upload-service/pkg/websocket"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	conn      *amqp.Connection
	channel   *amqp.Channel
	cancel    context.CancelFunc
	queueName string
	wsManager *websocket.Manager
}

func NewConsumer(ctx context.Context, cfg *config.Cfg, wsManager *websocket.Manager) (*Consumer, error) {
	conn, err := amqp.Dial(cfg.Queue.URL)
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	_, err = channel.QueueDeclare(
		cfg.Queue.ResponseQueueName,
		true,  // durable
		false, // autoDelete
		false, // exclusive
		false, // noWait
		nil,   // arguments
	)
	if err != nil {
		channel.Close()
		conn.Close()
		return nil, err
	}

	consumerCtx, cancel := context.WithCancel(ctx)

	consumer := &Consumer{
		conn:      conn,
		channel:   channel,
		cancel:    cancel,
		queueName: cfg.Queue.ResponseQueueName,
		wsManager: wsManager,
	}

	err = consumer.Start(consumerCtx)
	if err != nil {
		channel.Close()
		conn.Close()
		return nil, err
	}

	return consumer, nil
}

func (c *Consumer) Start(ctx context.Context) error {
	msgs, err := c.channel.Consume(
		c.queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case msg := <-msgs:
				c.handleMessage(msg.Body)
			case <-ctx.Done():
				c.channel.Close()
				c.conn.Close()
				return
			}
		}
	}()

	return nil
}

func (c *Consumer) Stop() error {
	c.cancel()
	return nil
}

func (c *Consumer) handleMessage(body []byte) {
	var message models.EvaluationResponseMessage

	log.Println("[Consumer]" + string(body))

	if err := json.Unmarshal(body, &message); err != nil {
		log.Println("[Consumer] failed to parse message:", err)
		return
	}

	if message.ImageID == "" {
		log.Println("[Consumer] empty channel_id")
		return
	}

	c.wsManager.Send(message.ImageID, message)
	c.wsManager.CloseChannel(message.ImageID)
}
