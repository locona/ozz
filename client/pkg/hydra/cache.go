package hydra

import (
	"github.com/garyburd/redigo/redis"
)

type cache struct {
	redis *redis.Pool
}

func NewCache(redisConn *redis.Pool) *cache {
	return &cache{redis: redisConn}
}

func (c *cache) Get() (string, error) {
	return redis.String(c.redis.Get().Do("GET", c.key()))
}

func (c *cache) Set(token string) error {
	_, err := c.redis.Get().Do("SET", c.key(), token)
	return err
}

func (c *cache) key() string {
	return "access_token"
}
