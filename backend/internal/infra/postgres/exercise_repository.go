package postgres

import (
	"context"
	"gym/internal/domain"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type exerciseRepository struct {
	pool *pgxpool.Pool
}

func NewExerciseRepository(pool *pgxpool.Pool) domain.ExerciseRepository {
	return (&exerciseRepository{pool: pool})
}

func (e exerciseRepository) Save(ctx context.Context, exercise *domain.Exercise) error {
	_, err := e.pool.Exec(ctx,
		`INSERT INTO gym.exercises (id, name, muscle_group, description)
         VALUES ($1, $2, $3, $4)
         ON CONFLICT (id) DO UPDATE SET name=$2, muscle_group=$3, description=$4`,
		exercise.ID, exercise.Name, exercise.MuscleGroup, exercise.Description,
	)
	return err
}

func (e exerciseRepository) List(ctx context.Context) ([]domain.Exercise, error) {
	rows, err := e.pool.Query(ctx,
		`SELECT id, name, muscle_group, description
		 FROM gym.exercises
		 ORDER BY name`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	exercises := make([]domain.Exercise, 0)
	for rows.Next() {
		var ex domain.Exercise
		if err := rows.Scan(&ex.ID, &ex.Name, &ex.MuscleGroup, &ex.Description); err != nil {
			return nil, err
		}
		exercises = append(exercises, ex)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return exercises, nil
}

func (e exerciseRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Exercise, error) {
	var ex domain.Exercise
	err := e.pool.QueryRow(ctx,
		`SELECT id, name, muscle_group, description
		 FROM gym.exercises
		 WHERE id = $1`,
		id,
	).Scan(&ex.ID, &ex.Name, &ex.MuscleGroup, &ex.Description)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, pgx.ErrNoRows
		}
		return nil, err
	}

	return &ex, nil
}

func (e exerciseRepository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := e.pool.Exec(ctx, `DELETE FROM gym.exercises WHERE id = $1`, id)
	return err
}
