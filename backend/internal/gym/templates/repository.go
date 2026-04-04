package templates

import (
	"context"
	"life_grid/internal/gym/domain"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TemplateRepository struct {
	pool *pgxpool.Pool
}

func NewTemplateRepository(pool *pgxpool.Pool) *TemplateRepository {
	return &TemplateRepository{pool: pool}
}

func (r *TemplateRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Template, error) {
	row := r.pool.QueryRow(ctx,
		`SELECT id, user_id, name, notes, created_at FROM templates WHERE id = $1`, id)

	t := &domain.Template{}
	if err := row.Scan(&t.ID, &t.UserID, &t.Name, &t.Notes, &t.CreatedAt); err != nil {
		return nil, err
	}
	return t, nil
}

func (r *TemplateRepository) ListByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.Template, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT id, user_id, name, notes, created_at FROM templates WHERE user_id = $1`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*domain.Template
	for rows.Next() {
		t := &domain.Template{}
		if err := rows.Scan(&t.ID, &t.UserID, &t.Name, &t.Notes, &t.CreatedAt); err != nil {
			return nil, err
		}
		result = append(result, t)
	}
	return result, rows.Err()
}

func (r *TemplateRepository) Create(ctx context.Context, template *domain.Template) error {
	_, err := r.pool.Exec(ctx,
		`INSERT INTO templates (id, user_id, name, notes, created_at) VALUES ($1, $2, $3, $4, $5)`,
		template.ID, template.UserID, template.Name, template.Notes, template.CreatedAt)
	return err
}

func (r *TemplateRepository) Update(ctx context.Context, template *domain.Template) error {
	_, err := r.pool.Exec(ctx,
		`UPDATE templates SET name = $1, notes = $2 WHERE id = $3`,
		template.Name, template.Notes, template.ID)
	return err
}

func (r *TemplateRepository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := r.pool.Exec(ctx, `DELETE FROM templates WHERE id = $1`, id)
	return err
}

// TemplateExerciseRepository

type TemplateExerciseRepository struct {
	pool *pgxpool.Pool
}

func NewTemplateExerciseRepository(pool *pgxpool.Pool) *TemplateExerciseRepository {
	return &TemplateExerciseRepository{pool: pool}
}

func (r *TemplateExerciseRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.TemplateExercise, error) {
	row := r.pool.QueryRow(ctx,
		`SELECT id, template_id, exercise_id, order_index, target_sets, target_reps, target_weight
		 FROM template_exercises WHERE id = $1`, id)

	te := &domain.TemplateExercise{}
	if err := row.Scan(&te.ID, &te.TemplateID, &te.ExerciseID, &te.OrderIndex,
		&te.TargetSets, &te.TargetReps, &te.TargetWeight); err != nil {
		return nil, err
	}
	return te, nil
}

func (r *TemplateExerciseRepository) ListByTemplateID(ctx context.Context, templateID uuid.UUID) ([]*domain.TemplateExercise, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT id, template_id, exercise_id, order_index, target_sets, target_reps, target_weight
		 FROM template_exercises WHERE template_id = $1 ORDER BY order_index`, templateID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*domain.TemplateExercise
	for rows.Next() {
		te := &domain.TemplateExercise{}
		if err := rows.Scan(&te.ID, &te.TemplateID, &te.ExerciseID, &te.OrderIndex,
			&te.TargetSets, &te.TargetReps, &te.TargetWeight); err != nil {
			return nil, err
		}
		result = append(result, te)
	}
	return result, rows.Err()
}

func (r *TemplateExerciseRepository) Create(ctx context.Context, te *domain.TemplateExercise) error {
	_, err := r.pool.Exec(ctx,
		`INSERT INTO template_exercises (id, template_id, exercise_id, order_index, target_sets, target_reps, target_weight)
		 VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		te.ID, te.TemplateID, te.ExerciseID, te.OrderIndex, te.TargetSets, te.TargetReps, te.TargetWeight)
	return err
}

func (r *TemplateExerciseRepository) Update(ctx context.Context, te *domain.TemplateExercise) error {
	_, err := r.pool.Exec(ctx,
		`UPDATE template_exercises SET order_index = $1, target_sets = $2, target_reps = $3, target_weight = $4 WHERE id = $5`,
		te.OrderIndex, te.TargetSets, te.TargetReps, te.TargetWeight, te.ID)
	return err
}

func (r *TemplateExerciseRepository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := r.pool.Exec(ctx, `DELETE FROM template_exercises WHERE id = $1`, id)
	return err
}
