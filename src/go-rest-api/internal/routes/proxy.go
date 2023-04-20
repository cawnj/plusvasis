package routes

import (
	"plusvasis/internal/controllers/nomad"
	"plusvasis/internal/controllers/proxy"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

func NomadProxy(e *echo.Echo) {
	controller := proxy.NomadProxyController{
		Client: &nomad.DefaultNomadClient{},
		Dialer: websocket.DefaultDialer,
	}

	e.GET("/job/:id/exec", controller.AllocExec)
	e.GET("/job/:id/logs", controller.StreamLogs)
}
