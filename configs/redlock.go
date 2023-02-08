package configs

import (
	"fmt"
	"time"
)

// Redis represents configuration of Redis caching database.
type RedLock struct {
	Host          string        `default:"127.0.0.1" envconfig:"REDLOCK_HOST"`
	Port          int           `default:"6379" envconfig:"REDLOCK_PORT"`
	Password      string        `default:"" envconfig:"REDLOCK_PASSWORD"`
	MaxIdle       int           `default:"300" envconfig:"RREDLOCK_MAX_IDLE"`
	MaxActive     int           `default:"0" envconfig:"REDIS_MAX_ACTIVE"`
	ExpiredTime   time.Duration `default:"5s" envconfig:"REDLOCK_EXPIRED_TIME"`
	RetryTime     int           `default:"500" envconfig:"REDLOCK_RETRY_TIME"`
	RetryDuration time.Duration `default:"10ms" envconfig:"REDLOCK_RETRY_DURATION"`
}

// URL return redis connection URL.
func (c *RedLock) URL() string {
	return fmt.Sprintf("%v:%v", c.Host, c.Port)
}
