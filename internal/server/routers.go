package server

import (
	"github.com/gin-gonic/gin"
	"github.com/gokure/gin-blog/internal/api"
)

func registerRoutes(engine *gin.Engine) *gin.Engine {
	route := engine.Group("/api")
	{
		route.GET("/ping", api.GetPing)
	}

	return engine
}
