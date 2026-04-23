package sessions

import (
	"context"
	"life_grid/internal/gym/domain"
	"time"

	"github.com/google/uuid"
)

type SessionService interface {
	GetSessionByID(ctx context.Context, id uuid.UUID) (*domain.Session, error)
	ListSessionByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.Session, error)
	StartSession(ctx context.Context, userID uuid.UUID, name string) (*domain.Session, error)
	EndSession(ctx context.Context, id uuid.UUID, endedAt time.Time) (*domain.Session, error)
	DeleteSession(ctx context.Context, id uuid.UUID) error
}

type sessionService struct {
	repo domain.SessionRepository
}

func NewSessionService(repo domain.SessionRepository) *sessionService {
	return &sessionService{repo: repo}
}

func (s *sessionService) GetSessionByID(ctx context.Context, id uuid.UUID) (*domain.Session, error) {
	session, err := s.repo.GetSessionByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (s *sessionService) ListSessionByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.Session, error) {
	sessions, err := s.repo.ListSessionsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return sessions, nil
}

func (s *sessionService) StartSession(ctx context.Context, userID uuid.UUID, name string) (*domain.Session, error) {
	session, err := domain.NewSession(userID, name)
	if err != nil {
		return nil, err
	}
	if err := s.repo.CreateSession(ctx, session); err != nil {
		return nil, err
	}
	return session, nil
}

func (s *sessionService) EndSession(ctx context.Context, id uuid.UUID, endedAt time.Time) (*domain.Session, error) {

	session, err := s.repo.GetSessionByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if err := session.End(endedAt); err != nil {
		return nil, err
	}
	if err := s.repo.UpdateSession(ctx, session); err != nil {
		return nil, err
	}
	return session, nil
}

func (s *sessionService) DeleteSession(ctx context.Context, id uuid.UUID) error {
	_, err := s.repo.GetSessionByID(ctx, id)
	if err != nil {
		return err
	}
	return s.repo.DeleteSession(ctx, id)
}
