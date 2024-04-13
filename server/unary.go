package main

import (
	"context"

	pb "github.com/Abhishekagarwal1597/basic-go-grpc/proto"
)

func (s *helloServer) SayHello(c context.Context, r *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}
