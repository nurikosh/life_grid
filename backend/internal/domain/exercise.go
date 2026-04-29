package domain

import (
	"context"

	"github.com/google/uuid"
)

type Exercise struct {
	ID          uuid.UUID
	Name        string
	MuscleGroup string
	Description string
}

type ExerciseRepository interface {
	Save(ctx context.Context, e *Exercise) error
	List(ctx context.Context) ([]Exercise, error)
	GetByID(ctx context.Context, id uuid.UUID) (*Exercise, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

func NewExercise(Name, MuscleGroup, Description string) (*Exercise, error) {
	if Name == "" {
		return nil, ErrNameIsRequired
	}

	if MuscleGroup == "" {
		return nil, ErrMuscleGroupIsRequired
	}

	if len(Description) > 250 {
		return nil, ErrDescriptionIsTooLong
	}

	return &Exercise{
		ID:          uuid.New(),
		Name:        Name,
		MuscleGroup: MuscleGroup,
		Description: Description,
	}, nil
}
