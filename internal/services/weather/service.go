package weather

import (
	"go.uber.org/fx"
	"weather-cache/internal/domain"
	"weather-cache/pkg/logger"
)

var Module = fx.Options(
	fx.Provide(New),
)

type Service struct {
	log             logger.Logger
	cache           domain.Cache
	mapsService     domain.MapsService
	weatherProvider domain.WeatherProvider
}

func New(log logger.Logger, cache domain.Cache, maps domain.MapsService) domain.WeatherService {
	return Service{
		log:         log,
		cache:       cache,
		mapsService: maps,
	}
}
