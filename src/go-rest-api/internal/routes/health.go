package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Health(e *echo.Echo) {

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{ Status: "OK" })
	})
}
