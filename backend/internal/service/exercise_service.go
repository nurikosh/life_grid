package service

import (
	"context"
	"fmt"
	"gym/internal/domain"

	"github.com/google/uuid"
)

type exerciseService struct {
	repo domain.ExerciseRepository
}

type ExerciseInput struct {
	Name        string
	MuscleGroup string
	Description string
}

type ExerciseService interface {
	Create(ctx context.Context, input ExerciseInput) (*domain.Exercise, error)
	List(ctx context.Context) ([]domain.Exercise, error)
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Exercise, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

func NewExerciseService(repo domain.ExerciseRepository) ExerciseService {
	return &exerciseService{repo: repo}
}

func (e exerciseService) Create(ctx context.Context, input ExerciseInput) (*domain.Exercise, error) {
	ex, err := domain.NewExercise(input.Name, input.MuscleGroup, input.Description)

	if err != nil {
		return nil, fmt.Errorf("exercise service: %w", err)
	}

	if err := e.repo.Save(ctx, ex); err != nil {
		return nil, err
	}
	return ex, nil

}

func (e exerciseService) List(ctx context.Context) ([]domain.Exercise, error) {
	return e.repo.List(ctx)
}

func (e exerciseService) GetByID(ctx context.Context, id uuid.UUID) (*domain.Exercise, error) {
	return e.repo.GetByID(ctx, id)
}

func (e exerciseService) Delete(ctx context.Context, id uuid.UUID) error {
	return e.repo.Delete(ctx, id)

}
