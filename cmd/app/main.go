package main

import (
	"log"

	"github.com/tagpro/go-template/internal/app"
	"github.com/tagpro/go-template/internal/config"
	"github.com/tagpro/go-template/pkg/logger"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	logger.InitLogger(config.Logger)

	server, err := app.NewServer(config)
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}

	err = server.StartServer()
	if err != nil {
		log.Fatalf("shutting down with error: %v", err)
	}
}
