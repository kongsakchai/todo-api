package todo

import "errors"

type Todo struct {
	ID          int64
	Title       string `json:"title" binding:"required"` // binding:"required" not able to test because it need gin context
	Description string
	Done        bool
}

var (
	ErrNotFound = errors.New("todo not found")
)
