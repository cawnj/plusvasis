package routes

import (
	"continens/internal/controllers/nomad"

	"github.com/labstack/echo/v4"
)

func NomadJobs(e *echo.Echo) {

	e.GET("/jobs", nomad.GetJobs)
	e.POST("/jobs", nomad.CreateJob)
	e.GET("/job/:id", nomad.ReadJob)
	e.DELETE("job/:id", nomad.StopJob)
	e.POST("/job/:id", nomad.UpdateJob)
}
