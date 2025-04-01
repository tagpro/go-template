package logger

import (
	"log/slog"
	"os"
)

type LogConfig struct {
	Level  slog.Level
	Format string // Can be "text" or "json"
}

func InitLogger(config *LogConfig) {
	var handler slog.Handler
	switch config.Format {
	case "json":
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: config.Level,
		})
	case "text":
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: config.Level,
		})
	default:
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: config.Level,
		})
	}
	logger := slog.New(handler)
	slog.SetDefault(logger)
}
