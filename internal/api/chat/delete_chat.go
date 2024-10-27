package chat

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/valek177/chat-server/grpc/pkg/chat_v1"
)

// DeleteChat removes chat
func (i *Implementation) DeleteChat(ctx context.Context, req *chat_v1.DeleteChatRequest) (
	*emptypb.Empty, error,
) {
	err := i.chatService.DeleteChat(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	log.Printf("deleted chat with id: %d", req.GetId())

	return &emptypb.Empty{}, nil
}
