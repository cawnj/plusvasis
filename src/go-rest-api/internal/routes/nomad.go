package routes

import (
	"context"
	"continens/internal/controllers/nomad"
	"continens/internal/fauth"
	"log"
	"strings"

	"firebase.google.com/go/auth"

	"github.com/labstack/echo/v4"
)

func isAuthorised(next echo.HandlerFunc, client *auth.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Header["Authorization"] != nil {

			token := strings.Split(c.Request().Header["Authorization"][0], "Bearer ")[1]
			user, err := client.VerifyIDToken(context.Background(), token)
			if err != nil {
				return echo.ErrUnauthorized
			}

			c.Set("uid", user.UID)
			return next(c)
		} else {
			return echo.ErrUnauthorized
		}
	}
}

func NomadJobs(e *echo.Echo) {

	client, err := fauth.InitAuth()
	if err != nil {
		log.Fatalln("failed to init firebase auth", err)
	}

	e.GET("/jobs", isAuthorised(nomad.GetJobs, client))
	e.POST("/jobs", isAuthorised(nomad.CreateJob, client))
	e.GET("/job/:id", isAuthorised(nomad.ReadJob, client))
	e.DELETE("job/:id", isAuthorised(nomad.StopJob, client))
	e.POST("/job/:id", isAuthorised(nomad.UpdateJob, client))
	e.GET("/job/:id/allocations", isAuthorised(nomad.ReadJobAllocs, client))
}
