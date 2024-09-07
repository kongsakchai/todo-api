package todo

import (
	"errors"
	"strconv"
	"todo-api/app"
)

type Storager interface {
	Todos() ([]Todo, error)
	Todo(id int64) (Todo, error)
	Create(todo Todo) (int64, error)
	Update(todo Todo) error
	Delete(id int64) error
}

type handler struct {
	storage Storager
}

func NewHandler(repo Storager) *handler {
	return &handler{
		storage: repo,
	}
}

func (h *handler) Todos(c app.Context) {
	todos, err := h.storage.Todos()
	if err != nil {
		c.InternalServer(err)
		return
	}

	c.OK(todos)
}

func (h *handler) Todo(c app.Context) {
	id := c.Param("id")
	todoID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.BadRequest(errors.New("invalid id"))
		return
	}

	todo, err := h.storage.Todo(todoID)
	if err == ErrNotFound {
		c.NotFound(err)
		return
	} else if err != nil {
		c.InternalServer(err)
		return
	}

	c.OK(todo)
}

func (h *handler) Create(c app.Context) {
	var todo Todo
	if err := c.Bind(&todo); err != nil {
		c.BadRequest(err)
		return
	}

	var err error
	todo.ID, err = h.storage.Create(todo)
	if err != nil {
		c.InternalServer(err)
		return
	}

	c.Created(todo)
}

func (h *handler) Update(c app.Context) {
	id := c.Param("id")
	todoID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.BadRequest(errors.New("invalid id"))
		return
	}

	var todo Todo
	if err := c.Bind(&todo); err != nil {
		c.BadRequest(err)
		return
	}

	todo.ID = todoID
	if err = h.storage.Update(todo); err != nil {
		c.InternalServer(err)
		return
	}

	c.OK(todo)
}

func (h *handler) Delete(c app.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.BadRequest(err)
		return
	}

	err = h.storage.Delete(id)
	if err != nil {
		c.InternalServer(err)
		return
	}

	c.OK(nil)
}
