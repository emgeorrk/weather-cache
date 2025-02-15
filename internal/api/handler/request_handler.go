package handler

import (
	"github.com/labstack/echo"
	"go.uber.org/fx"
	"strings"
	"weather-cache/config"
	"weather-cache/pkg/logger"
)

var Module = fx.Options(
	fx.Provide(NewRequestHandler),
)

type RequestHandler struct {
	*echo.Echo
}

func NewRequestHandler(lg logger.Logger, config config.Config) RequestHandler {
	e := echo.New()

	e.Use(logger.EchoMiddleware(lg))
	e.Logger.SetOutput(&lg)

	e.HideBanner = true
	e.HidePort = true
	e.Debug = strings.ToLower(config.Env) == "dev"

	return RequestHandler{Echo: e}
}
