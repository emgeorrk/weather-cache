package weather

import (
	"context"
	"weather-cache/internal/model"
)

func (s Service) GetWeatherByCoords(ctx context.Context, lat, lon float64) (model.Weather, error) {
	loc, err := s.mapsService.GetLocationByCoords(ctx, lat, lon)
	if err != nil {
		s.log.Error("failed to get address by coords", "err", err)
		return model.Weather{}, err
	}

	weather, ok := s.cache.Get(ctx, loc.City)
	if !ok {
		weather, err = s.UpdateWeather(ctx, loc.City)
		if err != nil {
			s.log.Error("failed to update weather", "err", err)
			return model.Weather{}, err
		}
	}

	return weather, nil
}
