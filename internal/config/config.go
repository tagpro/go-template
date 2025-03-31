package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("config error: failed to load .env file: %w", err)
		return nil, fmt.Errorf("config error: failed to load .env file: %w", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}

	return &Config{
		Port: port,
	}, nil
}
