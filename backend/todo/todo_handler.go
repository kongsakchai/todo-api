package todo

import (
	"strconv"
	"todo-api/app"
)

type Storager interface {
	Todos() ([]Todo, error)
	Todo(id int64) (Todo, error)
	Create(todo Todo) error
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

func (h *handler) Todos(app app.Context) {
	todos, err := h.storage.Todos()
	if err != nil {
		app.InternalServer(err.Error())
		return
	}

	app.OK(todos)
}

func (h *handler) Todo(app app.Context) {
	idStr := app.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		app.BadRequest("Invalid ID")
		return
	}

	todo, err := h.storage.Todo(id)
	if err != nil {
		app.InternalServer(err.Error())
		return
	}

	app.OK(todo)
}

func (h *handler) Create(app app.Context) {
	var todo Todo
	if err := app.Bind(&todo); err != nil {
		app.BadRequest("Invalid request body")
		return
	}

	err := h.storage.Create(todo)
	if err != nil {
		app.InternalServer(err.Error())
		return
	}

	app.Created(todo)
}

func (h *handler) Update(app app.Context) {
	idStr := app.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		app.BadRequest("Invalid ID")
		return
	}

	var todo Todo
	if err := app.Bind(&todo); err != nil {
		app.BadRequest("Invalid request body")
		return
	}

	todo.ID = id
	err = h.storage.Update(todo)
	if err != nil {
		app.InternalServer(err.Error())
		return
	}

	app.OK(todo)
}

func (h *handler) Delete(app app.Context) {
	idStr := app.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		app.BadRequest("Invalid ID")
		return
	}

	err = h.storage.Delete(id)
	if err != nil {
		app.InternalServer(err.Error())
		return
	}

	app.OK(nil)
}
