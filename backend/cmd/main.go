package main

import (
	"life_grid/internal/gym/domain"
	"life_grid/internal/gym/exercises"
	"life_grid/internal/gym/session_exercises"
	"life_grid/internal/gym/sessions"
	"life_grid/internal/gym/sets"
	"life_grid/internal/gym/templates"
	"life_grid/internal/shared"
	"log"
)

func main() {

	dbConfig := shared.Config()

	connPool, err := shared.NewDBPool(dbConfig)
	if err != nil {
		log.Fatal("Failed to create a connection pool, error: ", err)
	}
	defer connPool.Close()
	log.Println("Successfully connected to the database!!")

	// initialize repositories

	// gym repositories
	var exerciseRepo domain.ExerciseRepository = exercises.NewExerciseRepository(connPool)
	var sessionRepo domain.SessionRepository = sessions.NewSessionRepository(connPool)
	var sessionExerciseRepo domain.SessionExerciseRepository = session_exercises.NewSessionExerciseRepository(connPool)
	var setRepo domain.SetRepository = sets.NewSetRepository(connPool)
	var templateRepo domain.TemplateRepository = templates.NewTemplateRepository(connPool)
	var templateExerciseRepo domain.TemplateExerciseRepository = templates.NewTemplateExerciseRepository(connPool)

	_ = exerciseRepo
	_ = sessionRepo
	_ = sessionExerciseRepo
	_ = setRepo
	_ = templateRepo
	_ = templateExerciseRepo

}
