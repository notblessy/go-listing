package db

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
	"github.com/notblessy/go-listing/config"
	"github.com/sirupsen/logrus"
)

func InitPool() *redis.Pool {
	pool := &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial(
				"tcp",
				fmt.Sprintf("%s:%s", config.RedisHost(), config.RedisPort()),
				redis.DialDatabase(config.RedisDB()),
			)
			if err != nil {
				logrus.Fatal("Failed to connect redis")
			}
			return conn, err
		},
	}

	return pool
}
