package todo

type Todo struct {
	ID          int64
	Title       string `json:"title" binding:"required"` // binding:"required" test on e2e
	Description string
	Done        bool
}
