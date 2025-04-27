package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"task-dependency-manager/internal/handlers"
)

func main() {
	r := gin.Default()

	r.GET("/health", handlers.Health)
	r.GET("/version", handlers.Version)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		panic(err)
	}
}
