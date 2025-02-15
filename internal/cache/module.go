package cache

import (
	"go.uber.org/fx"
	"weather-cache/config"
	inmemory "weather-cache/internal/cache/in-memory"
	"weather-cache/internal/constants"
	"weather-cache/internal/domain"
)

var Module = fx.Options(
	fx.Provide(NewCache),
)

func NewCache(config config.Config) (domain.Cache, error) {
	switch config.CacheType {
	case "in-memory":
		return inmemory.New(config), nil
	default:
		return nil, constants.ErrCacheTypeNotSupported
	}
}
