package starline_maps

import (
	"weather-cache/config"
	"weather-cache/internal/constants"
	"weather-cache/internal/domain"
	"weather-cache/pkg/logger"
)

type API struct {
	log    logger.Logger
	URL    string
	APIKey string
}

func New(log logger.Logger, config *config.Config) domain.MapsProvider {
	return API{
		log:    log,
		URL:    config.Maps.APIs[constants.StarLineMapsProvider].URL,
		APIKey: config.Maps.APIs[constants.StarLineMapsProvider].Key,
	}
}
