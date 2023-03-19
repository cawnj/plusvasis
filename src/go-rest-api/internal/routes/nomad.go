package routes

import (
	"context"
	"continens/internal/controllers/nomad"
	"continens/internal/fauth"
	"fmt"
	"log"

	"firebase.google.com/go/auth"

	"github.com/labstack/echo/v4"
)

func isAuthorised(next echo.HandlerFunc, client *auth.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Header["Token"] != nil {

			token, err := client.VerifyIDToken(context.Background(), c.Request().Header["Token"][0])
			if err != nil {
				fmt.Println("error verifying ID token: ", err)
				return echo.ErrUnauthorized
			}

			fmt.Println("Verified ID token: ", token)
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
