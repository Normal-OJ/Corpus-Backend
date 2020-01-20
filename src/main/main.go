package main

import (
	"github.com/gin-gonic/gin"
	"main.main/src/route"
)

func main() {
	engine := gin.Default()
	engine.POST("/api/mlt", route.MltRequestHandler)
	engine.Run(":8787")
}
