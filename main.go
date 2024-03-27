package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"main.main/src/db"
	"main.main/src/route"
)

func main() {
	data, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatalf("error when open db: %s", err)
	}
	defer data.Close()

	err = db.Init(data)
	if err != nil {
		log.Fatalf("error when initialize db: %s", err)
	}

	engine := gin.Default()
	route.RegisterRouter(engine)

	srv := &http.Server{
		Addr:    ":8787",
		Handler: engine,
	}

	go func() {
		// Wait for interrupt signal to gracefully shutdown the server with
		// a timeout of 5 seconds.
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		log.Println("Shutdown Server...")

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Fatal("Server Shutdown:", err)
		}
		close(quit)
	}()

	// service connections
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}

	log.Println("Server exiting")

}
