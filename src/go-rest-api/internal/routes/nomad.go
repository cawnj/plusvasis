package routes

import (
	"plusvasis/internal/controllers/nomad"

	"github.com/labstack/echo/v4"
)

func NomadJobs(e *echo.Echo) {
	controller := nomad.NomadController{Client: &nomad.DefaultNomadClient{}}

	e.GET("/jobs", controller.GetJobs)
	e.POST("/jobs", controller.CreateJob)
	e.GET("/job/:id", controller.ReadJob)
	e.DELETE("job/:id", controller.StopJob)
	e.POST("/job/:id", controller.UpdateJob)
	e.POST("/job/:id/restart", controller.RestartJob)
	e.GET("/job/:id/start", controller.StartJob)
}
