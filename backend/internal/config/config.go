package config

import (
	"fmt"
	"log"
	"os"
)

type Cfg struct {
	Service  Service
	Queue    RabbitMQ
	Postgres Postgres
}

type Service struct {
	Host     string
	MainPort string
}

type RabbitMQ struct {
	Host      string
	PortUI    string
	PortAMQP  string
	User      string
	Password  string
	QueueName string
	URL       string
}

type Postgres struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	DSN      string
}

func (cfg *Cfg) MustLoad() {
	log.Println("Loading configuration from environment variables")

	cfg.loadFromEnv()
	cfg.validate()

	log.Println("Configuration loaded successfully")
}

func (cfg *Cfg) loadFromEnv() {
	cfg.Service.Host = getEnv("BACKEND_HOST", "0.0.0.0")
	cfg.Service.MainPort = getEnv("BACKEND_EXTERNAL_PORT", "8080")

	cfg.Queue.Host = getEnv("RABBITMQ_HOST", "localhost")
	cfg.Queue.PortUI = getEnv("RABBITMQ_PORT_UI", "15672")
	cfg.Queue.PortAMQP = getEnv("RABBITMQ_PORT_AMQP", "5672")
	cfg.Queue.User = getEnv("RABBITMQ_USER", "guest")
	cfg.Queue.Password = getEnv("RABBITMQ_PASSWORD", "guest")
	cfg.Queue.QueueName = getEnv("RABBITMQ_PHOTO_QUEUE_NAME", "photos_queue")

	cfg.Queue.URL = fmt.Sprintf("amqp://%s:%s@%s:%s/",
		cfg.Queue.User,
		cfg.Queue.Password,
		cfg.Queue.Host,
		cfg.Queue.PortAMQP,
	)

	cfg.Postgres.DBName = getEnv("POSTGRES_DB", "evaluator")
	cfg.Postgres.User = getEnv("POSTGRES_USER", "admin")
	cfg.Postgres.Password = getEnv("POSTGRES_PASSWORD", "postgres")
	cfg.Postgres.Port = getEnv("POSTGRES_PORT", "5432")
	cfg.Postgres.Host = getEnv("POSTGRES_HOST", "localhost")

	cfg.Postgres.DSN = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.DBName,
	)
}

func (cfg *Cfg) validate() {
	if cfg.Queue.Host == "" {
		log.Fatal("RABBITMQ_HOST is required")
	}
	if cfg.Queue.QueueName == "" {
		log.Fatal("RABBITMQ_PHOTO_QUEUE_NAME is required")
	}
	if cfg.Postgres.DBName == "" {
		log.Fatal("POSTGRES_DB is required")
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
