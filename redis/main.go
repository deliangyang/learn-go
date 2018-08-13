package main

import (
	redis2 "github.com/garyburd/redigo/redis"
	"os"
	"fmt"
)

func main() {
	redis, err := redis2.Dial("tcp", "www.ydl.com:6379")
	if err != nil {
		panic(err)
		os.Exit(-1)
	}
	res, err := redis.Do("set", "test", 100)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)
}
