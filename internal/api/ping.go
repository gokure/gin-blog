package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GET /api/ping
func GetPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "PONG"})
}
