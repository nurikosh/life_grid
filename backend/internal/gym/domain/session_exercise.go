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

func NewSessionExercise(sessionID, exerciseID uuid.UUID) (*SessionExercise, error) {
	if sessionID == uuid.Nil {
		return nil, ErrSessionIDRequired
	}
	if exerciseID == uuid.Nil {
		return nil, ErrExerciseIDRequired
	}
	return &SessionExercise{
		ID:         uuid.New(),
		SessionID:  sessionID,
		ExerciseID: exerciseID,
	}, nil
}

func (se *SessionExercise) TotalVolume() float64 {
	var total float64
	for _, s := range se.Sets {
		total += s.Weight * float64(s.Reps)
	}
	return total
}

type SessionExerciseRepository interface {
	GetSessionExerciseByID(ctx context.Context, id uuid.UUID) (*SessionExercise, error)
	ListSessionExercisesBySessionID(ctx context.Context, sessionID uuid.UUID) ([]*SessionExercise, error)
	CreateSessionExercise(ctx context.Context, sessionExercise *SessionExercise) error
	DeleteSessionExercise(ctx context.Context, id uuid.UUID) error
}
