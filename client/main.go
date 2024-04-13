package main

import (
	"fmt"

	pb "github.com/Abhishekagarwal1597/basic-go-grpc/proto"

	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const port = ":8000"

func main() {
	clientConnection, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Fprintln(os.Stdout, []any{"Error in GRPC Server Connection..."}...)
	}
	defer clientConnection.Close()

	svcClient := pb.NewGreetServiceClient(clientConnection)

	names := &pb.NamesList{
		Names: []string{"Abhishek", "Himanshu", "Vijay", "Savita"},
	}

	// callSayHello(svcClient)
	// callServerStream(svcClient, names)
	// callClientStream(svcClient, names)
	callSayHelloBidirectionalStream(svcClient, names)
	// client.Greet(context.Background(), &pb.GreetRequest{
	// clientConnection.Connect()
	// fmt.Println("Client Connected to Server...")

}
