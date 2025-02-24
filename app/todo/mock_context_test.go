package todo

import (
	"encoding/json"
	"net/http"
	"todo-api/app"
)

type mockContext[T any] struct {
	app.Context
	payload  string
	params   map[string]string
	status   int
	response T
	message  string
}

func (c *mockContext[T]) Param(key string) string {
	return c.params[key]
}

func (c *mockContext[T]) Bind(obj any) error {
	return json.Unmarshal([]byte(c.payload), obj)
}

func (c *mockContext[T]) OK(obj any) {
	c.status = http.StatusOK
	if obj != nil {
		c.response = obj.(T)
	}
}

func (c *mockContext[T]) Created(obj any) {
	c.status = http.StatusCreated
	if obj != nil {
		c.response = obj.(T)
	}
}

func (c *mockContext[T]) NotFound(err error) {
	c.status = http.StatusOK
	c.message = err.Error()
}

func (c *mockContext[T]) InternalServer(err error) {
	c.status = http.StatusInternalServerError
	c.message = err.Error()
}

func (c *mockContext[T]) BadRequest(err error) {
	c.status = http.StatusBadRequest
	c.message = err.Error()
}

func (c *mockContext[T]) JSON(code int, obj any) {
	c.status = code
	if obj != nil {
		c.response = obj.(T)
	}
}
