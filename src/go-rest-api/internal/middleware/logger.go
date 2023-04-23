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
	return []byte(fmt.Sprintf("[%s] method=%s, uri=%s, status=%d, user=%s, error=%v\n",
		timestamp,
		entry.Data["method"],
		entry.Data["uri"],
		entry.Data["status"],
		entry.Data["user"],
		entry.Data["error"],
	)), nil
}

var log = logrus.New()
var DefaultRequestLoggerConfig = middleware.RequestLoggerConfig{
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
}

func Logger() echo.MiddlewareFunc {
	log.SetFormatter(&Formatter{})
	c := DefaultRequestLoggerConfig
	return LoggerWithConfig(c)
}

func LoggerWithConfig(config middleware.RequestLoggerConfig) echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(DefaultRequestLoggerConfig)
}
