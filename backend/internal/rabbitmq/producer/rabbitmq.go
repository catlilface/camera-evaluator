package producer

import (
	"photo-upload-service/internal/config"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	conn           *amqp.Connection
	channel        *amqp.Channel
	queueName      string
	publishTimeout time.Duration
}

func NewPublisher(cfg *config.Cfg) (*Publisher, error) {
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
		cfg.Queue.QueueName,
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

	return &Publisher{
		conn:           conn,
		channel:        channel,
		queueName:      cfg.Queue.QueueName,
		publishTimeout: cfg.Queue.PublishTimeout,
	}, nil
}

func (p *Publisher) Close() error {
	var err error
	if p.channel != nil {
		err = p.channel.Close()
	}
	if p.conn != nil {
		if closeErr := p.conn.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
	}
	return err
}
