package main

import (
	"fmt"
	"log"
	"os"

	"continens/internal/auth"
	"continens/internal/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client, err := auth.InitAuth()
	if err != nil {
		log.Fatalln("failed to init firebase auth", err)
	}

	fmt.Println(client)

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.Health(e)
	routes.NomadJobs(e)

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
