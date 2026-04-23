package sets

import (
	"context"
	"life_grid/internal/gym/domain"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SetRepository struct {
	pool *pgxpool.Pool
}

func NewSetRepository(pool *pgxpool.Pool) *SetRepository {
	return &SetRepository{pool: pool}
}

func (r *SetRepository) GetSetByID(ctx context.Context, id uuid.UUID) (*domain.Set, error) {
	row := r.pool.QueryRow(ctx,
		`SELECT id, session_exercise_id, reps, weight, order_num FROM sets WHERE id = $1`, id)

	s := &domain.Set{}
	if err := row.Scan(&s.ID, &s.SessionExerciseID, &s.Reps, &s.Weight, &s.OrderNum); err != nil {
		return nil, err
	}
	return s, nil
}

func (r *SetRepository) ListSetsBySessionExerciseID(ctx context.Context, sessionExerciseID uuid.UUID) ([]*domain.Set, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT id, session_exercise_id, reps, weight, order_num FROM sets WHERE session_exercise_id = $1 ORDER BY order_num`,
		sessionExerciseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*domain.Set
	for rows.Next() {
		s := &domain.Set{}
		if err := rows.Scan(&s.ID, &s.SessionExerciseID, &s.Reps, &s.Weight, &s.OrderNum); err != nil {
			return nil, err
		}
		result = append(result, s)
	}
	return result, rows.Err()
}

func (r *SetRepository) CreateSet(ctx context.Context, set *domain.Set) error {
	_, err := r.pool.Exec(ctx,
		`INSERT INTO sets (id, session_exercise_id, reps, weight, order_num) VALUES ($1, $2, $3, $4, $5)`,
		set.ID, set.SessionExerciseID, set.Reps, set.Weight, set.OrderNum)
	return err
}

func (r *SetRepository) UpdateSet(ctx context.Context, set *domain.Set) error {
	_, err := r.pool.Exec(ctx,
		`UPDATE sets SET reps = $1, weight = $2, order_num = $3 WHERE id = $4`,
		set.Reps, set.Weight, set.OrderNum, set.ID)
	return err
}

func (r *SetRepository) DeleteSet(ctx context.Context, id uuid.UUID) error {
	_, err := r.pool.Exec(ctx, `DELETE FROM sets WHERE id = $1`, id)
	return err
}
