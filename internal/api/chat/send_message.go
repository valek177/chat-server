package chat

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/valek177/chat-server/grpc/pkg/chat_v1"
)

// SendMessage sends message to server
func (i *Implementation) SendMessage(ctx context.Context, req *chat_v1.SendMessageRequest) (
	*emptypb.Empty, error,
) {
	log.Printf("Send message from %s with text: %s",
		req.GetMessage().GetFrom(), req.GetMessage().GetText())

	chatID, err := i.chatService.GetChatIdByName(ctx, req.Chatname)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	err = i.chatService.SendMessage(ctx, chatID, req.Message)
	if err != nil {
		return &emptypb.Empty{}, fmt.Errorf("error while sending message %v", err.Error())
	}

	return &emptypb.Empty{}, nil
}
