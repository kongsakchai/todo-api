package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"todo-api/app"
	"todo-api/app/todo"
	"todo-api/config"
	"todo-api/database"
)

func main() {
	cfg := config.Get()
	db, close := database.NewPostgres(cfg)
	defer close()

	r := app.NewGinRouter()

	{
		s := todo.NewStorage(db)
		h := todo.NewHandler(s)
		r.GET("/todos", h.Todos)
		r.GET("/todos/:id", h.Todo)
		r.POST("/todos", h.Create)
		r.PUT("/todos/:id", h.Update)
		r.DELETE("/todos/:id", h.Delete)
	}

	serv := http.Server{
		Addr:              fmt.Sprintf(":%s", cfg.Port),
		Handler:           r,
		ReadHeaderTimeout: 5 * time.Second,
	}

	idleConnsClosed := make(chan struct{})
	go gracefulShutdown(&serv, close, idleConnsClosed)

	fmt.Println("Server started at", cfg.Port)
	if err := serv.ListenAndServe(); err != http.ErrServerClosed {
		log.Println(err)
	}

	<-idleConnsClosed
}

func gracefulShutdown(serv *http.Server, dbClose func(), idle chan struct{}) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()
	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := serv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
	dbClose()

	fmt.Println("Server shutdown gracefully")
	close(idle)
}
