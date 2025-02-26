package domain

import (
	"context"
	"weather-cache/internal/model"
)

//go:generate go run go.uber.org/mock/mockgen -source=weather_provider.go -destination=./mocks/weather_provider_mock.go -package=mocks
type WeatherProvider interface {
	FetchWeather(ctx context.Context, city model.Location) (model.Weather, error)
}
