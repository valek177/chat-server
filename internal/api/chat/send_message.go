package chat

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/valek177/chat-server/grpc/pkg/chat_v1"
)

// SendMessage sends message to server
func (i *Implementation) SendMessage(_ context.Context, req *chat_v1.SendMessageRequest) (
	*emptypb.Empty, error,
) {
	log.Printf("Send message from %s with text: %s", req.GetFrom(), req.GetText())

	return &emptypb.Empty{}, nil
}
