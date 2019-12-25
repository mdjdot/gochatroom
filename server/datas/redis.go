package datas

import "github.com/garyburd/redigo/redis"

// Pool redis 连接池
var Pool *redis.Pool

func init() {
	Pool = &redis.Pool{
		MaxIdle:     8,
		MaxActive:   8,
		IdleTimeout: 100,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", "127.0.0.1:6379")
			return conn, err
		},
	}
}
