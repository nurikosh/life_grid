package app

import (
	"gym/internal/infra/postgres"
	"gym/internal/service"
	"gym/internal/transport"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	Router http.Handler
}

func New(db *pgxpool.Pool) *App {
	// repos
	exerciseRepo := postgres.NewExerciseRepository(db)

	// services
	exerciseService := service.NewExerciseService(exerciseRepo)

	// handlers
	handlers := &transport.Handlers{
		Exercise: transport.NewExerciseHandler(exerciseService),
	}

	return &App{
		Router: transport.NewRouter(handlers),
	}
}
