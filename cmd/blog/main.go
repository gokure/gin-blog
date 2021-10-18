package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gokure/gin-blog/internal/server"
)

func main() {
	r := gin.Default()
	r = server.InitRouter(r)
	r.Run(":8080")
}
