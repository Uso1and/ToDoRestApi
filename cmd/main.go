package main

import (
	"ToDoRestApi/internal/app/handlers"
	"ToDoRestApi/internal/domain/infrastructure/database"
	"ToDoRestApi/internal/domain/infrastructure/repositories"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := database.Init(); err != nil {
		log.Fatalf("Ошибка подключения к бд: %v", err)
	}

	taskRepo := repositories.NewTaskRepository(database.DB)
	taskHandler := handlers.NewTaskHandler(taskRepo)

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/", handlers.IndexPage)

	r.POST("/tasks", taskHandler.CreateTask)
	r.Run(":8080")
}
