package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Health(c echo.Context) error {
	return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
}

func HealthRoutes(e *echo.Echo) {
	e.GET("/health", Health)
}
