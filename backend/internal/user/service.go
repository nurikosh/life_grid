package user

import (
	"context"
	"life_grid/internal/user/domain"

	"github.com/google/uuid"
)

type UserService interface {
	GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	ListUsers(ctx context.Context) ([]*domain.User, error)
	Register(ctx context.Context, email, passwordHash, fullName string, weight, height float64) (*domain.User, error)
	Login(ctx context.Context, email, password string) (string, error)
	UpdateUserProfile(ctx context.Context, id uuid.UUID, fullName string, weight, height float64) (*domain.User, error)
	DeleteUser(ctx context.Context, id uuid.UUID) error
}

type userService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	return s.repo.GetUserByID(ctx, id)
}

func (s *userService) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	return s.repo.GetUserByEmail(ctx, email)
}

func (s *userService) ListUsers(ctx context.Context) ([]*domain.User, error) {
	return s.repo.ListUsers(ctx)
}

func (s *userService) Register(ctx context.Context, email, passwordHash, fullName string, weight, height float64) (*domain.User, error) {
	existingUser, err := s.repo.GetUserByEmail(ctx, email)
	if err == nil && existingUser != nil {
		return nil, domain.ErrEmailExists
	}

	user, err := domain.NewUser(email, passwordHash, fullName, weight, height)
	if err != nil {
		return nil, err
	}

	if err := s.repo.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)

	if err != nil {
		return "", domain.ErrInvalidCredentials
	}

	if err := user.CheckPassword(password); err != nil {
		return "", domain.ErrInvalidCredentials
	}

	token, err := GenerateJWT(user.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}
func (s *userService) UpdateUserProfile(ctx context.Context, id uuid.UUID, fullName string, weight, height float64) (*domain.User, error) {
	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	user.UpdateProfile(fullName, weight, height)

	if err := s.repo.UpdateUser(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteUser(ctx, id)
}
