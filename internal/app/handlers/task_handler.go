package handlers

import (
	"ToDoRestApi/internal/domain"
	"ToDoRestApi/internal/domain/infrastructure/repositories"
	"database/sql"
	"log"
	"net/http"
	"strconv"

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

// @Summary Получить задачу по ID
// @Description Получение задачи по её идентификатору
// @Tags task
// @Accept json
// @Produce json
// @Param id path int true "ID задачи"
// @Success 200 {object} domain.Task
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tasks/{id} [get]
func (h *TaskHandler) GetTaskHandler(c *gin.Context) {
	idStr := c.Param("id")

	taskId, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task"})
		return
	}

	task, err := h.taskRepo.GetTask(c.Request.Context(), taskId)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
			return
		}

		log.Printf("err get task %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

// @Summary Обновить задачу
// @Description Обновление задачи по ID
// @Tags task
// @Accept json
// @Produce json
// @Param id path int true "ID задачи"
// @Param task body domain.Task true "Новые данные задачи"
// @Success 200 {object} domain.Task
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tasks/{id} [put]
func (h *TaskHandler) UpdateTaskHandler(c *gin.Context) {
	idStr := c.Param("id")

	taskID, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
		return
	}
	var updatedTask domain.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTask.ID = taskID

	if err := h.taskRepo.UpdateTask(c.Request.Context(), &updatedTask); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedTask)
}

// @Summary Удалить задачу
// @Description Удаление задачи по ID
// @Tags task
// @Accept json
// @Produce json
// @Param id path int true "ID задачи"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tasks/{id} [delete]
func (h *TaskHandler) DeleteTaskHandler(c *gin.Context) {
	idSrt := c.Param("id")

	taskID, err := strconv.Atoi(idSrt)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
		return
	}

	if err := h.taskRepo.DeleteTask(c.Request.Context(), taskID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
