package handler

import (
	"net/http"

	"short-url/interfaces"
	"short-url/models/api"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type URLHandler struct {
	urlService interfaces.IUrlService
}

func NewURLHandler(urlService interfaces.IUrlService) *URLHandler {
	handler := &URLHandler{
		urlService: urlService,
	}

	return handler
}

func (urlHandler *URLHandler) CreateURL(c echo.Context) error {
	var request api.URLCreate
	if err := c.Bind(&request); err != nil {
		log.Errorf("error parsing request, error is %v agent is %v and app id is %s", err, c.Request().UserAgent())

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	response := urlHandler.urlService.Create(c.Request().Context(), &request)
	if response.Error != "" {
		return echo.NewHTTPError(http.StatusInternalServerError, response)
	}

	return c.JSON(http.StatusCreated, response)
}

func (urlHandler *URLHandler) RedirectURL(c echo.Context) error {
	url, err := urlHandler.urlService.GetOriginal(c.Request().Context(), c.Param("shortURL"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.Redirect(http.StatusMovedPermanently, url)
}

func (urlHandler *URLHandler) DeleteURL(c echo.Context) error {
	err := urlHandler.urlService.Delete(c.Request().Context(), c.Param("shortURL"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.NoContent(200)
}