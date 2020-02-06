package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"main.main/src/db"
	"main.main/src/route"
)

func main() {
	data, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		print("error when open db:", err.Error())
	}
	err = db.Init(data)
	if err != nil {
		print("error when initialize db:", err.Error())
	}
	engine := gin.Default()
	route.RegisterRouter(engine)
	engine.Run(":8787")
}
