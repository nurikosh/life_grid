package domain

import "errors"

var (
	ErrNameRequired        = errors.New("name is required")
	ErrMuscleGroupRequired = errors.New("muscle group is required")
	ErrUserIDRequired      = errors.New("user ID is required")
	ErrExerciseIDRequired  = errors.New("exercise ID is required")
	ErrSessionIDRequired   = errors.New("session ID is required")
	ErrTemplateIDRequired  = errors.New("template ID is required")
	ErrRepsPositive        = errors.New("reps must be greater than zero")
	ErrWeightNegative      = errors.New("weight must be non-negative")
	ErrTargetSetsPositive  = errors.New("target sets must be greater than zero")
	ErrTargetRepsPositive  = errors.New("target reps must be greater than zero")
	ErrOrderIndexNegative  = errors.New("order index must be non-negative")
	ErrSessionEnded        = errors.New("session has already ended")
	ErrDescriptionTooLong  = errors.New("description must be 100 characters or less")
)
