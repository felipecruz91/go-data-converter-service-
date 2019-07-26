package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"../proto"
)

type server struct{}

func main() {

	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	proto.RegisterConverterServiceServer(s, &server{})
	reflection.Register(s)

	if e := s.Serve(listener); e != nil {
		panic(e)
	}

}

func (s *server) Convert(ctx context.Context, request *proto.Request) (*proto.Response, error) {

	body := request.GetBody()

	fmt.Printf("Body: %s", body)

	return &proto.Response{Result: []byte(body)}, nil
}
