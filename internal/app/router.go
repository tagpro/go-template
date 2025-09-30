package app

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	// Health check endpoints
	r.Get("/z/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	r.Get("/z/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "pong"})
	})

	// Define your API routes here
	r.Route("/v1", func(r chi.Router) {
		// Example: Add your routes here
		// r.Get("/users", userHandler.List)
		// r.Post("/users", userHandler.Create)
		// r.Route("/users/{id}", func(r chi.Router) {
		//     r.Get("/", userHandler.Get)
		//     r.Put("/", userHandler.Update)
		//     r.Delete("/", userHandler.Delete)
		// })
	})

	return r
}
