package server

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/fx"
	"net/http"
	"time"
	"weather-cache/config"
	"weather-cache/internal/api/handler"
	"weather-cache/pkg/logger"
)

func Serve(lc fx.Lifecycle, log logger.Logger, r handler.RequestHandler, config *config.Config) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Info("Starting the application ☀️")

			go func() {
				err := r.Start(fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port))
				if err != nil && !errors.Is(err, http.ErrServerClosed) {
					log.Error("Failed to start the server", log.Err(err))
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("Shutting down server ❄️")
			ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
			defer cancel()
			return r.Shutdown(ctx)
		},
	})
}
