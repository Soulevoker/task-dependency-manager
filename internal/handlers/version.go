package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Version(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"version": "1.0.0"})
}
