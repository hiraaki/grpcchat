package chat

import (
	"fmt"
	"grpcchat/pb"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

type Server struct {
	pb.UnimplementedChatServiceServer
}

func (s *Server) SayHello(ctx context.Context, in *pb.Message) (*pb.Message, error) {

	log.Printf("Receive message body from client: %s %s", in.Client, in.Body)
	md, _ := metadata.FromIncomingContext(ctx)
	fmt.Printf("%v", md["authority"])
	return &pb.Message{Body: "Hello From the Server!"}, nil
}

func (s *Server) BroadcastMessage(ctx context.Context, in *pb.Message) (*pb.Message, error) {
	return &pb.Message{Body: "Hello To All"}, nil
}
