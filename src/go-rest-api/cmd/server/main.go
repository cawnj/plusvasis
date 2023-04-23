package main

import (
	"os"

	customMw "plusvasis/internal/middleware"
	"plusvasis/internal/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "plusvasis/docs"
)

//	@title			PlusVasis API
//	@version		1.0
//	@description	API backend for PlusVasis service

//	@host		api.plusvasis.xyz
//	@schemes	https

//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization

func setupMiddlewares(e *echo.Echo) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:*",
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

	e.GET("/swagger/*", echoSwagger.WrapHandler)
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
