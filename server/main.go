package main

import (
	"net"

	pb "github.com/deliangyang/hello-test/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	Address = "127.0.0.1:5000"
)

type helloService struct {
}

func (helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	resp := new(pb.HelloReply)
	resp.Message = "hello world " + in.Name
	return resp, nil
}

var HelloService = helloService{}

var (
	options []grpc.ServerOption
)

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		panic(err)
	}

	var interceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		println(ctx)
		return handler(ctx, req)
	}

	options = append(options, grpc.UnaryInterceptor(interceptor))

	s := grpc.NewServer(options...)
	pb.RegisterHelloServer(s, HelloService)
	_ = s.Serve(listen)
}
