package main

import (
	"context"
	"fmt"
	"go.uber.org/fx"
	"os"
	"os/signal"
	"syscall"
	"weather-cache/config"
	"weather-cache/internal/api/controller"
	"weather-cache/internal/api/handler"
	"weather-cache/internal/api/routes"
	"weather-cache/internal/cache"
	"weather-cache/internal/infrastructure/providers"
	"weather-cache/internal/server"
	"weather-cache/internal/services/maps"
	"weather-cache/internal/services/weather"
	"weather-cache/pkg/logger"
)

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	app := fx.New(
		config.Module,
		logger.Module,
		weather.Module,
		maps.Module,
		cache.Module,
		providers.Module,
		controller.Module,
		handler.Module,
		fx.Invoke(routes.Setup),
		fx.Invoke(server.Serve),
	)

	go func() {
		if err := app.Start(context.Background()); err != nil {
			fmt.Println("Error starting app:", err)
		}
	}()

	<-stop

	err := app.Stop(context.Background())
	if err != nil {
		fmt.Println("Error stopping app:", err)
	}
}
