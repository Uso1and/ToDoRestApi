package handlers

import (
	"ToDoRestApi/internal/domain"
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) CreateTask(ctx context.Context, task *domain.Task) error {
	args := m.Called(ctx, task)
	return args.Error(0)
}

func (m *MockTaskRepository) GetTask(ctx context.Context, taskID int) (*domain.Task, error) {
	args := m.Called(ctx, taskID)
	return args.Get(0).(*domain.Task), args.Error(1)
}

func (m *MockTaskRepository) UpdateTask(ctx context.Context, task *domain.Task) error {
	args := m.Called(ctx, task)
	return args.Error(0)
}

func (m *MockTaskRepository) DeleteTask(ctx context.Context, taskID int) error {
	args := m.Called(ctx, taskID)
	return args.Error(0)
}

func TestCreateTaskHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := new(MockTaskRepository)
	handler := NewTaskHandler(mockRepo)

	testTask := domain.Task{Title: "Title_test"}

	mockRepo.On("CreateTask", mock.Anything, &testTask).Return(nil)

	body, _ := json.Marshal(testTask)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/tasks", bytes.NewBuffer(body))

	handler.CreateTask(c)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestGetTaskHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := new(MockTaskRepository)
	handler := NewTaskHandler(mockRepo)
	testTime := time.Date(2025, time.July, 22, 12, 0, 0, 0, time.UTC)
	testTask := &domain.Task{
		ID:          1,
		Title:       "TitleTest",
		Description: "DescpTest",
		Done:        true,
		CreatedAt:   testTime,
	}

	mockRepo.On("GetTask", mock.Anything, 1).
		Return(testTask, nil).
		Once()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/tasks/1", nil)
	c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}

	handler.GetTaskHandler(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var reponseTask domain.Task

	err := json.Unmarshal(w.Body.Bytes(), &reponseTask)

	assert.NoError(t, err)
	assert.Equal(t, *testTask, reponseTask)
	mockRepo.AssertExpectations(t)
}

func TestUpdateTaskHandler_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockRepo := new(MockTaskRepository)
	handler := NewTaskHandler(mockRepo)

	testTime := time.Date(2025, time.July, 22, 12, 0, 0, 0, time.UTC)
	inputTask := domain.Task{
		Title:       "TestTitle",
		Description: "DescritTest",
		Done:        false,
		CreatedAt:   testTime,
	}

	mockRepo.On("UpdateTask", mock.Anything, mock.MatchedBy(func(task *domain.Task) bool {
		return task.Title == inputTask.Title &&
			task.Description == inputTask.Description &&
			task.Done == inputTask.Done &&
			task.CreatedAt.Equal(inputTask.CreatedAt)
	})).Return(nil).Once()

	body, _ := json.Marshal(inputTask)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("PUT", "/tasks/1", bytes.NewBuffer(body))
	c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}

	handler.UpdateTaskHandler(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var responseTask domain.Task
	err := json.Unmarshal(w.Body.Bytes(), &responseTask)
	assert.NoError(t, err)

	assert.Equal(t, 1, responseTask.ID)
	assert.Equal(t, inputTask.Title, responseTask.Title)
	assert.Equal(t, inputTask.Description, responseTask.Description)
	assert.Equal(t, inputTask.Done, responseTask.Done)
	assert.True(t, inputTask.CreatedAt.Equal(responseTask.CreatedAt))

	mockRepo.AssertExpectations(t)
}

func TestDeleteTaskHandler_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := new(MockTaskRepository)
	handler := NewTaskHandler(mockRepo)

	mockRepo.On("DeleteTask", mock.Anything, 1).Return(nil).Once()

	r := gin.Default()
	r.DELETE("/tasks/:id", handler.DeleteTaskHandler)

	req, _ := http.NewRequest("DELETE", "/tasks/1", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Empty(t, w.Body.String())
	mockRepo.AssertExpectations(t)
}
