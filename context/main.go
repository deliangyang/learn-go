package main

import (
	"context"
	"fmt"
	"github.com/oschwald/geoip2-golang"
	"net"
)


type Int int64

func (i Int) test() {
	fmt.Println(i)
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for i := 0; i < 10; i++  {
			fmt.Println("current i: ", i)
			if i == 5 {
				cancel()
			}
		}
	}(ctx)

	db, err := geoip2.Open("context/GeoLite2-Country.mmdb")
	if err != nil {
		panic(err)
	}

	ip := net.ParseIP("120.24.37.249")

	record, err := db.City(ip)
	if err != nil {
		panic(err)
	}
	fmt.Println(record.City)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("done")
			return
		}
	}


}
