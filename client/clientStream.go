package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/Abhishekagarwal1597/basic-go-grpc/proto"
)

func callClientStream(client pb.GreetServiceClient, namelist *pb.NamesList) {
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("Error while calling GreetClientStream: %v\n", err)
	}

	for _, name := range namelist.Names {
		fmt.Println("sending req with server:", name)
		req := &pb.HelloRequest{
			Name: name,
		}
		err := stream.Send(req)
		if err != nil {
			log.Fatalf("Error while sending %v", err)
		}
		time.Sleep(2 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving %v", err)
	}
	log.Printf("Receive Message from Server %v", res.Messages)

	log.Printf("Client Streaming finished")
	// return nil
}
