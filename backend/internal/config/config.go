package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/sethvargo/go-envconfig"
)

type Cfg struct {
	Service Service
	Queue   RabbitMQ
}

type Service struct {
	Host            string        `env:"BACKEND_HOST,default=0.0.0.0"`
	MainPort        string        `env:"BACKEND_EXTERNAL_PORT,default=8080"`
	ReadTimeout     time.Duration `env:"READ_TIMEOUT_IN_SEC,default=10s"`
	WriteTimeout    time.Duration `env:"WRITE_TIMEOUT_IN_SEC,default=10s"`
	IdleTimeout     time.Duration `env:"IDLE_TIMEOUT_IN_SEC,default=15s"`
	ShutdownTimeout time.Duration `env:"SHUTDOWN_TIMEOUT_IN_SEC,default=3s"`
}

type RabbitMQ struct {
	Host              string        `env:"RABBITMQ_HOST,default=localhost"`
	PortUI            string        `env:"RABBITMQ_PORT_UI,default=15672"`
	PortAMQP          string        `env:"RABBITMQ_PORT_AMQP,default=5672"`
	User              string        `env:"RABBITMQ_USER,default=guest"`
	Password          string        `env:"RABBITMQ_PASSWORD,default=guest"`
	QueueName         string        `env:"RABBITMQ_PHOTO_QUEUE_NAME,default=photos_queue"`
	ResponseQueueName string        `env:"RABBITMQ_RESPONSE_QUEUE_NAME,default=response_queue"`
	PublishTimeout    time.Duration `env:"RABBITMQ_PUBLISH_TIMEOUT_IS_SEC,default=30s"`
	URL               string
}

func (cfg *Cfg) MustLoad() {
	ctx := context.Background()

	if err := envconfig.Process(ctx, cfg); err != nil {
		log.Fatal("Failed to load configuration: ", err)
	}

	cfg.Queue.URL = fmt.Sprintf("amqp://%s:%s@%s:%s/",
		cfg.Queue.User,
		cfg.Queue.Password,
		cfg.Queue.Host,
		cfg.Queue.PortAMQP,
	)

	log.Println("Configuration loaded successfully")
}
