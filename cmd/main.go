package main

import (
	"ToDoRestApi/internal/app/handlers"
	"ToDoRestApi/internal/domain/infrastructure/database"
	"ToDoRestApi/internal/domain/infrastructure/repositories"
	"log"

	_ "ToDoRestApi/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // Добавь эту строку
	ginSwagger "github.com/swaggo/gin-swagger" // Этот у тебя уже есть
)

func main() {
	if err := database.Init(); err != nil {
		log.Fatalf("Ошибка подключения к бд: %v", err)
	}

	taskRepo := repositories.NewTaskRepository(database.DB)
	taskHandler := handlers.NewTaskHandler(taskRepo)

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/", handlers.IndexPage)

	r.POST("/tasks", taskHandler.CreateTask)
	r.Run(":8080")
}
