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

func NewExercise(name, description, muscleGroup string) (*Exercise, error) {
	if name == "" {
		return nil, ErrNameRequired
	}
	if muscleGroup == "" {
		return nil, ErrMuscleGroupRequired
	}

	if description == "" {
		description = "No description provided"
	}

	if len(description) > 300 {
		return nil, ErrDescriptionTooLong
	}
	return &Exercise{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
		MuscleGroup: muscleGroup,
	}, nil
}

func (e *Exercise) Update(name, description, muscleGroup string) error {
	if name == "" {
		return ErrNameRequired
	}
	if muscleGroup == "" {
		return ErrMuscleGroupRequired
	}

	if description == "" {
		description = "No description provided"
	}

	if len(description) > 300 {
		return ErrDescriptionTooLong
	}

	e.Name = name
	e.Description = description
	e.MuscleGroup = muscleGroup

	return nil
}

type ExerciseRepository interface {
	GetExerciseByID(ctx context.Context, id uuid.UUID) (*Exercise, error)
	ListExercises(ctx context.Context) ([]*Exercise, error)
	CreateExercise(ctx context.Context, exercise *Exercise) error
	UpdateExercise(ctx context.Context, exercise *Exercise) error
	DeleteExercise(ctx context.Context, id uuid.UUID) error
}
