package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

type Cfg struct {
	Service Service  `yaml:"service"`
	Queue   RabbitMQ `yaml:"queue"`
}

type Service struct {
	Host      string `yaml:"host"`
	MainPort  string `yaml:"mainPort"`
	DebugPort string `yaml:"debugPort"`
}

type RabbitMQ struct {
	RabbitMQURL string `yaml:"rabbitMQURL"`
	QueueName   string `yaml:"queueName"`
}

func (cfg *Cfg) Init() error {
	confPath := os.Getenv("CONFIG_FILE")
	log.Printf("Config path: %s", confPath)
	viper.SetConfigFile(confPath)
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error to read file: %w", err)
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return fmt.Errorf("error to decode file: %w", err)
	}
	return nil
}
