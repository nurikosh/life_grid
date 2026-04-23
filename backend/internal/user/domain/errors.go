package domain

import "errors"

var (
	// Identity errors.
	ErrUserIDRequired     = errors.New("user ID is required")
	ErrInvalidCredentials = errors.New("invalid credentials")

	// Email errors.
	ErrEmailRequired = errors.New("email is required")
	ErrEmailInvalid  = errors.New("email is invalid")
	ErrEmailExists   = errors.New("email already exists")

	// Password errors.
	ErrPasswordRequired   = errors.New("password hash is required")
	ErrPasswordIsRequired = errors.New("password is required")
	ErrPasswordTooShort   = errors.New("password must be at least 8 characters")

	// Profile metrics errors.
	ErrWeightInvalid = errors.New("weight cannot be negative")
	ErrHeightInvalid = errors.New("height cannot be negative")
)
