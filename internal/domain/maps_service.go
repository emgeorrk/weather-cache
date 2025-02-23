package domain

import (
	"context"
	"weather-cache/internal/model"
)

//go:generate go run go.uber.org/mock/mockgen -source=maps_service.go -destination=./mocks/maps_service_mock.go -package=mocks
type MapsService interface {
	GetLocationByCoords(ctx context.Context, lat, lon float64) (model.Location, error)
}
