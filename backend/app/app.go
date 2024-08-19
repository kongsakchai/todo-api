package app

type Context interface {
	Param(key string) string
	Bind(obj any) error
	OK(obj any)
	Created(obj any)
	InternalServer(message string)
	BadRequest(message string)
	JSON(code int, obj any)
}

type Router interface {
	GET(path string, handler func(c Context))
	POST(path string, handler func(c Context))
	PUT(path string, handler func(c Context))
	DELETE(path string, handler func(c Context))
	PATCH(path string, handler func(c Context))
}

type response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type app struct {
	Router Router
}

func New(router Router) *app {
	return &app{Router: router}
}

func NewGin() *app {
	return New(NewGinRouter())
}
