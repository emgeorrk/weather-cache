package domain

import (
	"context"
	"weather-cache/internal/model"
)

type MapsService interface {
	GetLocationByCoords(ctx context.Context, lat, lon float64) (model.Location, error)
}
