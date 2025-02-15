package weather

import (
	"context"
	"weather-cache/internal/model"
)

func (s Service) UpdateWeather(ctx context.Context, city string) (model.Weather, error) {
	weather, err := s.weatherProvider.FetchWeather(ctx, model.Location{City: city})
	if err != nil {
		s.log.Error("failed to fetch weather", "err", err)
		return model.Weather{}, err
	}

	s.cache.Set(ctx, city, weather)

	return weather, nil
}
