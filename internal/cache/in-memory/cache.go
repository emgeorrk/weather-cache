package in_memory

import (
	"sync"
	"time"
	"weather-cache/config"
	"weather-cache/internal/constants"
)

type InMemoryCache struct {
	sync.Map
	ttl time.Duration
}

func New(config config.Config) *InMemoryCache {
	ttl, _ := time.ParseDuration(config.CacheTTL)
	if ttl == 0 {
		ttl = constants.DefaultCacheTTL
	}

	return &InMemoryCache{
		Map: sync.Map{},
		ttl: ttl,
	}
}
