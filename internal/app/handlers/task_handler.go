package handlers

import (
	"ToDoRestApi/internal/domain"
	"ToDoRestApi/internal/domain/infrastructure/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	taskRepo *repositories.TaskRepository
}

func NewTaskHandler(taskRepo *repositories.TaskRepository) *TaskHandler {
	return &TaskHandler{taskRepo: taskRepo}
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var task domain.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if task.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title is required"})
		return
	}
	if err := h.taskRepo.CreateTask(c.Request.Context(), &task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, task)
}
