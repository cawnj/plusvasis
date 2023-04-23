package main

import (
	"fmt"
	"os"
	"time"

	"plusvasis/internal/middleware/firebase"
	"plusvasis/internal/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

type CustomFormatter struct{}

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := entry.Time.Format(time.RFC3339)
	return []byte(fmt.Sprintf("[%s] method=%s, uri=%s, status=%d, user=%s, error=%v\n",
		timestamp,
		entry.Data["method"],
		entry.Data["uri"],
		entry.Data["status"],
		entry.Data["user"],
		entry.Data["error"],
	)), nil
}

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:5173",
			"https://*.plusvasis.xyz",
		},
		AllowHeaders: []string{"*"},
	}))

	log := logrus.New()
	log.SetFormatter(&CustomFormatter{})
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogMethod: true,
		LogError:  true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			uid := c.Get("uid")
			if uid != nil {
				uid = uid.(string)
			} else {
				uid = "nil"
			}
			log.WithFields(logrus.Fields{
				"method": values.Method,
				"uri":    values.URI,
				"status": values.Status,
				"user":   uid,
				"error":  values.Error,
			}).Info()
			return nil
		},
	}))

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
