package repositories

import (
	"ToDoRestApi/internal/domain"
	"context"
	"database/sql"
	"fmt"
)

type TaskRepositoryInterface interface {
	CreateTask(ctx context.Context, task *domain.Task) error
	GetTask(ctx context.Context, taskID int) (*domain.Task, error)
	UpdateTask(ctx context.Context, task *domain.Task) error
	DeleteTask(ctx context.Context, taskID int) error
}
type TaskRepository struct {
	db *sql.DB
}

var _ TaskRepositoryInterface = (*TaskRepository)(nil)

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) CreateTask(c context.Context, task *domain.Task) error {

	query := `INSERT INTO tasks (title, description, done, created_at) VALUES ($1, $2, $3, $4) RETURNING id`

	return r.db.QueryRowContext(
		c,
		query,
		task.Title,
		task.Description,
		task.Done,
		task.CreatedAt,
	).Scan(&task.ID)
}

func (r *TaskRepository) GetTask(c context.Context, taskID int) (*domain.Task, error) {

	task := &domain.Task{}

	query := `SELECT title, description, done, created_at FROM tasks WHERE id = $1`

	err := r.db.QueryRowContext(c, query, taskID).Scan(
		&task.Title,
		&task.Description,
		&task.Done,
		&task.CreatedAt,
	)

	if err != nil {
		return nil, err
	}
	return task, nil
}

func (r *TaskRepository) UpdateTask(c context.Context, task *domain.Task) error {
	query := `UPDATE tasks SET title = $1, description = $2, done = $3, created_at = $4 WHERE id = $5`

	_, err := r.db.ExecContext(c, query, task.Title, task.Description, task.Done, task.CreatedAt, task.ID)

	if err != nil {
		return fmt.Errorf("error update task:%v", err)
	}
	return nil
}

func (r *TaskRepository) DeleteTask(c context.Context, taskID int) error {

	query := `DELETE FROM tasks WHERE id = $1`

	result, err := r.db.ExecContext(c, query, taskID)

	if err != nil {
		return fmt.Errorf("failet to delite task: %v", err)
	}

	rowsAffect, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failet to get rows affect: %w", err)
	}
	if rowsAffect == 0 {
		return sql.ErrNoRows
	}
	return nil
}
