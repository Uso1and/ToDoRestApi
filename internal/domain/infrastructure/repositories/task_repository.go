package repositories

import (
	"ToDoRestApi/internal/domain"
	"context"
	"database/sql"
)

type TaskRepository struct {
	db *sql.DB
}

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
