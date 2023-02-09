package ratelimit

import (
	"context"
	"sync/atomic"

	"gitlab.id.vin/gami/gami-common/adapters/cache"
	"gitlab.id.vin/gami/gami-common/logger"
)

type Limiter interface {
	// Allow a request to be executed.
	//
	// Allow return true if it could be executed, otherwise it returns false then it should be discarded
	Allow() bool
	// Allow a request with a specific key to be executed.
	//
	// AllowWithKey should be used in case limiting a specific a user from using the resources
	AllowWithKey(key string) bool
}

type Config struct {
	Service        string
	Cache          cache.CachedV2Adapter
	CounterType    CounterType
	OnBurst        func(service string, key string) // trigger when the ratelimit will be bursted in the next request
	OnReadyToAllow func(service string, key string) // trigger when the ratelimit is ready for allowing again
}

func (c Config) Clone() Config {
	return Config{
		Service:     c.Service,
		Cache:       c.Cache,
		CounterType: c.CounterType,
	}
}

type CounterType string

const (
	// CounterTypeLocal using a counter storaged in local
	CounterTypeLocal CounterType = "LOCAL"
	// CounterTyRedis using a counter storaged in redis, redis client must be provided in Config when this type is using.
	//
	// This type could be used for multi instances rate limiting
	CounterTypeRedis CounterType = "REDIS"
)

type Counter interface {
	// Store val to the counter
	Store(val int)
	// Load val from the counter
	Load() int
	// Add val to the counter and returns the new value.
	Add(val int) int
}

type localCounter struct {
	counter int32
}

func (c *localCounter) Store(val int) {
	atomic.StoreInt32(&c.counter, int32(val))
}

func (c *localCounter) Load() int {
	return int(atomic.LoadInt32(&c.counter))
}

func (c *localCounter) Add(val int) int {
	counter := atomic.AddInt32(&c.counter, int32(val))
	return int(counter)
}

type redisCounter struct {
	cache cache.CachedV2Adapter
	key   string
}

func (c *redisCounter) Store(val int) {
	if err := c.cache.Set(context.Background(), c.key, val, 0); err != nil {
		logger.Errorf("Set counter error: %v", err)
	}
}

func (c *redisCounter) Load() int {
	var val int
	_ = c.cache.Get(context.Background(), c.key, &val)
	return val
}

func (c *redisCounter) Add(val int) int {
	old, err := c.cache.IncrBy(context.Background(), c.key, int64(val))
	if err != nil {
		logger.Errorf("Incr counter error: %v", err)
	}
	return int(old)
}

func NewLocalCounter() Counter {
	return &localCounter{}
}

func NewLocalCounterWith(val int) Counter {
	return &localCounter{counter: int32(val)}
}

func NewRedisCounter(cache cache.CachedV2Adapter, key string) Counter {
	return &redisCounter{
		cache: cache,
		key:   key,
	}
}
