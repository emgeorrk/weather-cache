package gismeteo

import (
	"weather-cache/config"
	"weather-cache/internal/domain"
	"weather-cache/pkg/logger"
)

type API struct {
	log    logger.Logger
	URL    string
	APIKey string
}

func New(log logger.Logger, config config.Config) domain.WeatherProvider {
	return API{
		log:    log,
		URL:    config.WeatherAPIURL,
		APIKey: config.WeatherAPIKey,
	}
}
