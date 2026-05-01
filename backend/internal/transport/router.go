package transport

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type Handlers struct {
	Exercise *ExerciseHandler
}

func NewRouter(h *Handlers) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	}))

	r.Route("/exercises", func(r chi.Router) {
		r.Post("/", h.Exercise.CreateExercise)
		r.Get("/", h.Exercise.ListExercises)
		r.Get("/{id}", h.Exercise.GetExerciseByID)
		r.Delete("/{id}", h.Exercise.DeleteExercise)
	})

	return r
}
