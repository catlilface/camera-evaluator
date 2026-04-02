package config

import (
	"context"
	"fmt"
	"github.com/sethvargo/go-envconfig"
	"log"
)

type Cfg struct {
	Service  Service
	Queue    RabbitMQ
	Postgres Postgres
}

type Service struct {
	Host     string `env:"BACKEND_HOST, default=0.0.0.0"`
	MainPort string `env:"BACKEND_EXTERNAL_PORT, default=8080"`
}

type RabbitMQ struct {
	Host      string `env:"RABBITMQ_HOST, default=localhost"`
	PortUI    string `env:"RABBITMQ_PORT_UI, default=15672"`
	PortAMQP  string `env:"RABBITMQ_PORT_AMQP, default=5672"`
	User      string `env:"RABBITMQ_USER, default=guest"`
	Password  string `env:"RABBITMQ_PASSWORD, default=guest"`
	QueueName string `env:"RABBITMQ_PHOTO_QUEUE_NAME, default=photos_queue"`
	URL       string
}

type Postgres struct {
	Host     string `env:"POSTGRES_HOST, default=localhost"`
	Port     string `env:"POSTGRES_PORT, default=5432"`
	User     string `env:"POSTGRES_USER, default=admin"`
	Password string `env:"POSTGRES_PASSWORD, default=postgres"`
	DBName   string `env:"POSTGRES_DB, default=evaluator"`
	DSN      string
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

	cfg.Postgres.DSN = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.DBName,
	)

	log.Println("Configuration loaded successfully")
}
