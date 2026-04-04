package domain

import (
	"context"

	"github.com/google/uuid"
)

type SessionExercise struct {
	ID         uuid.UUID
	ExerciseID uuid.UUID
	SessionID  uuid.UUID
	Exercise   *Exercise
	Sets       []Set
}

func (se *SessionExercise) TotalVolume() float64 {
	var total float64
	for _, s := range se.Sets {
		total += s.Weight * float64(s.Reps)
	}
	return total
}

type SessionExerciseRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*SessionExercise, error)
	ListBySessionID(ctx context.Context, sessionID uuid.UUID) ([]*SessionExercise, error)
	Create(ctx context.Context, sessionExercise *SessionExercise) error
	Delete(ctx context.Context, id uuid.UUID) error
}
