package ratelimit

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/Workiva/go-datastructures/queue"
	"gitlab.id.vin/gami/ps2-gami-common/logger"
)

type SlidingWindowConfig struct {
	Config
	Interval time.Duration
	Burst    int
}

func (c *SlidingWindowConfig) Clone() SlidingWindowConfig {
	return SlidingWindowConfig{
		Config:   c.Config.Clone(),
		Interval: c.Interval,
		Burst:    c.Burst,
	}
}

type SlidingWindow struct {
	limits sync.Map
	limit  *slidingWindowLimiter
	config SlidingWindowConfig
}

type slidingWindowLimiter struct {
	SlidingWindowConfig
	queue   queue.Queue
	mu      sync.Mutex
	counter Counter
	key     string
}

func NewSlidingWindowLimiter(config SlidingWindowConfig) Limiter {
	window := &SlidingWindow{
		config: config,
		limit:  newSlidingWindowLimiter(config, "basic"),
	}

	return window
}

func (s *SlidingWindow) Allow() bool {
	return s.limit.Allow()
}

func (l *SlidingWindow) AllowWithKey(key string) bool {
	config := l.config.Clone()
	i, _ := l.limits.LoadOrStore(key, newSlidingWindowLimiter(config, key))
	limit, _ := i.(*slidingWindowLimiter)
	return limit.Allow()
}

func (s *slidingWindowLimiter) Allow() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	now := time.Now()
	counter := s.update(now) + 1

	if counter > s.Burst {
		return false
	}

	s.counter.Add(1)

	if err := s.queue.Put(now); err != nil {
		return false
	}

	return true
}

func (s *slidingWindowLimiter) update(now time.Time) int {
	/*
		all request before the time will be removed
		thus the counter will be minus by the length of obsoleted requests
	*/

	if s.queue.Empty() {
		return s.counter.Load()
	}

	last := now.Add(-s.Interval)
	rs, err := s.queue.TakeUntil(func(item interface{}) bool {
		t, ok := item.(time.Time)
		return ok && t.Before(last)
	})

	if err != nil {
		// TakeUntil error
		return 0
	}

	counter := s.counter.Add(-len(rs))
	return counter
}

func (s *slidingWindowLimiter) getKey(service, key string) string {
	return fmt.Sprintf("%v:%v:%v", "window", service, key)
}

func newSlidingWindowLimiter(config SlidingWindowConfig, key string) *slidingWindowLimiter {
	limiter := &slidingWindowLimiter{
		SlidingWindowConfig: config,
		key:                 key,
	}

	switch config.Config.CounterType {
	case CounterTypeLocal:
		limiter.counter = NewLocalCounter()
	case CounterTypeRedis:
		if config.Config.Cache == nil {
			logger.Fatalf(errors.New("Redis is nil"), "Redis is nil")
		}
		limiter.counter = NewRedisCounter(config.Config.Cache, limiter.getKey(config.Config.Service, key))
	default:
		limiter.counter = NewLocalCounter()
	}
	return limiter
}
