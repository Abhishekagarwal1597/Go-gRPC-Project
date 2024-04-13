package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Abhishekagarwal1597/basic-go-grpc/proto"
)

func callServerStream(client pb.GreetServiceClient, namelist *pb.NamesList) {
	log.Printf("Streaming started...")
	// for range namelist {
	stream, err := client.SayHelloServerStreaming(context.Background(), namelist)
	if err != nil {
		log.Fatalf("could not send names: %v", err)
	}
	for {
		msg, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while streaming %v", err)
			}
		}
		log.Println(msg)

	}

	log.Printf("Streaming finished")
}
