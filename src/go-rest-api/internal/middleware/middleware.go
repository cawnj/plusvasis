package middleware

import (
	"context"
	"strings"

	"firebase.google.com/go/auth"
	"github.com/labstack/echo/v4"
)

func IsAuthorised(next echo.HandlerFunc, client *auth.Client) echo.HandlerFunc {
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
