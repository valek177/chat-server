package chat

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/valek177/chat-server/grpc/pkg/chat_v1"
)

// SendMessage sends message to server
func (i *Implementation) SendMessage(_ context.Context, req *chat_v1.SendMessageRequest) (
	*emptypb.Empty, error,
) {
	log.Printf("Send message from %s with text: %s",
		req.GetMessage().GetFrom(), req.GetMessage().GetText())

	i.mxChannel.RLock()
	chatChan, ok := i.channels[req.GetChatId()]
	i.mxChannel.RUnlock()

	if !ok {
		return nil, status.Errorf(codes.NotFound, "chat not found")
	}

	chatChan <- req.GetMessage()

	return &emptypb.Empty{}, nil
}
