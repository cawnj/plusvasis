package routes

import (
	"continens/internal/controllers/nomad"
	"continens/internal/fauth"
	"continens/internal/middleware"
	"log"

	"github.com/labstack/echo/v4"
)

func NomadJobs(e *echo.Echo) {

	client, err := fauth.InitAuth()
	if err != nil {
		log.Fatalln("failed to init firebase auth", err)
	}

	e.GET("/jobs", middleware.IsAuthorised(nomad.GetJobs, client))
	e.POST("/jobs", middleware.IsAuthorised(nomad.CreateJob, client))
	e.GET("/job/:id", middleware.IsAuthorised(nomad.ReadJob, client))
	e.DELETE("job/:id", middleware.IsAuthorised(nomad.StopJob, client))
	e.POST("/job/:id", middleware.IsAuthorised(nomad.UpdateJob, client))
	e.GET("/job/:id/allocations", middleware.IsAuthorised(nomad.ReadJobAllocs, client))
	e.GET("/job/:id/alloc", middleware.IsAuthorised(nomad.ReadJobAlloc, client))
}
