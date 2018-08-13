package main

import (
	"time"
	"fmt"
)

func main() {
	now := time.Now()
	fmt.Println("day:", now.Day())
	fmt.Println("unix time:", now.Unix())
	fmt.Println("unix millisecond time:", now.UnixNano()/1e6)
}
