package model

import "time"

type CacheValue struct {
	Weather
	Expiration time.Time
}
