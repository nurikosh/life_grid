package main

import (
	"context"
	"gym/internal/app"
	"gym/internal/config"
	"gym/internal/infra/postgres"
	"log"
	"net/http"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	err = postgres.RunMigrations(cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	pool, err := postgres.NewPool(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connected to database")
	defer pool.Close()

	application := app.New(pool)

	log.Printf("server starting on %s", cfg.Port)
	if err := http.ListenAndServe(cfg.Port, application.Router); err != nil {
		log.Fatal(err)
	}

}
