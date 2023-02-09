package redlock

import (
	"gitlab.id.vin/gami/gami-common/configs"

	"github.com/go-redsync/redsync"
	"github.com/gomodule/redigo/redis"
)

type RedLock interface {
	NewMutex(key string) *redsync.Mutex
}

type redLock struct {
	redSync *redsync.Redsync
	opts    []redsync.Option
}

func NewRedLock(redLockCfg configs.RedLock) (RedLock, error) {
	pool := &redis.Pool{
		MaxIdle:   redLockCfg.MaxIdle,
		MaxActive: redLockCfg.MaxActive,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", redLockCfg.URL(), redis.DialPassword(configs.AppConfig.RedLock.Password))
			if err != nil {
				return nil, err
			}
			return conn, nil
		},
	}

	conn := pool.Get()
	_, err := conn.Do("PING")
	if err != nil {
		return nil, err
	}

	redSync := redsync.New([]redsync.Pool{pool})

	opts := []redsync.Option{}
	opts = append(opts, redsync.SetExpiry(redLockCfg.ExpiredTime))
	opts = append(opts, redsync.SetTries(redLockCfg.RetryTime))
	opts = append(opts, redsync.SetRetryDelay(redLockCfg.RetryDuration))

	return &redLock{
		redSync: redSync,
		opts:    opts,
	}, nil
}

func (r *redLock) NewMutex(key string) *redsync.Mutex {
	return r.redSync.NewMutex(key, r.opts...)
}
