package logger

import (
	"github.com/labstack/echo"
	"time"
)

// EchoMiddleware создает middleware для логирования запросов
func EchoMiddleware(logger Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			err := next(c)

			req := c.Request()
			res := c.Response()

			logger.Info("HTTP request",
				"method", req.Method,
				"path", req.URL.Path,
				"status", res.Status,
				"latency", time.Since(start),
				"remote_ip", c.RealIP(),
			)

			return err
		}
	}
}
