package session_exercises

import (
	"context"
	"errors"
	"life_grid/internal/gym/domain"

	"github.com/google/uuid"
)

var (
	ErrSessionExerciseNotFound = errors.New("session exercise not found")
)

type SessionExerciseService interface {
	GetSessionExerciseByID(ctx context.Context, id uuid.UUID) (*domain.SessionExercise, error)
	ListSessionExercisesBySessionID(ctx context.Context, sessionID uuid.UUID) ([]*domain.SessionExercise, error)
	CreateSessionExercise(ctx context.Context, sessionID, exerciseID uuid.UUID) (*domain.SessionExercise, error)
	DeleteSessionExercise(ctx context.Context, id uuid.UUID) error
}

type sessionExerciseService struct {
	repo *SessionExerciseRepository
}

func NewSessionExerciseService(repo *SessionExerciseRepository) SessionExerciseService {
	return &sessionExerciseService{repo: repo}
}

func (s *sessionExerciseService) GetSessionExerciseByID(ctx context.Context, id uuid.UUID) (*domain.SessionExercise, error) {
	se, err := s.repo.GetSessionExerciseByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return se, nil
}

func (s *sessionExerciseService) ListSessionExercisesBySessionID(ctx context.Context, sessionID uuid.UUID) ([]*domain.SessionExercise, error) {
	items, err := s.repo.ListSessionExercisesBySessionID(ctx, sessionID)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *sessionExerciseService) CreateSessionExercise(ctx context.Context, sessionID, exerciseID uuid.UUID) (*domain.SessionExercise, error) {
	se, err := domain.NewSessionExercise(sessionID, exerciseID)
	if err != nil {
		return nil, err
	}

	if err := s.repo.CreateSessionExercise(ctx, se); err != nil {
		return nil, err
	}

	return se, nil
}

func (s *sessionExerciseService) DeleteSessionExercise(ctx context.Context, id uuid.UUID) error {
	_, err := s.repo.GetSessionExerciseByID(ctx, id)
	if err != nil {
		return ErrSessionExerciseNotFound
	}
	return s.repo.DeleteSessionExercise(ctx, id)
}
