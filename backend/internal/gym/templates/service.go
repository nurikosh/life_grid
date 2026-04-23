package templates

import (
	"context"
	"errors"
	"life_grid/internal/gym/domain"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

var (
	ErrTemplateNotFound         = errors.New("template not found")
	ErrTemplateExerciseNotFound = errors.New("template exercise not found")
)

type TemplateService interface {
	GetTemplateByID(ctx context.Context, id uuid.UUID) (*domain.Template, error)
	ListTemplatesByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.Template, error)
	CreateTemplate(ctx context.Context, userID uuid.UUID, name, notes string) (*domain.Template, error)
	UpdateTemplate(ctx context.Context, id uuid.UUID, name, notes string) (*domain.Template, error)
	DeleteTemplate(ctx context.Context, id uuid.UUID) error
}

type TemplateExerciseService interface {
	GetTemplateExerciseByID(ctx context.Context, id uuid.UUID) (*domain.TemplateExercise, error)
	ListTemplateExercisesByTemplateID(ctx context.Context, templateID uuid.UUID) ([]*domain.TemplateExercise, error)
	CreateTemplateExercise(ctx context.Context, templateID, exerciseID uuid.UUID, orderIndex, targetSets, targetReps int, targetWeight float64) (*domain.TemplateExercise, error)
	UpdateTemplateExercise(ctx context.Context, id uuid.UUID, orderIndex, targetSets, targetReps int, targetWeight float64) (*domain.TemplateExercise, error)
	DeleteTemplateExercise(ctx context.Context, id uuid.UUID) error
}

type templateService struct {
	repo domain.TemplateRepository
}

type templateExerciseService struct {
	repo domain.TemplateExerciseRepository
}

func NewTemplateService(repo domain.TemplateRepository) TemplateService {
	return &templateService{repo: repo}
}

func NewTemplateExerciseService(repo domain.TemplateExerciseRepository) TemplateExerciseService {
	return &templateExerciseService{repo: repo}
}

func (s *templateService) GetTemplateByID(ctx context.Context, id uuid.UUID) (*domain.Template, error) {
	template, err := s.repo.GetTemplateByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrTemplateNotFound
		}
		return nil, err
	}
	return template, nil
}

func (s *templateService) ListTemplatesByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.Template, error) {
	templates, err := s.repo.ListTemplatesByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return templates, nil
}

func (s *templateService) CreateTemplate(ctx context.Context, userID uuid.UUID, name, notes string) (*domain.Template, error) {
	template, err := domain.NewTemplate(userID, name, notes)
	if err != nil {
		return nil, err
	}

	if err := s.repo.CreateTemplate(ctx, template); err != nil {
		return nil, err
	}

	return template, nil
}

func (s *templateService) UpdateTemplate(ctx context.Context, id uuid.UUID, name, notes string) (*domain.Template, error) {
	template, err := s.repo.GetTemplateByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrTemplateNotFound
		}
		return nil, err
	}

	updatedTemplate, err := domain.NewTemplate(template.UserID, name, notes)
	if err != nil {
		return nil, err
	}

	template.Name = updatedTemplate.Name
	template.Notes = updatedTemplate.Notes

	if err := s.repo.UpdateTemplate(ctx, template); err != nil {
		return nil, err
	}

	return template, nil
}

func (s *templateService) DeleteTemplate(ctx context.Context, id uuid.UUID) error {
	_, err := s.repo.GetTemplateByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrTemplateNotFound
		}
		return err
	}

	return s.repo.DeleteTemplate(ctx, id)
}

func (s *templateExerciseService) GetTemplateExerciseByID(ctx context.Context, id uuid.UUID) (*domain.TemplateExercise, error) {
	templateExercise, err := s.repo.GetTemplateExerciseByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrTemplateExerciseNotFound
		}
		return nil, err
	}
	return templateExercise, nil
}

func (s *templateExerciseService) ListTemplateExercisesByTemplateID(ctx context.Context, templateID uuid.UUID) ([]*domain.TemplateExercise, error) {
	templateExercises, err := s.repo.ListTemplateExercisesByTemplateID(ctx, templateID)
	if err != nil {
		return nil, err
	}
	return templateExercises, nil
}

func (s *templateExerciseService) CreateTemplateExercise(ctx context.Context, templateID, exerciseID uuid.UUID, orderIndex, targetSets, targetReps int, targetWeight float64) (*domain.TemplateExercise, error) {
	templateExercise, err := domain.NewTemplateExercise(templateID, exerciseID, orderIndex, targetSets, targetReps, targetWeight)
	if err != nil {
		return nil, err
	}

	if err := s.repo.CreateTemplateExercise(ctx, templateExercise); err != nil {
		return nil, err
	}

	return templateExercise, nil
}

func (s *templateExerciseService) UpdateTemplateExercise(ctx context.Context, id uuid.UUID, orderIndex, targetSets, targetReps int, targetWeight float64) (*domain.TemplateExercise, error) {
	templateExercise, err := s.repo.GetTemplateExerciseByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrTemplateExerciseNotFound
		}
		return nil, err
	}

	updatedTemplateExercise, err := domain.NewTemplateExercise(templateExercise.TemplateID, templateExercise.ExerciseID, orderIndex, targetSets, targetReps, targetWeight)
	if err != nil {
		return nil, err
	}

	templateExercise.OrderIndex = updatedTemplateExercise.OrderIndex
	templateExercise.TargetSets = updatedTemplateExercise.TargetSets
	templateExercise.TargetReps = updatedTemplateExercise.TargetReps
	templateExercise.TargetWeight = updatedTemplateExercise.TargetWeight

	if err := s.repo.UpdateTemplateExercise(ctx, templateExercise); err != nil {
		return nil, err
	}

	return templateExercise, nil
}

func (s *templateExerciseService) DeleteTemplateExercise(ctx context.Context, id uuid.UUID) error {
	_, err := s.repo.GetTemplateExerciseByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrTemplateExerciseNotFound
		}
		return err
	}

	return s.repo.DeleteTemplateExercise(ctx, id)
}
