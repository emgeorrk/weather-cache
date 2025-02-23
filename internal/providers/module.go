package providers

import (
	"go.uber.org/fx"
	"weather-cache/config"
	"weather-cache/internal/constants"
	"weather-cache/internal/domain"
	"weather-cache/internal/providers/geo/starline_maps"
	"weather-cache/internal/providers/weather/gismeteo"
	"weather-cache/pkg/logger"
)

var Module = fx.Options(
	fx.Provide(NewWeatherProvider),
	fx.Provide(NewMapsProvider),
)

func NewWeatherProvider(log logger.Logger, config *config.Config) (domain.WeatherProvider, error) {
	switch config.Weather.APIType {
	case constants.GismeteoProvider:
		return gismeteo.New(log, config), nil
	default:
		return nil, constants.ErrWeatherAPITypeNotSupported
	}
}

func NewMapsProvider(log logger.Logger, config *config.Config) (domain.MapsProvider, error) {
	switch config.Maps.APIType {
	case constants.StarLineMapsProvider:
		return starline_maps.New(log, config), nil
	default:
		return nil, constants.ErrMapsAPITypeNotSupported
	}
}
