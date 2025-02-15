package domain

import "context"

type MapsProvider interface {
	GetCityByCoords(ctx context.Context, lat, lon float64) (city string, err error)
}
