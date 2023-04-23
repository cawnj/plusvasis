package main

import (
	"os"

	customMw "plusvasis/internal/middleware"
	"plusvasis/internal/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func setupMiddlewares(e *echo.Echo) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:5173",
			"https://*.plusvasis.xyz",
		},
		AllowHeaders: []string{"*"},
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))

	e.Use(customMw.Logger())
	e.Use(customMw.Firebase())
}

func setupRoutes(e *echo.Echo) {
	routes.HealthRoutes(e)
	routes.NomadJobs(e)
	routes.NomadProxy(e)
}

func main() {
	e := echo.New()

	setupMiddlewares(e)
	setupRoutes(e)

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
