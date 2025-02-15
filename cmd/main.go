package main

import (
	"go.uber.org/fx"
	"weather-cache/config"
	"weather-cache/internal/api/controller"
	"weather-cache/internal/api/handler"
	"weather-cache/internal/cache"
	"weather-cache/internal/infrastructure/providers"
	"weather-cache/internal/services/maps"
	"weather-cache/internal/services/weather"
	"weather-cache/pkg/logger"
)

func main() {
	app := fx.New(
		config.Module,
		logger.Module,
		weather.Module,
		maps.Module,
		cache.Module,
		providers.Module,
		controller.Module,
		handler.Module,
		fx.Invoke(func(log logger.Logger, r handler.RequestHandler) {
			log.Info("Starting the application")

			err := r.Start("localhost:8080")
			if err != nil {
				log.Error("Failed to start the server", "error", err)
				return
			}
		}),
	)

	app.Run()
}
