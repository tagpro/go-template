package app

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/tagpro/go-template/internal/hello"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	// Define your routes here, for example:
	r.Route("/hello", func(r chi.Router) {
		helloHandler := hello.New()
		r.Get("/", helloHandler.Get)

		r.Route("/{name}", func(r chi.Router) {
			r.Use(HelloMiddleware)
			r.Get("/", helloHandler.GetWithName)
		})
	})

	return r
}

func HelloMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		ctx := context.WithValue(r.Context(), hello.NameKey, name)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
