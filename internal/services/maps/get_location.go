package maps

import (
	"context"
	"weather-cache/internal/model"
)

func (s Service) GetLocationByCoords(ctx context.Context, lat, lon float64) (model.Location, error) {
	city, err := s.mapsProvider.GetCityByCoords(ctx, lat, lon)
	if err != nil {
		s.log.Error("failed to get address by coords", "err", err)
		return model.Location{}, err
	}

	return model.Location{
		City: city,
		Lat:  lat,
		Lon:  lon,
	}, nil
}
