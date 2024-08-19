package app

import (
	"encoding/json"
	"net/http"
)

type ContextMock struct {
	Context
	Params   map[string]string
	Err      error
	BindData any
	Status   int
	Response any
	Message  string
}

func (c *ContextMock) Param(key string) string {
	return c.Params[key]
}

func (c *ContextMock) Bind(obj any) error {
	if c.Err != nil {
		return c.Err
	}

	b, err := json.Marshal(c.BindData)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, obj)
	if err != nil {
		return err
	}
	return err
}

func (c *ContextMock) OK(obj any) {
	c.Status = http.StatusOK
	c.Response = obj
}

func (c *ContextMock) Created(obj any) {
	c.Status = http.StatusCreated
	c.Response = obj
}

func (c *ContextMock) InternalServer(message string) {
	c.Status = http.StatusInternalServerError
	c.Message = message
}

func (c *ContextMock) BadRequest(message string) {
	c.Status = http.StatusBadRequest
	c.Message = message
}
