package routes

import (
	"plusvasis/internal/controllers/nomad"
	"plusvasis/internal/controllers/websocket"

	"github.com/labstack/echo/v4"
)

func Websocket(e *echo.Echo) {
	controller := websocket.WebsocketController{NomadClient: &nomad.DefaultNomadClient{}}

	e.GET("/alloc/:id/exec", controller.AllocExec)
}
