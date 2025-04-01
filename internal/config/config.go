package config

import (
	"fmt"
	"log/slog"

	"github.com/spf13/viper"
	"github.com/tagpro/go-template/pkg/logger"
)

type Config struct {
	Port   int
	Logger *logger.LogConfig
}

// LoadConfig loads config from file. You should be able to use env file for config as well.
// The default config file is located at ./config/config.yaml.
func LoadConfig() (*Config, error) {
	config := &Config{
		Port: 9000,
		Logger: &logger.LogConfig{
			Level:  slog.LevelDebug,
			Format: "json",
		},
	}

	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		slog.Error("config error: failed to find config file", slog.Any("error", err))
		slog.Info("continue with default config")
		return config, nil
	}

	err = viper.Unmarshal(config)
	if err != nil {
		slog.Error("config error: failed to unmarshal config", slog.Any("error", err))
		return nil, fmt.Errorf("config error: failed to unmarshal config: %w", err)
	}

	return config, nil
}
