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
func TestGetTask(t *testing.T) {
	db := setupTestDB(t)
	repo := NewTaskRepository(db)
	ctx := context.Background()

	createdTask := &domain.Task{
		Title:       "Test Get Task",
		Description: "Description for get test",
		Done:        true,
		CreatedAt:   time.Now(),
	}

	err := repo.CreateTask(ctx, createdTask)
	assert.NoError(t, err)
	assert.NotZero(t, createdTask.ID)

	t.Logf("Created task ID: %d", createdTask.ID)

	retrievedTask, err := repo.GetTask(ctx, createdTask.ID)
	assert.NoError(t, err)

	assert.Equal(t, createdTask.ID, retrievedTask.ID)
	assert.Equal(t, createdTask.Title, retrievedTask.Title)
	assert.Equal(t, createdTask.Description, retrievedTask.Description)
	assert.Equal(t, createdTask.Done, retrievedTask.Done)
	assert.WithinDuration(t, createdTask.CreatedAt, retrievedTask.CreatedAt, time.Second)
}

func TestGetTask_NotFound(t *testing.T) {
	db := setupTestDB(t)
	repo := NewTaskRepository(db)
	ctx := context.Background()

	_, err := repo.GetTask(ctx, 99999)
	assert.Error(t, err)
	assert.Equal(t, sql.ErrNoRows, err)
}

func TestUpdateTask(t *testing.T) {
	db := setupTestDB(t)
	repo := NewTaskRepository(db)
	ctx := context.Background()

	originalTask := &domain.Task{
		Title:       "Original Title",
		Description: "Original Description",
		Done:        false,
		CreatedAt:   time.Now(),
	}

	err := repo.CreateTask(ctx, originalTask)
	assert.NoError(t, err)
	assert.NotZero(t, originalTask.ID)

	updatedTask := &domain.Task{
		ID:          originalTask.ID,
		Title:       "Updated Title",
		Description: "Updated Description",
		Done:        true,
		CreatedAt:   originalTask.CreatedAt.Add(24 * time.Hour),
	}

	err = repo.UpdateTask(ctx, updatedTask)
	assert.NoError(t, err)

	retrievedTask, err := repo.GetTask(ctx, originalTask.ID)
	assert.NoError(t, err)

	assert.Equal(t, updatedTask.ID, retrievedTask.ID)
	assert.Equal(t, updatedTask.Title, retrievedTask.Title)
	assert.Equal(t, updatedTask.Description, retrievedTask.Description)
	assert.Equal(t, updatedTask.Done, retrievedTask.Done)
	assert.WithinDuration(t, updatedTask.CreatedAt, retrievedTask.CreatedAt, time.Second)
}

func TestDeleteTask(t *testing.T) {
	db := setupTestDB(t)
	repo := NewTaskRepository(db)
	ctx := context.Background()

	taskToDelete := &domain.Task{
		Title:       "Task to delete",
		Description: "Will be deleted soon",
		Done:        false,
		CreatedAt:   time.Now(),
	}

	err := repo.CreateTask(ctx, taskToDelete)
	assert.NoError(t, err)
	assert.NotZero(t, taskToDelete.ID)

	err = repo.DeleteTask(ctx, taskToDelete.ID)
	assert.NoError(t, err)

	_, err = repo.GetTask(ctx, taskToDelete.ID)
	assert.Error(t, err)
	assert.Equal(t, sql.ErrNoRows, err)
}

func TestDeleteTask_NotFound(t *testing.T) {
	db := setupTestDB(t)
	repo := NewTaskRepository(db)
	ctx := context.Background()

	err := repo.DeleteTask(ctx, 999999)
	assert.Error(t, err)
	assert.Equal(t, sql.ErrNoRows, err)
}

func TestDeleteTask_VerifyRowsAffected(t *testing.T) {
	db := setupTestDB(t)
	repo := NewTaskRepository(db)
	ctx := context.Background()

	task := &domain.Task{
		Title:       "For deletion check",
		Description: "Check rows affected",
		Done:        true,
		CreatedAt:   time.Now(),
	}
	err := repo.CreateTask(ctx, task)
	assert.NoError(t, err)

	err = repo.DeleteTask(ctx, task.ID)
	assert.NoError(t, err)

	err = repo.DeleteTask(ctx, task.ID)
	assert.Error(t, err)
	assert.Equal(t, sql.ErrNoRows, err)
}
