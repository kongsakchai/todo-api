package app

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type ginContext struct {
	*gin.Context
}

func (c ginContext) Param(key string) string {
	return c.Context.Param(key)
}

func (c ginContext) Bind(obj any) error {
	return c.Context.ShouldBindJSON(obj)
}

func (c ginContext) OK(obj any) {
	c.JSON(http.StatusOK, response{
		Success: true,
		Data:    obj,
	})
}

func (c ginContext) Created(obj any) {
	c.JSON(http.StatusCreated, response{
		Success: true,
		Data:    obj,
	})
}

func (c ginContext) InternalServer(message string) {
	c.JSON(http.StatusInternalServerError, response{
		Success: false,
		Message: message,
	})
}

func (c ginContext) BadRequest(message string) {
	c.JSON(http.StatusBadRequest, response{
		Success: false,
		Message: message,
	})
}

func (c ginContext) JSON(code int, obj any) {
	c.Context.JSON(http.StatusUnauthorized, obj)
}

type ginRoute struct {
	*gin.Engine
}

func NewGinRouter() *ginRoute {
	r := gin.Default()

	config := cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"X-Requested-With", "Authorization", "Origin", "Content-Length", "Content-Type", "TransactionID"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}

	r.Use(cors.New(config))

	return &ginRoute{r}
}

func NewGinHandler(handler func(Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(&ginContext{c})
	}
}

func (r *ginRoute) GET(path string, handler func(Context)) {
	r.Engine.GET(path, NewGinHandler(handler))
}

func (r *ginRoute) POST(path string, handler func(Context)) {
	r.Engine.POST(path, NewGinHandler(handler))
}

func (r *ginRoute) PUT(path string, handler func(Context)) {
	r.Engine.PUT(path, NewGinHandler(handler))
}

func (r *ginRoute) DELETE(path string, handler func(Context)) {
	r.Engine.DELETE(path, NewGinHandler(handler))
}

func (r *ginRoute) PATCH(path string, handler func(Context)) {
	r.Engine.PATCH(path, NewGinHandler(handler))
}
