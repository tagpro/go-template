package hello

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/tagpro/go-template/templates/helloworld"
)

type Hello interface {
	Get(w http.ResponseWriter, r *http.Request)
	GetWithName(w http.ResponseWriter, r *http.Request)
}

type contextKey string

const (
	NameKey contextKey = "name"
)

func New() Hello {
	return &hello{}
}

type hello struct{}

func (h *hello) Get(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello, World!")
	slog.Info("Hello, world ")
	if err != nil {
		log.Printf("Error writing response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *hello) GetWithName(w http.ResponseWriter, r *http.Request) {
	name := r.Context().Value(NameKey).(string)
	slog.Info("Hello", slog.String("name", name))
	component := helloworld.Hello(name)
	err := component.Render(r.Context(), w)
	if err != nil {
		log.Printf("Error rendering component: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
