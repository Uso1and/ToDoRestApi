package repositories

import (
	"ToDoRestApi/internal/domain"
	"context"
	"database/sql"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", "user=postgres password=root dbname=test_todo sslmode=disable")

	if err != nil {
		t.Fatal(err)
	}
	return db
}

func TestCreateTask(t *testing.T) {
	db := setupTestDB(t)

	repo := NewTaskRepository(db)

	c := context.Background()

	task := &domain.Task{
		Title:       "Test_title",
		Description: "testDescript",
		Done:        false,
		CreatedAt:   time.Now(),
	}

	err := repo.CreateTask(c, task)
	assert.NoError(t, err)
	assert.NotZero(t, task.ID)
}
