package domain

import (
	"context"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uuid.UUID
	Email     string
	Password  string
	FullName  string
	Weight    float64
	Height    float64
	CreatedAt time.Time
}

func NewUser(email, password, fullName string, weight, height float64) (*User, error) {
	if strings.TrimSpace(email) == "" {
		return nil, ErrEmailRequired
	}
	if !strings.Contains(email, "@") {
		return nil, ErrEmailInvalid
	}
	if strings.TrimSpace(password) == "" {
		return nil, ErrPasswordRequired
	}

	if weight < 0 {
		return nil, ErrWeightInvalid
	}

	if height < 0 {
		return nil, ErrHeightInvalid
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:        uuid.New(),
		Email:     strings.TrimSpace(strings.ToLower(email)),
		Password:  string(hash),
		FullName:  strings.TrimSpace(fullName),
		Weight:    weight,
		Height:    height,
		CreatedAt: time.Now(),
	}, nil
}

func (u *User) UpdateProfile(fullName string, weight, height float64) {
	u.FullName = strings.TrimSpace(fullName)
	u.Weight = weight
	u.Height = height
}

func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

type UserRepository interface {
	GetUserByID(ctx context.Context, id uuid.UUID) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	ListUsers(ctx context.Context) ([]*User, error)
	CreateUser(ctx context.Context, user *User) error
	UpdateUser(ctx context.Context, user *User) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
}
