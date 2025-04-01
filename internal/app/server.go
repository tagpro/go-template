package app

import (
	"log"
	"net/http"
	"strconv"

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
	log.Printf("Server listening on port %d", s.config.Port)
	return http.ListenAndServe(":"+strconv.Itoa(s.config.Port), r)
}
