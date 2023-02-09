package ratelimit

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"gitlab.id.vin/gami/ps2-gami-common/logger"
)

type LeakyBucketConfig struct {
	Config
	LeakyRate Rate
	Burst     int
}

func (c LeakyBucketConfig) Clone() LeakyBucketConfig {
	return LeakyBucketConfig{
		Config:    c.Config.Clone(),
		LeakyRate: c.LeakyRate,
		Burst:     c.Burst,
	}
}

type LeakyBucket struct {
	limits sync.Map
	limit  *leakyBucketLimiter
	config LeakyBucketConfig
}

func NewLeakyBucketLimiter(config LeakyBucketConfig) Limiter {
	leaky := &LeakyBucket{
		config: config,
		limit:  newLeakyBucketLimiter(config, "basic"),
	}

	return leaky
}

func (l *LeakyBucket) Allow() bool {
	return l.limit.Allow()
}

func (l *LeakyBucket) AllowWithKey(key string) bool {
	config := l.config.Clone()
	limiter := newLeakyBucketLimiter(config, key)
	i, _ := l.limits.LoadOrStore(key, limiter)
	limit, _ := i.(*leakyBucketLimiter)

	return limit.Allow()
}

type leakyBucketLimiter struct {
	LeakyBucketConfig
	Counter Counter
	Key     string
}

func (l *leakyBucketLimiter) Allow() bool {
	var (
		wait time.Duration
	)

	drop := l.Counter.Add(1)
	defer func() {
		if l.Counter.Add(-1) == l.Burst-1 {
			// trigger ready to allow
			l.OnReadyToAllow(l.Service, l.Key)
		}
	}()

	if drop == l.Burst {
		l.OnBurst(l.Service, l.Key)
	}

	if drop > l.Burst {
		return false
	}

	wait = l.LeakyRate.Period() * time.Duration(drop)
	time.Sleep(wait)
	return true
}

func (l *leakyBucketLimiter) getKey(service string, key string) string {
	return fmt.Sprintf("%v:%v:%v", "leaky", service, key)
}

func newLeakyBucketLimiter(config LeakyBucketConfig, key string) *leakyBucketLimiter {
	limiter := &leakyBucketLimiter{
		LeakyBucketConfig: config,
		Key:               key,
	}

	switch config.Config.CounterType {
	case CounterTypeLocal:
		limiter.Counter = NewLocalCounter()
	case CounterTypeRedis:
		if config.Config.Cache == nil {
			logger.Fatalf(errors.New("Redis is nil"), "Redis is nil")
		}
		limiter.Counter = NewRedisCounter(config.Config.Cache, limiter.getKey(config.Config.Service, key))
	default:
		limiter.Counter = NewLocalCounter()
	}

	return limiter
}

// Rate determining the number of requests can be processed in a interval
type Rate struct {
	// Amount of requests can be processed in a interval
	Throughput int
	// The interval
	Interval time.Duration
}

// Period returns the duration a request needed to be processed
func (r *Rate) Period() time.Duration {
	return r.Interval / time.Duration(r.Throughput)
}

// Returns new pointer of Rate by through put & interval
//
// The new rate can be read as "throughput per interval"
// By default, throughput and interval should be positive
func NewRate(throughput int, interval time.Duration) *Rate {
	if throughput <= 0 {
		throughput = 1
	}

	if interval <= 0 {
		interval = time.Second
	}

	return &Rate{
		Throughput: throughput,
		Interval:   interval,
	}
}

// Returns new pointer of Rate by through put per second
func NewRateBySecond(throughput int) *Rate {
	return NewRate(throughput, time.Second)
}
