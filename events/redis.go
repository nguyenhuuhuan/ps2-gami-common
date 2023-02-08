package events

import (
	"context"

	"gitlab.id.vin/gami/ps2-gami-common/adapters/cache"
	"gitlab.id.vin/gami/ps2-gami-common/logger"
)

type redisEvent struct {
	adapter cache.CachedV2Adapter
}

// RedisEvent interface
type RedisEvent interface {
	HandleFunc(a []interface{})
}

// NewRedisEvent struct
func NewRedisEvent(adapter cache.CachedV2Adapter) RedisEvent {
	return &redisEvent{adapter: adapter}
}

// HandleFunc when some function publish event
func (e *redisEvent) HandleFunc(args []interface{}) {
	ctx := context.Background()
	for _, arg := range args {
		key, ok := arg.(string)
		if ok {
			if err := e.adapter.Del(ctx, key); err != nil {
				logger.Errorf("Redis Del Key Error: %v ", err)
			}
		}
	}
}
