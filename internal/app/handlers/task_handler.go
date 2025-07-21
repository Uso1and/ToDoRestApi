package handlers

import (
	"ToDoRestApi/internal/domain"
	"ToDoRestApi/internal/domain/infrastructure/repositories"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	taskRepo *repositories.TaskRepository
}

func NewTaskHandler(taskRepo *repositories.TaskRepository) *TaskHandler {
	return &TaskHandler{taskRepo: taskRepo}
}

// @Summary Создание Таски
// @Description Добавление Задачи
// @Tags task
// @Accept json
// @Produce json
// @Param   task body domain.Task true "Данные задачи"
// @Example  { "title": "Пример задачи", "description": "Описание", "done": false, "created_at": "2025-07-22T12:00:00Z" }
// @Success 201 {object} domain.Task
// @Failure 400 {object} map[string]string
// @Router /tasks [post]
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
		log.Printf("Ошибка создания такси: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, task)
}
