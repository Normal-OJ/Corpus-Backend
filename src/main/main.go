package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"main.main/src/db"
)

func main() {
	data, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		print("error when open db:", err.Error())
	}
	err = db.Init(data)
	if err != nil {
		print("error wheen initialize db:", err.Error())
	}
	/*
		engine := gin.Default()
		route.RegisterRouter(engine)
		engine.Run(":8787")*/
}
