package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/valek177/chat-server/grpc/pkg/chat_v1"

	"github.com/brianvoe/gofakeit"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

const grpcPort = 50051

type server struct {
	chat_v1.UnimplementedChatV1Server
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	chat_v1.RegisterChatV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// CreateChat creates new chat with parameters
func (s *server) CreateChat(_ context.Context, req *chat_v1.CreateChatRequest) (
	*chat_v1.CreateChatResponse, error,
) {
	log.Printf("Create new chat with name %s", req.GetName())

	id := gofakeit.Int64()

	return &chat_v1.CreateChatResponse{
		Id: id,
	}, nil
}

// DeleteChat removes chat
func (s *server) DeleteChat(_ context.Context, req *chat_v1.DeleteChatRequest) (
	*emptypb.Empty, error,
) {
	log.Printf("Delete chat with id %d", req.GetId())

	return &emptypb.Empty{}, nil
}

// SendMessage sends message to server
func (s *server) SendMessage(_ context.Context, req *chat_v1.SendMessageRequest) (
	*emptypb.Empty, error,
) {
	log.Printf("Send message from %s with text: %s", req.GetFrom(), req.GetText())

	return &emptypb.Empty{}, nil
}
