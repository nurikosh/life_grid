package user

import (
	"context"
	"errors"
	"life_grid/internal/user/domain"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrEmailExists  = errors.New("email already exists")
)

type userRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *userRepository {
	return &userRepository{pool: pool}
}

func (r *userRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	query := `SELECT id, email, password_hash, full_name, weight, height, created_at FROM identity.users WHERE id = $1`

	u := &domain.User{}
	if err := r.pool.QueryRow(ctx, query, id).Scan(
		&u.ID,
		&u.Email,
		&u.Password,
		&u.FullName,
		&u.Weight,
		&u.Height,
		&u.CreatedAt,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return u, nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := `SELECT id, email, password_hash, full_name, weight, height, created_at FROM identity.users WHERE email = $1`

	normalizedEmail := strings.TrimSpace(strings.ToLower(email))
	u := &domain.User{}
	if err := r.pool.QueryRow(ctx, query, normalizedEmail).Scan(
		&u.ID,
		&u.Email,
		&u.Password,
		&u.FullName,
		&u.Weight,
		&u.Height,
		&u.CreatedAt,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return u, nil
}

func (r *userRepository) ListUsers(ctx context.Context) ([]*domain.User, error) {
	query := `SELECT id, email, password_hash, full_name, weight, height, created_at FROM identity.users ORDER BY created_at DESC`

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*domain.User
	for rows.Next() {
		u := &domain.User{}
		if err := rows.Scan(
			&u.ID,
			&u.Email,
			&u.Password,
			&u.FullName,
			&u.Weight,
			&u.Height,
			&u.CreatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, rows.Err()
}

func (r *userRepository) CreateUser(ctx context.Context, user *domain.User) error {
	query := `INSERT INTO identity.users (id, email, password_hash, full_name, weight, height, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := r.pool.Exec(ctx, query,
		user.ID,
		strings.TrimSpace(strings.ToLower(user.Email)),
		user.Password,
		user.FullName,
		user.Weight,
		user.Height,
		user.CreatedAt,
	)
	if err == nil {
		return nil
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == "23505" {
		return ErrEmailExists
	}

	return err
}

func (r *userRepository) UpdateUser(ctx context.Context, user *domain.User) error {
	query := `UPDATE identity.users SET full_name = $1, weight = $2, height = $3 WHERE id = $4`

	cmd, err := r.pool.Exec(ctx, query, user.FullName, user.Weight, user.Height, user.ID)
	if err != nil {
		return err
	}
	if cmd.RowsAffected() == 0 {
		return ErrUserNotFound
	}

	return nil
}

func (r *userRepository) DeleteUser(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM identity.users WHERE id = $1`

	cmd, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if cmd.RowsAffected() == 0 {
		return ErrUserNotFound
	}

	return nil
}
