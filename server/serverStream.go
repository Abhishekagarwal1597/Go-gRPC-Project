package main

import (
	"fmt"
	"time"

	pb "github.com/Abhishekagarwal1597/basic-go-grpc/proto"
)

func (s *helloServer) SayHelloServerStreaming(namelist *pb.NamesList, stream pb.GreetService_SayHelloServerStreamingServer) error {
	fmt.Println("nmelist server:", namelist)
	for _, namelist := range namelist.Names {
		res := &pb.HelloResponse{
			Message: "Hello " + namelist,
		}
		err := stream.Send(res)
		if err != nil {
			return err
		}
		time.Sleep(1 * time.Second)

	}
	return nil
}
