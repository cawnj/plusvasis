package routes

import (
	"continens/internal/controllers/nomad"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func isAuthorised(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Header["Token"] != nil {

			token, err := jwt.Parse(c.Request().Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return []byte(os.Getenv("SAMPLE_KEY")), nil
			})

			if err != nil {
				return echo.ErrUnauthorized
			}

			if token.Valid {
				return next(c)
			}
		} else {

			return echo.ErrUnauthorized
		}
		return nil
	}
}

func NomadJobs(e *echo.Echo) {

	e.GET("/jobs", isAuthorised(nomad.GetJobs))
	e.POST("/jobs", isAuthorised(nomad.CreateJob))
	e.GET("/job/:id", isAuthorised(nomad.ReadJob))
	e.DELETE("job/:id", isAuthorised(nomad.StopJob))
	e.POST("/job/:id", isAuthorised(nomad.UpdateJob))
	e.GET("/job/:id/allocations", isAuthorised(nomad.ReadJobAllocs))
}
