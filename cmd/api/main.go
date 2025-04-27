package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/health", handleHealth)
	router.GET("/version", handleVersion)

	err := router.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}

func handleHealth(c *gin.Context) {
	c.JSON(200, gin.H{"status": "OK"})
}

func handleVersion(c *gin.Context) {
	c.JSON(200, gin.H{"version": "1.0.0"})
}
