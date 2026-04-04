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

func (s *Session) InProgress() bool {
	return s.EndedAt == nil
}

type SessionRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*Session, error)
	ListByUserID(ctx context.Context, userID uuid.UUID) ([]*Session, error)
	Create(ctx context.Context, session *Session) error
	Update(ctx context.Context, session *Session) error
	Delete(ctx context.Context, id uuid.UUID) error
}
