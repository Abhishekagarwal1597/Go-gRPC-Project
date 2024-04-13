package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/Abhishekagarwal1597/basic-go-grpc/proto"
	"google.golang.org/grpc"
)

type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Errorf(err.Error())
	}
	fmt.Println("server started at ", lis.Addr())
	server := grpc.NewServer()
	// client := &greetServiceClient{}

	pb.RegisterGreetServiceServer(server, &helloServer{})
	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}

}
