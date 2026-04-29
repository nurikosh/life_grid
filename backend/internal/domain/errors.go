package domain

import "errors"

var (
	//exercise
	ErrNameIsRequired        = errors.New("name is required")
	ErrMuscleGroupIsRequired = errors.New("muscle group is required")
	ErrDescriptionIsTooLong  = errors.New("description is too long")
)
