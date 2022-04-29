package app

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

var RedisConn *redis.Pool

func InitRedis() {
	cfg := GetConfig()
	RedisConn = &redis.Pool{
		MaxIdle:     cfg.Redis.MaxIdle,
		MaxActive:   cfg.Redis.MaxActive,
		IdleTimeout: time.Duration(cfg.Redis.IdleTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", cfg.Redis.Address)
			if err != nil {
				return nil, err
			}
			if cfg.Redis.Password != "" {
				if _, err := c.Do("AUTH", cfg.Redis.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
	}

	if cfg.RunMode == "debug" {
		RedisConn.TestOnBorrow = func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		}
	}
}
