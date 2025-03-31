package main

import (
	"log"

	"github.com/tagpro/go-template/internal/app"
	"github.com/tagpro/go-template/internal/config"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	server, err := app.NewServer(config)
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}

	err = server.StartServer()
	if err != nil {
		log.Fatalf("shutting down with error: %v", err)
	}
}
