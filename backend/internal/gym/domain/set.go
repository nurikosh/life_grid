package domain

import (
	"context"

	"github.com/google/uuid"
)

type Set struct {
	ID                uuid.UUID
	SessionExerciseID uuid.UUID
	Reps              int
	Weight            float64
	OrderNum          int
}

type SetRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*Set, error)
	ListBySessionExerciseID(ctx context.Context, sessionExerciseID uuid.UUID) ([]*Set, error)
	Create(ctx context.Context, set *Set) error
	Update(ctx context.Context, set *Set) error
	Delete(ctx context.Context, id uuid.UUID) error
}
