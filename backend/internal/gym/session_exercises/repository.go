package session_exercises

import (
	"context"
	"life_grid/internal/gym/domain"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SessionExerciseRepository struct {
	pool *pgxpool.Pool
}

func NewSessionExerciseRepository(pool *pgxpool.Pool) *SessionExerciseRepository {
	return &SessionExerciseRepository{pool: pool}
}

func (r *SessionExerciseRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.SessionExercise, error) {
	row := r.pool.QueryRow(ctx,
		`SELECT id, exercise_id, session_id FROM session_exercises WHERE id = $1`, id)

	se := &domain.SessionExercise{}
	if err := row.Scan(&se.ID, &se.ExerciseID, &se.SessionID); err != nil {
		return nil, err
	}
	return se, nil
}

func (r *SessionExerciseRepository) ListBySessionID(ctx context.Context, sessionID uuid.UUID) ([]*domain.SessionExercise, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT id, exercise_id, session_id FROM session_exercises WHERE session_id = $1`, sessionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*domain.SessionExercise
	for rows.Next() {
		se := &domain.SessionExercise{}
		if err := rows.Scan(&se.ID, &se.ExerciseID, &se.SessionID); err != nil {
			return nil, err
		}
		result = append(result, se)
	}
	return result, rows.Err()
}

func (r *SessionExerciseRepository) Create(ctx context.Context, sessionExercise *domain.SessionExercise) error {
	_, err := r.pool.Exec(ctx,
		`INSERT INTO session_exercises (id, exercise_id, session_id) VALUES ($1, $2, $3)`,
		sessionExercise.ID, sessionExercise.ExerciseID, sessionExercise.SessionID)
	return err
}

func (r *SessionExerciseRepository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := r.pool.Exec(ctx, `DELETE FROM session_exercises WHERE id = $1`, id)
	return err
}
