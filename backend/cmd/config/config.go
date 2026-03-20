package config

import (
	"fmt"
	"github.com/spf13/viper"
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
	viper.Set("config-file", "./configs/local.yaml")
	viper.SetConfigFile(viper.GetString("config-file"))
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error to read file: %w", err)
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return fmt.Errorf("error to decode file: %w", err)
	}
	return nil
}
