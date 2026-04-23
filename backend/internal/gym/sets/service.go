package sets

import (
	"context"
	"errors"
	"life_grid/internal/gym/domain"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

var (
	ErrSetNotFound = errors.New("set not found")
)

type SetService interface {
	GetSetByID(ctx context.Context, id uuid.UUID) (*domain.Set, error)
	ListSetsBySessionExerciseID(ctx context.Context, sessionExerciseID uuid.UUID) ([]*domain.Set, error)
	CreateSet(ctx context.Context, sessionExerciseID uuid.UUID, reps int, weight float64, orderNum int) (*domain.Set, error)
	UpdateSet(ctx context.Context, id uuid.UUID, reps int, weight float64, orderNum int) (*domain.Set, error)
	DeleteSet(ctx context.Context, id uuid.UUID) error
}

type setService struct {
	repo domain.SetRepository
}

func NewSetService(repo domain.SetRepository) SetService {
	return &setService{repo: repo}
}

func (s *setService) GetSetByID(ctx context.Context, id uuid.UUID) (*domain.Set, error) {
	set, err := s.repo.GetSetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrSetNotFound
		}
		return nil, err
	}
	return set, nil
}

func (s *setService) ListSetsBySessionExerciseID(ctx context.Context, sessionExerciseID uuid.UUID) ([]*domain.Set, error) {
	sets, err := s.repo.ListSetsBySessionExerciseID(ctx, sessionExerciseID)
	if err != nil {
		return nil, err
	}
	return sets, nil
}

func (s *setService) CreateSet(ctx context.Context, sessionExerciseID uuid.UUID, reps int, weight float64, orderNum int) (*domain.Set, error) {
	set, err := domain.NewSet(sessionExerciseID, reps, weight, orderNum)
	if err != nil {
		return nil, err
	}

	if err := s.repo.CreateSet(ctx, set); err != nil {
		return nil, err
	}

	return set, nil
}

func (s *setService) UpdateSet(ctx context.Context, id uuid.UUID, reps int, weight float64, orderNum int) (*domain.Set, error) {
	set, err := s.repo.GetSetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrSetNotFound
		}
		return nil, err
	}

	updatedSet, err := domain.NewSet(set.SessionExerciseID, reps, weight, orderNum)
	if err != nil {
		return nil, err
	}

	set.Reps = updatedSet.Reps
	set.Weight = updatedSet.Weight
	set.OrderNum = updatedSet.OrderNum

	if err := s.repo.UpdateSet(ctx, set); err != nil {
		return nil, err
	}

	return set, nil
}

func (s *setService) DeleteSet(ctx context.Context, id uuid.UUID) error {
	_, err := s.repo.GetSetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrSetNotFound
		}
		return err
	}

	return s.repo.DeleteSet(ctx, id)
}
