package weather

import (
	"context"
	"weather-cache/internal/model"
)

func (s Service) GetWeatherByCity(ctx context.Context, city string) (model.Weather, error) {
	weather, ok := s.cache.Get(ctx, city)
	if !ok {
		return model.Weather{}, nil
	}

	return weather, nil
}
