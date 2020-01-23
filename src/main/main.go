package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"main.main/src/route"
)

func main() {
	engine := gin.Default()
	engine.POST("/api/mlt", route.MltRequestHandler)
	route.RegisterRouter(engine)
	engine.Run(":8787")
}
