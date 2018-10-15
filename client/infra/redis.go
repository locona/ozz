package infra

import (
	"log"

	"github.com/garyburd/redigo/redis"
	"github.com/locona/ozz/client/config"
)

var Redis *redis.Pool

// redis のセッティング
func InitRedis() {
	host := config.Config.Redis.Host
	port := config.Config.Redis.Port
	maxConnections := config.Config.Redis.MaxConn
	conn := redis.NewPool(func() (redis.Conn, error) {
		conn, err := redis.Dial("tcp", host+":"+port)
		if err != nil {
			log.Fatal(err)
		}
		return conn, err
	}, maxConnections)

	_, err := conn.Get().Do("PING")
	if err != nil {
		log.Fatal(err)
	}

	Redis = conn
}
