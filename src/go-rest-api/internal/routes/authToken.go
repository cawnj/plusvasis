package routes

import (
	auth "continens/internal/controllers/authToken"

	"github.com/labstack/echo/v4"
)

func AuthToken(e *echo.Echo) {

	e.POST("/token", auth.GetToken)
}
