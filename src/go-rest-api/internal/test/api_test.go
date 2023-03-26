package test

import (
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/steinfletcher/apitest"
)

func testApp() *echo.Echo {
	app := echo.New()
	app.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})
	/*	app.GET("/jobs", nomad.GetJobs)
		app.POST("/jobs", nomad.CreateJob)
		app.GET("/job/:id", nomad.ReadJob)
		app.DELETE("job/:id", nomad.StopJob)
		app.POST("/job/:id", nomad.UpdateJob)
		app.GET("/job/:id/allocations", nomad.ReadJobAllocs)
		app.GET("/job/:id/alloc", nomad.ReadJobAlloc) */
	return app
}

func TestHealth(t *testing.T) {
	apitest.New().
		Report(apitest.SequenceDiagram()).
		Handler(testApp()).
		Get("/health").
		Expect(t).
		Status(http.StatusOK).
		End()
}
