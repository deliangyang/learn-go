package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

var (
	pool         *redis.Pool
	redisService = flag.String("redisServer", ":6379", "")
)

func main() {
	flag.Parse()
	pool = newPool(*redisService)
	c := pool.Get()
	fmt.Println(c)
	fmt.Println(pool.Stats())
	c = pool.Get()
	fmt.Println(c)
	fmt.Println(pool.Stats())
	c = pool.Get()
	fmt.Println(c)
	fmt.Println(pool.Stats())
	c = pool.Get()
	fmt.Println(c)
	fmt.Println(pool.Stats())
	c = pool.Get()
	fmt.Println(c)
	fmt.Println(pool.Stats())
}

func newPool(addr string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		MaxActive: 3,
		IdleTimeout: 240 * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", addr) },
	}
}
