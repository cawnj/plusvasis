package middleware

import (
	"context"
	"net/http"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
	"google.golang.org/api/option"
)

type Config struct {
	CredentialsFile string
	Skipper         middleware.Skipper
}

// disable auth for health check
func DefaultSkipper(c echo.Context) bool {
	return c.Path() == "/health"
}

var DefaultConfig = Config{
	CredentialsFile: "firebase.json",
	Skipper:         DefaultSkipper,
}

func Firebase() echo.MiddlewareFunc {
	c := DefaultConfig
	return FirebaseWithConfig(c)
}

func FirebaseWithConfig(config Config) echo.MiddlewareFunc {
	opt := option.WithCredentialsFile(config.CredentialsFile)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic(errors.Wrap(err, "error initializing app"))
	}
	client, err := app.Auth(context.Background())
	if err != nil {
		panic(errors.Wrap(err, "error getting auth client"))
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}

			token := ""
			if c.Request().Header["Authorization"] != nil {
				token = strings.Split(c.Request().Header["Authorization"][0], "Bearer ")[1]
			} else if c.QueryParam("access_token") != "" {
				token = c.QueryParam("access_token")
			} else {
				return echo.NewHTTPError(http.StatusUnauthorized, "missing auth token")
			}

			user, err := client.VerifyIDToken(context.Background(), token)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
			}

			c.Set("uid", user.UID)
			return next(c)
		}
	}
}
