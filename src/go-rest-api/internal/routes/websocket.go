package routes

import (
	"plusvasis/internal/controllers/websocket"

	"github.com/labstack/echo/v4"
)

func Websocket(e *echo.Echo) {
	e.GET("/alloc/:id/exec", websocket.AllocExec)
}
