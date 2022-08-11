package component

import (
	"short-url/handler"

	"github.com/labstack/echo/v4"
)

type Health struct {
	handler *handler.HealthCheckHandler
	router  *echo.Group
}

func NewHealth(handler *handler.HealthCheckHandler, router *echo.Group) *Health {
	return &Health{
		handler,
		router,
	}
}

func (h *Health) Routes() {
	h.router.GET("/health", h.HandleHealth())
}

func (h *Health) HandleHealth() echo.HandlerFunc {
	return func(c echo.Context) error {
		return h.handler.Check(c)
	}
}
