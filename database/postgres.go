package database

import (
	"database/sql"
	"log"
	"todo-api/config"

	_ "github.com/lib/pq"
)

func NewPostgres(cfg config.Config) (*sql.DB, func()) {
	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Connect to database error", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Ping to database error", err)
	}

	close := func() {
		_ = db.Close()
	}

	return db, close
}
