package main

import (
	"log"

	chat "grpcchat/pb"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	response, err := c.SayHello(context.Background(), &chat.Message{Client: "eco", Body: "Hello From Client!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)

	response, err = c.BroadcastMessage(context.Background(), &chat.Message{Client: "eco", Body: "Hello From Client!"})
	if err != nil {
		log.Fatalf("Error when calling BroadcastMessage: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)

}
