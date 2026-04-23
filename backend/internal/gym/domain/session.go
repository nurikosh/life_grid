package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Name      string
	CreatedAt time.Time
	EndedAt   *time.Time
}

func NewSession(userID uuid.UUID, name string) (*Session, error) {
	if userID == uuid.Nil {
		return nil, ErrUserIDRequired
	}
	if name == "" {
		return nil, ErrNameRequired
	}
	return &Session{
		ID:        uuid.New(),
		UserID:    userID,
		Name:      name,
		CreatedAt: time.Now(),
	}, nil
}

func (s *Session) End(at time.Time) error {
	if !s.InProgress() {
		return ErrSessionEnded
	}
	s.EndedAt = &at
	return nil
}

func (s *Session) InProgress() bool {
	return s.EndedAt == nil
}

type SessionRepository interface {
	GetSessionByID(ctx context.Context, id uuid.UUID) (*Session, error)
	ListSessionsByUserID(ctx context.Context, userID uuid.UUID) ([]*Session, error)
	CreateSession(ctx context.Context, session *Session) error
	UpdateSession(ctx context.Context, session *Session) error
	DeleteSession(ctx context.Context, id uuid.UUID) error
}
