package chat

import (
	"context"
	"log"
)

type Server struct {
	UnimplementedChatServiceServer
}

func (c *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Recieved messaage body from client: %s", message.Body)
	//return new Message
	return &Message{Body: "Hello from the server"}, nil
}
