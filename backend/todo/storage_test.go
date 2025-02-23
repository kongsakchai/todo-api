package todo

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	_ "modernc.org/sqlite"
)

var todoTable string = `
CREATE TABLE IF NOT EXISTS todo (
    id Integer PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    done BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);`

var insertTodos = `
INSERT INTO todo (title, description, done)
VALUES
	("todo 1", "description 1", true),
	("todo 2", "description 2", false),
	("todo 3", "description 3", false);`

func TestGetTodosStorage(t *testing.T) {
	t.Run("should return all todos", func(t *testing.T) {
		// Arrange
		db, err := sql.Open("sqlite", "file:TestGetTodosStorage?cache=shared&mode=memory")
		if err != nil {
			t.Fatal(err)
		}
		defer db.Close()

		_, err = db.Exec(todoTable)
		assert.NoError(t, err)
		_, err = db.Exec(insertTodos)
		assert.NoError(t, err)

		expectedResp := []Todo{
			{ID: 1, Title: "todo 1", Description: "description 1", Done: true},
			{ID: 2, Title: "todo 2", Description: "description 2", Done: false},
			{ID: 3, Title: "todo 3", Description: "description 3", Done: false},
		}

		s := NewStorage(db)

		// Act
		resp, err := s.Todos()

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedResp, resp)
	})
}

func TestGetTodoStorage(t *testing.T) {
	t.Run("should return todo by id", func(t *testing.T) {
		// Arrange
		db, err := sql.Open("sqlite", "file:TestGetTodosStorage?cache=shared&mode=memory")
		if err != nil {
			t.Fatal(err)
		}
		defer db.Close()

		_, err = db.Exec(todoTable)
		assert.NoError(t, err)
		_, err = db.Exec(insertTodos)
		assert.NoError(t, err)

		expectedResp := Todo{ID: 1, Title: "todo 1", Description: "description 1", Done: true}

		s := NewStorage(db)

		// Act
		resp, err := s.Todo(1)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedResp, resp)
	})
}

func TestNotFoundGetTodoStorage(t *testing.T) {
	t.Run("should return error when data not found", func(t *testing.T) {
		// Arrange
		db, err := sql.Open("sqlite", "file:TestGetTodosStorage?cache=shared&mode=memory")
		if err != nil {
			t.Fatal(err)
		}
		defer db.Close()

		_, err = db.Exec(todoTable)
		assert.NoError(t, err)
		_, err = db.Exec(insertTodos)
		assert.NoError(t, err)

		s := NewStorage(db)

		// Act
		_, err = s.Todo(4)

		// Assert
		assert.Equal(t, ErrNotFound, err)
	})
}

func TestCreateTodoStorage(t *testing.T) {
	t.Run("should return id when create todo success", func(t *testing.T) {
		// Arrange
		db, err := sql.Open("sqlite", "file:TestGetTodosStorage?cache=shared&mode=memory")
		if err != nil {
			t.Fatal(err)
		}
		defer db.Close()

		_, err = db.Exec(todoTable)
		assert.NoError(t, err)
		_, err = db.Exec(insertTodos)
		assert.NoError(t, err)

		todo := Todo{Title: "todo", Description: "description", Done: true}

		s := NewStorage(db)

		// Act
		id, err := s.Create(todo)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, int64(4), id)

		var count int64
		row := db.QueryRow("select count(*) from todo")
		err = row.Scan(&count)

		assert.NoError(t, err)
		assert.Equal(t, int64(4), count)
	})
}

func TestUpdateTodoStorage(t *testing.T) {
	t.Run("should change name todo when update todo success", func(t *testing.T) {
		// Arrange
		db, err := sql.Open("sqlite", "file:TestGetTodosStorage?cache=shared&mode=memory")
		if err != nil {
			t.Fatal(err)
		}
		defer db.Close()

		_, err = db.Exec(todoTable)
		assert.NoError(t, err)
		_, err = db.Exec(insertTodos)
		assert.NoError(t, err)

		todo := Todo{ID: 2, Title: "todo", Description: "description", Done: true}

		s := NewStorage(db)

		// Act
		err = s.Update(todo)

		// Assert
		assert.NoError(t, err)

		var newTodo Todo
		row := db.QueryRow("SELECT id,title,description,done FROM todo WHERE id = $1", 2)
		err = row.Scan(&newTodo.ID, &newTodo.Title, &newTodo.Description, &newTodo.Done)

		assert.NoError(t, err)
		assert.Equal(t, todo, newTodo)
	})
}

func TestDeleteTodo(t *testing.T) {
	t.Run("should delete todo by id", func(t *testing.T) {
		// Arrange
		db, err := sql.Open("sqlite", "file:TestGetTodosStorage?cache=shared&mode=memory")
		if err != nil {
			t.Fatal(err)
		}
		defer db.Close()

		_, err = db.Exec(todoTable)
		assert.NoError(t, err)
		_, err = db.Exec(insertTodos)
		assert.NoError(t, err)

		s := NewStorage(db)

		// Act
		err = s.Delete(2)

		// Assert
		assert.NoError(t, err)

		var count int64
		row := db.QueryRow("select count(*) from todo")
		err = row.Scan(&count)

		assert.NoError(t, err)
		assert.Equal(t, int64(2), count)
	})
}
