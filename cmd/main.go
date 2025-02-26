package main

import (
	"context"
	"fmt"
	"go.uber.org/fx"
	"os"
	"os/signal"
	"syscall"
	"weather-cache/config"
	"weather-cache/internal/api"
	"weather-cache/internal/api/controller"
	"weather-cache/internal/api/handler"
	"weather-cache/internal/cache"
	"weather-cache/internal/providers"
	"weather-cache/internal/server"
	"weather-cache/internal/services/maps"
	"weather-cache/internal/services/weather"
	"weather-cache/pkg/logger"
)

var Module = fx.Options(
	config.Module,
	logger.Module,
	weather.Module,
	maps.Module,
	cache.Module,
	providers.Module,
	controller.Module,
	handler.Module,
)

var Invokes = fx.Invoke(
	api.Setup,
	server.Serve,
)

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	app := fx.New(
		Module,
		Invokes,
	)

	if err := app.Err(); err != nil {
		os.Exit(1)
	}

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
