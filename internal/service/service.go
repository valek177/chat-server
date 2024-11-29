package service

import (
	"context"

	"github.com/valek177/chat-server/grpc/pkg/chat_v1"
)

// ChatService is interface for chat on service
type ChatService interface {
	ConnectChat(ctx context.Context, chatID int64, username string,
		stream chat_v1.ChatV1_ConnectChatServer,
	) error
	SendMessage(ctx context.Context, chatID int64, message *chat_v1.Message) error
	CreateChat(ctx context.Context, req *chat_v1.CreateChatRequest) (int64, error)
	DeleteChat(ctx context.Context, id int64) error
}

// AccessService is interface for access service
type AccessService interface {
	TokenCtx(ctx context.Context) (context.Context, error)
}
