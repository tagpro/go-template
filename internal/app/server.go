package app

import (
	"log"
	"net/http"

	"github.com/tagpro/go-template/internal/config"
)

type Server struct {
	config *config.Config
}

func NewServer(c *config.Config) (*Server, error) {
	return &Server{
		config: c,
	}, nil
}

func (s *Server) StartServer() error {
	r := NewRouter()
	log.Printf("Server listening on port %s", s.config.Port)
	return http.ListenAndServe(":"+s.config.Port, r)
}
