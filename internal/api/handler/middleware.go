package handler

import (
	"context"
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
				logger.RequestID(c.Request().Context()),
			)

			return err
		}
	}
}

// RequestIDMiddleware creates a new middleware that generates a new request ID.
func RequestIDMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		requestID := uuid.New().String()

		ctx := context.WithValue(c.Request().Context(), constants.RequestID, requestID)

		c.Response().Header().Set("X-Request-ID", requestID)
		c.Set(constants.RequestID, requestID)
		c.SetRequest(c.Request().WithContext(ctx))

		return next(c)
	}
}
