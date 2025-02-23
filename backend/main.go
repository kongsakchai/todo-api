package main

import (
	"fmt"
	"todo-api/app"
	"todo-api/config"
	"todo-api/database"
	"todo-api/todo"
)

func main() {
	cfg := config.Get()
	db, close := database.NewPostgres(cfg)
	defer close()

	s := todo.NewStorage(db)
	h := todo.NewHandler(s)

	r := app.NewGinRouter()
	r.GET("/todos", h.Todos)
	r.GET("/todos/:id", h.Todo)
	r.POST("/todos", h.Create)
	r.PUT("/todos/:id", h.Update)
	r.DELETE("/todos/:id", h.Delete)

	r.Run(fmt.Sprintf(":%s", cfg.Port))
}
