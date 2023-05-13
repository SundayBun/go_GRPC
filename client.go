package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"simpleGRPCServer/chat"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	messageClient := chat.Message{Body: "Hello from the client. "}
	response, err := c.SayHello(context.Background(), &messageClient)
	if err != nil {
		log.Fatalf("Error when calling Say Hello: %s", err)
	}

	log.Printf("Response from Server: %s", response.Body)
}
