package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port string
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

	config := &Config{}
	err = viper.Unmarshal(config)
	if err != nil {
		log.Println("config error: failed to unmarshal config: %w", err)
		return nil, fmt.Errorf("config error: failed to unmarshal config: %w", err)
	}

	return config, nil
}
