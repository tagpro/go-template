package app

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/tagpro/go-template/internal/config"
	"github.com/tagpro/go-template/templates/helloworld"
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
	r := chi.NewRouter()

	// Define your routes here, for example:
	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		// return a templ response with Hello world
		component := helloworld.Hello("World")
		component.Render(r.Context(), w)
	})

	log.Printf("Server listening on port %s", s.config.Port)
	return http.ListenAndServe(":"+s.config.Port, r)
}
