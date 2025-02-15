package providers

import (
	"go.uber.org/fx"
	"weather-cache/config"
	"weather-cache/internal/constants"
	"weather-cache/internal/domain"
	"weather-cache/internal/infrastructure/providers/gismeteo"
	"weather-cache/internal/infrastructure/providers/starline_maps"
	"weather-cache/pkg/logger"
)

var Module = fx.Options(
	fx.Provide(NewWeatherProvider),
	fx.Provide(NewMapsProvider),
)

func NewWeatherProvider(log logger.Logger, config config.Config) (domain.WeatherProvider, error) {
	switch config.WeatherAPIType {
	case "gismeteo":
		return gismeteo.New(log, config), nil
	default:
		return nil, constants.ErrWeatherAPITypeNotSupported
	}
}

func NewMapsProvider(log logger.Logger, config config.Config) (domain.MapsProvider, error) {
	switch config.MapsAPIType {
	case "starline_maps":
		return starline_maps.New(log, config), nil
	default:
		return nil, constants.ErrMapsAPITypeNotSupported
	}
}
