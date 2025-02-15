package maps

import (
	"go.uber.org/fx"
	"weather-cache/internal/domain"
	"weather-cache/pkg/logger"
)

var Module = fx.Options(
	fx.Provide(New),
)

type Service struct {
	log          logger.Logger
	mapsProvider domain.MapsProvider
	URL          string
}

func New(log logger.Logger) domain.MapsService {
	return Service{
		log: log,
	}
}
