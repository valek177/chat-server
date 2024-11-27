package service

import (
	"context"

	"github.com/valek177/chat-server/grpc/pkg/chat_v1"
)

// ChatService is interface for chat on service
type ChatService interface {
	// ConnectChat(ctx context.Context, id int64, user_id int64) error //(chan *model.Message, error)
	// SendMessage()
	CreateChat(ctx context.Context, req *chat_v1.CreateChatRequest) (int64, error)
	DeleteChat(ctx context.Context, id int64) error
}

// AccessService is interface for access service
type AccessService interface {
	TokenCtx(ctx context.Context) (context.Context, error)
}
