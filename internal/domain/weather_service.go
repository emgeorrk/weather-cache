package domain

import (
	"context"
	"weather-cache/internal/model"
)

//go:generate go run go.uber.org/mock/mockgen -source=weather_service.go -destination=./mocks/weather_service_mock.go -package=mocks
type WeatherService interface {
	GetWeatherByCoords(ctx context.Context, lat, lon float64) (model.Weather, error)
	GetWeatherByCity(ctx context.Context, city string) (model.Weather, error)
	UpdateWeather(ctx context.Context, city string) (model.Weather, error)
}
