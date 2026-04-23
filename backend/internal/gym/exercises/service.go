package exercises

import (
	"context"
	"life_grid/internal/gym/domain"

	"github.com/google/uuid"
)

type ExerciseService interface {
	GetExerciseByID(ctx context.Context, id uuid.UUID) (*domain.Exercise, error)
	ListExercises(ctx context.Context) ([]*domain.Exercise, error)
	CreateExercise(ctx context.Context, name, description, muscleGroup string) (*domain.Exercise, error)
	UpdateExercise(ctx context.Context, id uuid.UUID, name, description, muscleGroup string) (*domain.Exercise, error)
	DeleteExercise(ctx context.Context, id uuid.UUID) error
}

type exerciseService struct {
	exerciseRepo domain.ExerciseRepository
}

func NewExerciseService(exerciseRepo domain.ExerciseRepository) *exerciseService {
	return &exerciseService{exerciseRepo: exerciseRepo}
}

func (e exerciseService) GetExerciseByID(ctx context.Context, id uuid.UUID) (*domain.Exercise, error) {
	exercise, err := e.exerciseRepo.GetExerciseByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return exercise, nil
}

func (e exerciseService) ListExercises(ctx context.Context) ([]*domain.Exercise, error) {
	return e.exerciseRepo.ListExercises(ctx)
}

func (e exerciseService) CreateExercise(ctx context.Context, name, description, muscleGroup string) (*domain.Exercise, error) {
	newExercise, err := domain.NewExercise(name, description, muscleGroup)

	if err != nil {
		return nil, err
	}

	if err := e.exerciseRepo.CreateExercise(ctx, newExercise); err != nil {
		return nil, err
	}

	return newExercise, nil
}

func (e exerciseService) UpdateExercise(ctx context.Context, id uuid.UUID, name, description, muscleGroup string) (*domain.Exercise, error) {
	exercise, err := e.exerciseRepo.GetExerciseByID(ctx, id)

	if err != nil {
		return nil, err
	}

	if err := exercise.Update(name, description, muscleGroup); err != nil {
		return nil, err
	}

	if err := e.exerciseRepo.UpdateExercise(ctx, exercise); err != nil {
		return nil, err
	}

	return exercise, nil

}

func (e exerciseService) DeleteExercise(ctx context.Context, id uuid.UUID) error {
	return e.exerciseRepo.DeleteExercise(ctx, id)
}
