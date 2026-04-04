package exercises

import (
	"context"
	"life_grid/internal/gym/domain"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ExerciseRepository struct {
	pool *pgxpool.Pool
}

func NewExerciseRepository(pool *pgxpool.Pool) *ExerciseRepository {
	return &ExerciseRepository{pool: pool}
}

func (r *ExerciseRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Exercise, error) {
	row := r.pool.QueryRow(ctx,
		`SELECT id, name, description, muscle_group FROM exercises WHERE id = $1`, id)

	e := &domain.Exercise{}
	if err := row.Scan(&e.ID, &e.Name, &e.Description, &e.MuscleGroup); err != nil {
		return nil, err
	}
	return e, nil
}

func (r *ExerciseRepository) List(ctx context.Context) ([]*domain.Exercise, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT id, name, description, muscle_group FROM exercises`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var exercises []*domain.Exercise
	for rows.Next() {
		e := &domain.Exercise{}
		if err := rows.Scan(&e.ID, &e.Name, &e.Description, &e.MuscleGroup); err != nil {
			return nil, err
		}
		exercises = append(exercises, e)
	}
	return exercises, rows.Err()
}

func (r *ExerciseRepository) Create(ctx context.Context, exercise *domain.Exercise) error {
	_, err := r.pool.Exec(ctx,
		`INSERT INTO exercises (id, name, description, muscle_group) VALUES ($1, $2, $3, $4)`,
		exercise.ID, exercise.Name, exercise.Description, exercise.MuscleGroup)
	return err
}

func (r *ExerciseRepository) Update(ctx context.Context, exercise *domain.Exercise) error {
	_, err := r.pool.Exec(ctx,
		`UPDATE exercises SET name = $1, description = $2, muscle_group = $3 WHERE id = $4`,
		exercise.Name, exercise.Description, exercise.MuscleGroup, exercise.ID)
	return err
}

func (r *ExerciseRepository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := r.pool.Exec(ctx, `DELETE FROM exercises WHERE id = $1`, id)
	return err
}
