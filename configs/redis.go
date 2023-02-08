package configs

import "fmt"

// Redis represents configuration of Redis caching database.
type Redis struct {
	Host         string `default:"127.0.0.1" envconfig:"REDIS_HOST"`
	Port         int    `default:"6379" envconfig:"REDIS_PORT"`
	Password     string `default:"" envconfig:"REDIS_PASSWORD"`
	Database     int    `default:"0" envconfig:"REDIS_DB"`
	MasterName   string `default:"mymaster" envconfig:"REDIS_MASTER_NAME"`
	PoolSize     int    `default:"2000" envconfig:"REDIS_POOL_SIZE"`
	MinIdleConns int    `default:"100" envconfig:"REDIS_MIN_IDLE_CONNS"`
}

// URL return redis connection URL.
func (c *Redis) URL() string {
	return fmt.Sprintf("%v:%v", c.Host, c.Port)
}
