package domain

import (
	"context"
	"weather-cache/internal/model"
)

type WeatherProvider interface {
	FetchWeather(ctx context.Context, city model.Location) (model.Weather, error)
}
