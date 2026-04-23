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

func NewTemplate(userID uuid.UUID, name, notes string) (*Template, error) {
	if userID == uuid.Nil {
		return nil, ErrUserIDRequired
	}
	if name == "" {
		return nil, ErrNameRequired
	}
	return &Template{
		ID:        uuid.New(),
		UserID:    userID,
		Name:      name,
		Notes:     notes,
		CreatedAt: time.Now(),
	}, nil
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

func NewTemplateExercise(templateID, exerciseID uuid.UUID, orderIndex, targetSets, targetReps int, targetWeight float64) (*TemplateExercise, error) {
	if templateID == uuid.Nil {
		return nil, ErrTemplateIDRequired
	}
	if exerciseID == uuid.Nil {
		return nil, ErrExerciseIDRequired
	}
	if orderIndex < 0 {
		return nil, ErrOrderIndexNegative
	}
	if targetSets <= 0 {
		return nil, ErrTargetSetsPositive
	}
	if targetReps <= 0 {
		return nil, ErrTargetRepsPositive
	}
	if targetWeight < 0 {
		return nil, ErrWeightNegative
	}
	return &TemplateExercise{
		ID:           uuid.New(),
		TemplateID:   templateID,
		ExerciseID:   exerciseID,
		OrderIndex:   orderIndex,
		TargetSets:   targetSets,
		TargetReps:   targetReps,
		TargetWeight: targetWeight,
	}, nil
}

type TemplateRepository interface {
	GetTemplateByID(ctx context.Context, id uuid.UUID) (*Template, error)
	ListTemplatesByUserID(ctx context.Context, userID uuid.UUID) ([]*Template, error)
	CreateTemplate(ctx context.Context, template *Template) error
	UpdateTemplate(ctx context.Context, template *Template) error
	DeleteTemplate(ctx context.Context, id uuid.UUID) error
}

type TemplateExerciseRepository interface {
	GetTemplateExerciseByID(ctx context.Context, id uuid.UUID) (*TemplateExercise, error)
	ListTemplateExercisesByTemplateID(ctx context.Context, templateID uuid.UUID) ([]*TemplateExercise, error)
	CreateTemplateExercise(ctx context.Context, templateExercise *TemplateExercise) error
	UpdateTemplateExercise(ctx context.Context, templateExercise *TemplateExercise) error
	DeleteTemplateExercise(ctx context.Context, id uuid.UUID) error
}
