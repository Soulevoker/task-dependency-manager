package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Health handles GET /health endpoint
func Health(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"status": "OK"})
}
