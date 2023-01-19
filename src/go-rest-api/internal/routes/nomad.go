package routes

import (
	"continens/internal/controllers/nomad"

	"github.com/labstack/echo/v4"
)

func NomadJobs(e *echo.Echo) {

	e.GET("/list-jobs", nomad.GetJobs)
}
