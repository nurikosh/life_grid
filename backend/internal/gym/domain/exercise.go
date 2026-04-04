package domain

import (
	"context"

	"github.com/google/uuid"
)

type Exercise struct {
	ID          uuid.UUID
	Name        string
	Description string
	MuscleGroup string
}

func NewExercise(name, description, muscleGroup string) *Exercise {
	return &Exercise{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
		MuscleGroup: muscleGroup,
	}

}

type ExerciseRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*Exercise, error)
	List(ctx context.Context) ([]*Exercise, error)
	Create(ctx context.Context, exercise *Exercise) error
	Update(ctx context.Context, exercise *Exercise) error
	Delete(ctx context.Context, id uuid.UUID) error
}
