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

func NewSet(sessionExerciseID uuid.UUID, reps int, weight float64, orderNum int) (*Set, error) {
	if sessionExerciseID == uuid.Nil {
		return nil, ErrSessionIDRequired
	}
	if reps <= 0 {
		return nil, ErrRepsPositive
	}
	if weight < 0 {
		return nil, ErrWeightNegative
	}
	if orderNum < 0 {
		return nil, ErrOrderIndexNegative
	}
	return &Set{
		ID:                uuid.New(),
		SessionExerciseID: sessionExerciseID,
		Reps:              reps,
		Weight:            weight,
		OrderNum:          orderNum,
	}, nil
}

type SetRepository interface {
	GetSetByID(ctx context.Context, id uuid.UUID) (*Set, error)
	ListSetsBySessionExerciseID(ctx context.Context, sessionExerciseID uuid.UUID) ([]*Set, error)
	CreateSet(ctx context.Context, set *Set) error
	UpdateSet(ctx context.Context, set *Set) error
	DeleteSet(ctx context.Context, id uuid.UUID) error
}
