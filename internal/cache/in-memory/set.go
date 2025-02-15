package in_memory

import (
	"context"
	"time"
	"weather-cache/internal/model"
)

func (c *InMemoryCache) Set(_ context.Context, key string, value model.Weather) {
	c.Map.Store(key, model.CacheValue{
		Weather:    value,
		Expiration: time.Now().Add(c.ttl),
	})
}
