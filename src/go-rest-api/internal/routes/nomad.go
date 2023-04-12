package routes

import (
	"log"

	"plusvasis/internal/controllers/nomad"
	"plusvasis/internal/fauth"
	"plusvasis/internal/middleware"

	"github.com/labstack/echo/v4"
)

func NomadJobs(e *echo.Echo) {
	client, err := fauth.InitAuth()
	if err != nil {
		log.Fatalln("failed to init firebase auth", err)
	}

	controller := nomad.NomadController{Client: &nomad.DefaultNomadClient{}}

	e.GET("/jobs", middleware.IsAuthorised(controller.GetJobs, client))
	e.POST("/jobs", middleware.IsAuthorised(controller.CreateJob, client))
	e.GET("/job/:id", middleware.IsAuthorised(controller.ReadJob, client))
	e.DELETE("job/:id", middleware.IsAuthorised(controller.StopJob, client))
	e.POST("/job/:id", middleware.IsAuthorised(controller.UpdateJob, client))
	e.GET("/job/:id/allocations", middleware.IsAuthorised(controller.ReadJobAllocs, client))
	e.GET("/job/:id/alloc", middleware.IsAuthorised(controller.ReadJobAlloc, client))
	e.POST("/job/:id/:allocId/:task/restart", middleware.IsAuthorised(controller.RestartJob, client))
}
