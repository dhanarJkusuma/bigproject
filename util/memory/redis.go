package memory

import (
	"bigproject/util/config"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func GetRedisConnectionPool() redis.Pool {
	conf := config.GetConfig()

	hostAndPort := fmt.Sprintf("%v:%v", conf.RedisHost, conf.RedisPort)

	pool := redis.Pool{
		MaxActive:   conf.RedisMaxActive,
		MaxIdle:     conf.RedisMaxIdle,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", hostAndPort)
		},
	}
	return pool
}
