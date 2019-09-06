package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

func newPool(addr string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		// Dial or DialContext must be set. When both are set, DialContext takes precedence over Dial.
		Dial: func() (redis.Conn, error) { return redis.Dial("tcp", addr) },
	}
}

var (
	pool        *redis.Pool
	redisServer = flag.String("redisServer", "127.0.0.1:6379", "")
)

func main() {
	flag.Parse()
	pool = newPool(*redisServer)
	p := pool.Get()
	ok, err := p.Do("SET", "xxxxxxxx", "123123123")

	if err != nil {
		panic(err)
	}

	fmt.Println(ok)

	reply, err := redis.String(pool.Get().Do("GET", "xxx"))
	if err != nil {
		panic(err)
	}

	fmt.Println(reply)

}
