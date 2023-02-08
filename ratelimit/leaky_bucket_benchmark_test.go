package ratelimit

import (
	"fmt"
	"testing"
	"time"
)

func BenchmarkLeakyBucket(b *testing.B) {
	var (
		limiter = NewLeakyBucketLimiter(LeakyBucketConfig{
			LeakyRate: *NewRate(1000, time.Millisecond),
			Burst:     1 << 14,
			Config: Config{
				Service:     "Benchmark",
				CounterType: CounterTypeLocal,
			},
		})
	)

	for loop := 0; loop < b.N; loop++ {
		do(limiter)
	}
}

func BenchmarkLeakyBucketWithKey(b *testing.B) {
	var (
		limiter = NewLeakyBucketLimiter(LeakyBucketConfig{
			LeakyRate: *NewRateBySecond(80),
			Burst:     1 << 10,
			Config: Config{
				Service:     "Benchmark",
				CounterType: CounterTypeLocal,
			},
		})
		user = "user"
	)

	for loop := 0; loop < b.N; loop++ {
		doWithKey(limiter, user)
	}
}

func BenchmarkLeakyBucketWithMultipleKey(b *testing.B) {
	var (
		limiter = NewLeakyBucketLimiter(LeakyBucketConfig{
			LeakyRate: *NewRateBySecond(80),
			Burst:     400,
			Config: Config{
				Service:     "Benchmark",
				CounterType: CounterTypeLocal,
			},
		})
		userAmount = 10000
	)

	for loop := 0; loop < b.N; loop++ {
		for loop2 := 0; loop < userAmount; loop2++ {
			doWithKey(limiter, fmt.Sprintf("User%v", loop2))
		}
	}
}

func do(limiter Limiter) {
	if !limiter.Allow() {
		return
	}
}

func doWithKey(limiter Limiter, key string) {
	if !limiter.AllowWithKey(key) {
		return
	}
}
