package handler

import (
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"time"
	"weather-cache/internal/constants"
	"weather-cache/pkg/logger"
)

// LoggerMiddleware creates a new middleware that logs HTTP requests.
func LoggerMiddleware(logger logger.Logger) echo.MiddlewareFunc {
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
				logger.RequestID(c.Get(constants.RequestID).(string)),
			)

			return err
		}
	}
}

// RequestIDMiddleware creates a new middleware that generates a new request ID.
func RequestIDMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		requestID := uuid.New().String()

		c.Response().Header().Set("X-Request-ID", requestID)
		c.Set(constants.RequestID, requestID)

		return next(c)
	}
}
