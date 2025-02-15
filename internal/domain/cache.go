package domain

import (
	"context"
	"weather-cache/internal/model"
)

type Cache interface {
	Get(ctx context.Context, key string) (model.Weather, bool)
	Set(ctx context.Context, key string, value model.Weather)
}
