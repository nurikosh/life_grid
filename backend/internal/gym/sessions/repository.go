package sessions

import (
	"context"
	"life_grid/internal/gym/domain"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SessionRepository struct {
	pool *pgxpool.Pool
}

func NewSessionRepository(pool *pgxpool.Pool) *SessionRepository {
	return &SessionRepository{pool: pool}
}

func (r *SessionRepository) GetSessionByID(ctx context.Context, id uuid.UUID) (*domain.Session, error) {
	row := r.pool.QueryRow(ctx,
		`SELECT id, user_id, name, created_at, ended_at FROM sessions WHERE id = $1`, id)

	s := &domain.Session{}
	if err := row.Scan(&s.ID, &s.UserID, &s.Name, &s.CreatedAt, &s.EndedAt); err != nil {
		return nil, err
	}
	return s, nil
}

func (r *SessionRepository) ListSessionsByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.Session, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT id, user_id, name, created_at, ended_at FROM sessions WHERE user_id = $1`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*domain.Session
	for rows.Next() {
		s := &domain.Session{}
		if err := rows.Scan(&s.ID, &s.UserID, &s.Name, &s.CreatedAt, &s.EndedAt); err != nil {
			return nil, err
		}
		result = append(result, s)
	}
	return result, rows.Err()
}

func (r *SessionRepository) CreateSession(ctx context.Context, session *domain.Session) error {
	_, err := r.pool.Exec(ctx,
		`INSERT INTO sessions (id, user_id, name, created_at, ended_at) VALUES ($1, $2, $3, $4, $5)`,
		session.ID, session.UserID, session.Name, session.CreatedAt, session.EndedAt)
	return err
}

func (r *SessionRepository) UpdateSession(ctx context.Context, session *domain.Session) error {
	_, err := r.pool.Exec(ctx,
		`UPDATE sessions SET name = $1, ended_at = $2 WHERE id = $3`,
		session.Name, session.EndedAt, session.ID)
	return err
}

func (r *SessionRepository) DeleteSession(ctx context.Context, id uuid.UUID) error {
	_, err := r.pool.Exec(ctx, `DELETE FROM sessions WHERE id = $1`, id)
	return err
}
