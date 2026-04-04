package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Template struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Name      string
	Notes     string
	CreatedAt time.Time
	Exercises []TemplateExercise
}

type TemplateExercise struct {
	ID           uuid.UUID
	TemplateID   uuid.UUID
	ExerciseID   uuid.UUID
	OrderIndex   int
	TargetSets   int
	TargetReps   int
	TargetWeight float64
	Exercise     *Exercise
}

type TemplateRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*Template, error)
	ListByUserID(ctx context.Context, userID uuid.UUID) ([]*Template, error)
	Create(ctx context.Context, template *Template) error
	Update(ctx context.Context, template *Template) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type TemplateExerciseRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*TemplateExercise, error)
	ListByTemplateID(ctx context.Context, templateID uuid.UUID) ([]*TemplateExercise, error)
	Create(ctx context.Context, templateExercise *TemplateExercise) error
	Update(ctx context.Context, templateExercise *TemplateExercise) error
	Delete(ctx context.Context, id uuid.UUID) error
}
