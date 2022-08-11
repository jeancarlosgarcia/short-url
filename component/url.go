package component

import (
	"short-url/handler"

	"github.com/labstack/echo/v4"
)

type Url struct {
	handler *handler.URLHandler
	router  *echo.Group
}

func NewUrl(handler *handler.URLHandler, router *echo.Group) *Url {
	return &Url{
		handler,
		router,
	}
}

func (e *Url) Routes() {
	e.router.POST("/", e.handler.CreateURL)
	e.router.GET("/:shortURL", e.handler.RedirectURL)
	e.router.DELETE("/:shortURL", e.handler.DeleteURL)
}