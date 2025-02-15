package in_memory

import (
	"context"
	"time"
	"weather-cache/internal/model"
)

func (c *InMemoryCache) Get(_ context.Context, key string) (model.Weather, bool) {
	val, ok := c.Map.Load(key)
	if !ok || val.(model.CacheValue).Expiration.Before(time.Now()) {
		c.Delete(key)
		return model.Weather{}, false
	}

	return val.(model.CacheValue).Weather, true
}
