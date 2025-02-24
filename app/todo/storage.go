package todo

import (
	"database/sql"
)

type storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *storage {
	return &storage{
		db: db,
	}
}

func (s *storage) Todos() ([]Todo, error) {
	rows, err := s.db.Query("SELECT id, title, description, done, created_at FROM todo ORDER BY created_at")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := []Todo{}
	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done, &todo.CreatedAt)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, err
}

func (s *storage) Todo(id int64) (Todo, error) {
	var todo Todo
	row := s.db.QueryRow("SELECT id, title, description, done, created_at FROM todo WHERE id = $1", id)
	err := row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done, &todo.CreatedAt)
	if err == sql.ErrNoRows {
		return todo, ErrNotFound
	} else if err != nil {
		return todo, err
	}

	return todo, nil
}

func (s *storage) Create(todo Todo) (int64, error) {
	query := `INSERT INTO todo (title, description, done) VALUES ($1, $2, $3) RETURNING id`
	var id int64
	row := s.db.QueryRow(query, todo.Title, todo.Description, todo.Done)
	err := row.Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (s *storage) Update(todo Todo) error {
	stmt, err := s.db.Prepare("UPDATE todo SET title = $1, description = $2, done = $3 WHERE id = $4")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(todo.Title, todo.Description, todo.Done, todo.ID)
	return err
}

func (s *storage) Delete(id int64) error {
	stmt, err := s.db.Prepare("DELETE FROM todo WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	return err
}
