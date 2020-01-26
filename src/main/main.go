package main

import (
	"github.com/gin-gonic/gin"
	"main.main/src/route"
	"main.main/src/view"
)

func main() {
	engine := gin.Default()
	engine.POST("/api/mlt", route.MltRequestHandler)
	engine.GET("/api/view", view.RequestHandler)
	engine.Run(":8787")
}
