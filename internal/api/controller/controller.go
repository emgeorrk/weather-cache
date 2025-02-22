package controller

import (
	"go.uber.org/fx"
	"weather-cache/internal/domain"
	"weather-cache/pkg/logger"
)

var Module = fx.Options(
	fx.Provide(NewController),
)

type Controller struct {
	log            logger.Logger
	weatherService domain.WeatherService
}

func NewController(log logger.Logger, weather domain.WeatherService) Controller {
	return Controller{
		log:            log,
		weatherService: weather,
	}
}
