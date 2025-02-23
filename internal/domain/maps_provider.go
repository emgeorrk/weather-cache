package domain

import "context"

//go:generate go run go.uber.org/mock/mockgen -source=maps_provider.go -destination=./mocks/maps_provider_mock.go -package=mocks
type MapsProvider interface {
	GetCityByCoords(ctx context.Context, lat, lon float64) (city string, err error)
}
