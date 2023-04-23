package middleware

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

type Formatter struct{}

func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := entry.Time.Format(time.RFC3339)
	return []byte(fmt.Sprintf("[%s] method=%s, uri=%s, ip=%s, status=%d, user=%v, error=%v\n",
		timestamp,
		entry.Data["method"],
		entry.Data["uri"],
		entry.Data["ip"],
		entry.Data["status"],
		entry.Data["user"],
		entry.Data["error"],
	)), nil
}

var (
	log                        = logrus.New()
	DefaultRequestLoggerConfig = middleware.RequestLoggerConfig{
		LogMethod:   true,
		LogURI:      true,
		LogRemoteIP: true,
		LogStatus:   true,
		LogError:    true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			log.WithFields(logrus.Fields{
				"method": values.Method,
				"uri":    values.URI,
				"ip":     values.RemoteIP,
				"status": values.Status,
				"user":   c.Get("uid"),
				"error":  values.Error,
			}).Info()
			return nil
		},
	}
)

func Logger() echo.MiddlewareFunc {
	log.SetFormatter(&Formatter{})
	c := DefaultRequestLoggerConfig
	return LoggerWithConfig(c)
}

func LoggerWithConfig(config middleware.RequestLoggerConfig) echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(DefaultRequestLoggerConfig)
}
