package main

import (
	"os"

	"plusvasis/internal/middleware/firebase"
	"plusvasis/internal/middleware/logger"
	"plusvasis/internal/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:5173",
			"https://*.plusvasis.xyz",
		},
		AllowHeaders: []string{"*"},
	}))

	e.Use(logger.CustomLogger())
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))
	e.Use(firebase.Auth())

	routes.HealthRoutes(e)
	routes.NomadJobs(e)
	routes.NomadProxy(e)

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
