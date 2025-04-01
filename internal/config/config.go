package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"github.com/tagpro/go-template/pkg/logger"
)

type Config struct {
	Port   string
	Logger *logger.LogConfig
}

// LoadConfig loads config from file. You should be able to use env file for config as well.
// The default config file is located at ./config/config.yaml.
func LoadConfig() (*Config, error) {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("config error: failed to load config: %w", err)
		return nil, fmt.Errorf("config error: failed to load config: %w", err)
	}

	config := &Config{
		Port: ":9000",
		Logger: &logger.LogConfig{
			Level:  0,
			Format: "json",
		},
	}
	err = viper.Unmarshal(config)
	if err != nil {
		log.Println("config error: failed to unmarshal config: %w", err)
		return nil, fmt.Errorf("config error: failed to unmarshal config: %w", err)
	}

	return config, nil
}
