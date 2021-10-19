package server

import (
	"github.com/gin-gonic/gin"
)

func Start() error {
	r := gin.Default()
	registerRoutes(r)
	return r.Run(":8080")
}
