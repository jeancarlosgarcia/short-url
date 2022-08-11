package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// HealthCheckHandler is a struct to bind the Check method in a request
type HealthCheckHandler struct {
}

// NewHealthCheckHandler creates a HealthCheckHandler struct
func NewHealthCheckHandler() *HealthCheckHandler {
	return &HealthCheckHandler{}
}

// Check returns a no content http OK request
func (hch *HealthCheckHandler) Check(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
