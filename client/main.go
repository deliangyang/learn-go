package main

import (
	"context"

	pb "github.com/deliangyang/hello-test/proto" // 引入编译生成的包
	"google.golang.org/grpc"
)

const (
	Address = "127.0.0.1:5000"
)

func main() {
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := pb.NewHelloClient(conn)

	reqBody := new(pb.HelloRequest)
	reqBody.Name = "li si"
	r, err := c.SayHello(context.Background(), reqBody)
	if err != nil {
		panic(err)
	}

	println(r.Message)
}
