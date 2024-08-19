package todo

import (
	"errors"
	"testing"
	"todo-api/app"
)

type mockStorage struct {
	Storager
	err  error
	todo Todo
}

func (m mockStorage) Todos() ([]Todo, error) {
	return []Todo{m.todo}, m.err
}

func (m mockStorage) Todo(id int64) (Todo, error) {
	return m.todo, m.err
}

func (m mockStorage) Create(todo Todo) error {
	return m.err
}

func (m mockStorage) Update(todo Todo) error {
	return m.err
}

func (m mockStorage) Delete(id int64) error {
	return m.err
}

func TestGetTodos(t *testing.T) {
	// Arrange
	repo := &mockStorage{}
	repo.todo = Todo{ID: 1, Title: "Title 1", Description: "Description 1", Done: false}
	ctx := app.ContextMock{}
	handler := NewHandler(repo)

	// Act
	handler.Todos(&ctx)

	// Assert
	if ctx.Status != 200 {
		t.Errorf("Expected status 200 but got %d", ctx.Status)
	}
}

func TestGetTodosError(t *testing.T) {
	// Arrange
	repo := &mockStorage{}
	repo.err = errors.New("error")
	ctx := app.ContextMock{}
	handler := NewHandler(repo)

	// Act
	handler.Todos(&ctx)

	// Assert
	if ctx.Status != 500 {
		t.Errorf("Expected status 500 but got %d", ctx.Status)
	}
}

func TestGetTodo(t *testing.T) {
	// Arrange
	repo := &mockStorage{}
	repo.todo = Todo{ID: 1, Title: "Title 1", Description: "Description 1", Done: false}
	ctx := app.ContextMock{}
	ctx.Params = map[string]string{"id": "1"}
	handler := NewHandler(repo)

	// Act
	handler.Todo(&ctx)

	// Assert
	if ctx.Status != 200 {
		t.Errorf("Expected status 200 but got %d", ctx.Status)
	}
}

func TestGetTodoError(t *testing.T) {
	// Arrange
	repo := &mockStorage{}
	repo.err = errors.New("error")
	ctx := app.ContextMock{}
	ctx.Params = map[string]string{"id": "1"}
	handler := NewHandler(repo)

	// Act
	handler.Todo(&ctx)

	// Assert
	if ctx.Status != 500 {
		t.Errorf("Expected status 500 but got %d", ctx.Status)
	}
}

func TestGetTodoBadRequest(t *testing.T) {
	// Arrange
	repo := &mockStorage{}
	ctx := app.ContextMock{}
	ctx.Params = map[string]string{"id": "a"}
	handler := NewHandler(repo)

	// Act
	handler.Todo(&ctx)

	// Assert
	if ctx.Status != 400 {
		t.Errorf("Expected status 400 but got %d", ctx.Status)
	}
}

func TestCreateTodo(t *testing.T) {
	// Arrange
	repo := &mockStorage{}
	ctx := app.ContextMock{}
	ctx.BindData = Todo{Title: "Title 1", Description: "Description 1", Done: false}
	handler := NewHandler(repo)

	// Act
	handler.Create(&ctx)

	// Assert
	if ctx.Status != 201 {
		t.Errorf("Expected status 201 but got %d", ctx.Status)
	}
}

func TestCreateTodoError(t *testing.T) {
	// Arrange
	repo := &mockStorage{}
	repo.err = errors.New("error")
	ctx := app.ContextMock{}
	ctx.BindData = Todo{Title: "Title 1", Description: "Description 1", Done: false}
	handler := NewHandler(repo)

	// Act
	handler.Create(&ctx)

	// Assert
	if ctx.Status != 500 {
		t.Errorf("Expected status 500 but got %d", ctx.Status)
	}
}

func TestCreateTodoBadRequest(t *testing.T) {
	// Arrange
	repo := &mockStorage{}
	ctx := app.ContextMock{}
	ctx.BindData = "invalid"
	handler := NewHandler(repo)

	// Act
	handler.Create(&ctx)

	// Assert
	if ctx.Status != 400 {
		t.Errorf("Expected status 400 but got %d", ctx.Status)
	}
}

func TestUpdateTodo(t *testing.T) {
	// Arrange
	repo := &mockStorage{}
	ctx := app.ContextMock{}
	ctx.BindData = Todo{ID: 1, Title: "Title 1", Description: "Description 1", Done: false}
	ctx.Params = map[string]string{"id": "1"}
	handler := NewHandler(repo)

	// Act
	handler.Update(&ctx)

	// Assert
	if ctx.Status != 200 {
		t.Errorf("Expected status 200 but got %d", ctx.Status)
	}
}

func TestUpdateTodoError(t *testing.T) {
	// Arrange
	repo := &mockStorage{}
	repo.err = errors.New("error")
	ctx := app.ContextMock{}
	ctx.BindData = Todo{ID: 1, Title: "Title 1", Description: "Description 1", Done: false}
	ctx.Params = map[string]string{"id": "1"}
	handler := NewHandler(repo)

	// Act
	handler.Update(&ctx)

	// Assert
	if ctx.Status != 500 {
		t.Errorf("Expected status 500 but got %d", ctx.Status)
	}
}

func TestUpdateTodoBadRequest(t *testing.T) {
	// Arrange
	repo := &mockStorage{}
	ctx := app.ContextMock{}
	ctx.BindData = "invalid"
	ctx.Params = map[string]string{"id": "1"}
	handler := NewHandler(repo)

	// Act
	handler.Update(&ctx)

	// Assert
	if ctx.Status != 400 {
		t.Errorf("Expected status 400 but got %d", ctx.Status)
	}
}

func TestUpdateTodoBadRequestID(t *testing.T) {
	// Arrange
	repo := &mockStorage{}
	ctx := app.ContextMock{}
	ctx.BindData = Todo{ID: 1, Title: "Title 1", Description: "Description 1", Done: false}
	ctx.Params = map[string]string{"id": "a"}
	handler := NewHandler(repo)

	// Act
	handler.Update(&ctx)

	// Assert
	if ctx.Status != 400 {
		t.Errorf("Expected status 400 but got %d", ctx.Status)
	}
}

func TestDeleteTodo(t *testing.T) {
	// Arrange
	repo := &mockStorage{}
	ctx := app.ContextMock{}
	ctx.Params = map[string]string{"id": "1"}
	handler := NewHandler(repo)

	// Act
	handler.Delete(&ctx)

	// Assert
	if ctx.Status != 200 {
		t.Errorf("Expected status 200 but got %d", ctx.Status)
	}
}

func TestDeleteTodoError(t *testing.T) {
	// Arrange
	repo := &mockStorage{}
	repo.err = errors.New("error")
	ctx := app.ContextMock{}
	ctx.Params = map[string]string{"id": "1"}
	handler := NewHandler(repo)

	// Act
	handler.Delete(&ctx)

	// Assert
	if ctx.Status != 500 {
		t.Errorf("Expected status 500 but got %d", ctx.Status)
	}
}

func TestDeleteTodoBadRequest(t *testing.T) {
	// Arrange
	repo := &mockStorage{}
	ctx := app.ContextMock{}
	ctx.Params = map[string]string{"id": "a"}
	handler := NewHandler(repo)

	// Act
	handler.Delete(&ctx)

	// Assert
	if ctx.Status != 400 {
		t.Errorf("Expected status 400 but got %d", ctx.Status)
	}
}
