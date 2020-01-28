package main

import (
	"github.com/gin-gonic/gin"
	"main.main/src/route"
)

func main() {
	engine := gin.Default()
	route.RegisterRouter(engine)
	engine.Run(":8787")
}
