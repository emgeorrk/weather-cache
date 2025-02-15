package domain

import (
	"context"
	"weather-cache/internal/model"
)

type WeatherService interface {
	GetWeatherByCoords(ctx context.Context, lat, lon float64) (model.Weather, error)
	GetWeatherByCity(ctx context.Context, city string) (model.Weather, error)
	UpdateWeather(ctx context.Context, city string) (model.Weather, error)
}
