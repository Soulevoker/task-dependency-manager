package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"task-dependency-manager/internal/handlers"
	"task-dependency-manager/internal/services"
	"task-dependency-manager/internal/storage"
)

func main() {
	r := gin.Default()
	// Base Routes
	r.GET("/health", handlers.Health)
	r.GET("/version", handlers.Version)

	// Task Routes
	store := storage.NewInMemoryStore()
	tasksService := services.NewTaskService(store)
	// Seed Tasks for local dev
	if gin.Mode() == gin.DebugMode {
		fmt.Println("[GIN-debug] Seeding Tasks For Local Development")
		if err := tasksService.SeedTasks(); err != nil {
			panic(err)
		}
	}
	tasksHandler := handlers.NewTaskHandler(tasksService)
	r.GET("/tasks/:id", tasksHandler.GetTask)
	r.DELETE("/tasks/:id", tasksHandler.DeleteTask)
	r.POST("/tasks/:id", tasksHandler.UpdateTask)
	r.POST("/tasks", tasksHandler.CreateTask)
	r.GET("/tasks", tasksHandler.ListTasks)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		panic(err)
	}
}
