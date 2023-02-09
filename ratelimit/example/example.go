package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"time"

	"github.com/go-redis/redis/v8"
	"gitlab.id.vin/gami/gami-common/adapters/cache"
	"gitlab.id.vin/gami/gami-common/logger"
	"gitlab.id.vin/gami/gami-common/ratelimit"
)

func main() {
	var (
		tokenBucket = ratelimit.NewTokenBucketLimiter(ratelimit.TokenBucketConfig{
			Config: ratelimit.Config{
				Service:        "Example",
				CounterType:    ratelimit.CounterTypeLocal,
				OnBurst:        OnBurst,
				OnReadyToAllow: OnReadyToAllow,
			},
			Interval: time.Millisecond * 4,
			Burst:    1,
		})

		leakyBucket = ratelimit.NewLeakyBucketLimiter(ratelimit.LeakyBucketConfig{
			Config: ratelimit.Config{
				Service:        "Example",
				CounterType:    ratelimit.CounterTypeLocal,
				OnBurst:        OnBurst,
				OnReadyToAllow: OnReadyToAllow,
			},
			LeakyRate: ratelimit.Rate{
				Throughput: 1,
				Interval:   time.Second * 2,
			},
			Burst: 2,
		})
		times     = 4
		rpi       = 8
		wait      = sync.WaitGroup{}
		requested = int64(0)
		allow     = int64(0)
		drop      = int64(0)
	)

	_ = tokenBucket
	_ = leakyBucket
	for loop := 0; loop < times; loop++ {
		for loop2 := 0; loop2 < rpi; loop2++ {
			go request(tokenBucket, &wait, &requested, &allow, &drop)
		}
		time.Sleep(time.Second)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func leakyBucketExample() {
	var (
		passed    int64
		dropped   int64
		requested int64

		redisClient = redis.NewFailoverClient(&redis.FailoverOptions{
			MasterName:    "mymaster",
			SentinelAddrs: []string{fmt.Sprintf("%v:%v", "0.0.0.0", 26379)},
			DB:            0,
			Password:      "",
		})
		_       = redisClient
		limiter = ratelimit.NewLeakyBucketLimiter(ratelimit.LeakyBucketConfig{
			LeakyRate: *ratelimit.NewRateBySecond(100),
			Config: ratelimit.Config{
				Service:     "Benchmark",
				Cache:       cache.NewCacheV2Adapter(redisClient),
				CounterType: ratelimit.CounterTypeLocal,
			},
			Burst: 400,
		})
		times              = 1
		requestPerInterval = 401
		wait               = sync.WaitGroup{}
	)

	for loop := 0; loop < times; loop++ {
		for loop := 0; loop < requestPerInterval; loop++ {
			go request(limiter, &wait, &requested, &passed, &dropped)
		}
		time.Sleep(time.Millisecond * 1)
	}

	wait.Wait()
	logger.Infof("Requested: %v - Passed: %v - Dropped: %v", requested, passed, dropped)
}

func multipleLeakyBucketExample() {
	var (
		passed    int64
		dropped   int64
		requested int64

		redisClient = redis.NewFailoverClient(&redis.FailoverOptions{
			MasterName:    "mymaster",
			SentinelAddrs: []string{fmt.Sprintf("%v:%v", "0.0.0.0", 26379)},
			DB:            0,
			Password:      "",
		})
		_ = redisClient

		limiter = ratelimit.NewLeakyBucketLimiter(ratelimit.LeakyBucketConfig{
			LeakyRate: *ratelimit.NewRateBySecond(100),
			Config: ratelimit.Config{
				Service:     "Benchmark",
				CounterType: ratelimit.CounterTypeLocal,
			},
			Burst: 400,
		})

		times              = 10
		requestPerInterval = 401
		wait               = sync.WaitGroup{}
		keys               = []string{"1", "2", "3"}
	)

	for loop := 0; loop < times; loop++ {
		for _, key := range keys {
			for loop := 0; loop < requestPerInterval; loop++ {
				go requestWithKey(limiter, key, &wait, &requested, &passed, &dropped)
			}
		}

		time.Sleep(time.Millisecond * 1)
	}

	wait.Wait()
	logger.Infof("Requested: %v - Passed: %v - Dropped: %v", requested, passed, dropped)
}

func request(limiter ratelimit.Limiter, wait *sync.WaitGroup, requestedCounter, passedCounter, droppedCounter *int64) {
	// now := time.Now()
	wait.Add(1)
	defer wait.Done()

	request := atomic.AddInt64(requestedCounter, 1)
	_ = request
	if !limiter.Allow() {
		logger.Infof("Request %v is dropped", request)
		atomic.AddInt64(droppedCounter, 1)
		return
	}

	passed := atomic.AddInt64(passedCounter, 1)
	_ = passed
	logger.Infof("%v", passed)
	// logger.Infof("Wait for: %v", time.Since(now))
}

func requestWithKey(limiter ratelimit.Limiter, key string, wait *sync.WaitGroup, requestedCounter, passedCounter, droppedCounter *int64) {
	now := time.Now()
	wait.Add(1)
	defer wait.Done()

	request := atomic.AddInt64(requestedCounter, 1)
	_ = request
	if !limiter.AllowWithKey(key) {
		// logger.Infof("Request %v is dropped", request)
		atomic.AddInt64(droppedCounter, 1)
		return
	}

	passed := atomic.AddInt64(passedCounter, 1)
	_ = passed
	logger.Infof("Wait for: %v", time.Since(now))
}
