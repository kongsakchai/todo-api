package todo

import (
	"encoding/json"
	"net/http"
	"todo-api/app"
)

type mockContext struct {
	app.Context
	payload  string
	params   map[string]string
	status   int
	response any
	message  string
}

func (c *mockContext) Param(key string) string {
	return c.params[key]
}

func (c *mockContext) Bind(obj any) error {
	return json.Unmarshal([]byte(c.payload), obj)
}

func (c *mockContext) OK(obj any) {
	c.status = http.StatusOK
	c.response = obj
}

func (c *mockContext) Created(obj any) {
	c.status = http.StatusCreated
	c.response = obj
}

func (c *mockContext) NotFound(err error) {
	c.status = http.StatusOK
	c.message = err.Error()
}

func (c *mockContext) InternalServer(err error) {
	c.status = http.StatusInternalServerError
	c.message = err.Error()
}

func (c *mockContext) BadRequest(err error) {
	c.status = http.StatusBadRequest
	c.message = err.Error()
}

func (c *mockContext) JSON(code int, obj any) {
	c.status = code
	c.response = obj
}
