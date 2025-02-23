package domain

import (
	"context"
	"weather-cache/internal/model"
)

//go:generate go run go.uber.org/mock/mockgen -source=cache.go -destination=./mocks/cache_mock.go -package=mocks
type Cache interface {
	Get(ctx context.Context, key string) (model.Weather, bool)
	Set(ctx context.Context, key string, value model.Weather)
}
