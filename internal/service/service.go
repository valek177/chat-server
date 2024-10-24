package service

import (
	"context"

	"github.com/valek177/chat-server/grpc/pkg/chat_v1"
)

// ChatService is interface for chat on service
type ChatService interface {
	CreateChat(ctx context.Context, req *chat_v1.CreateChatRequest) (int64, error)
	DeleteChat(ctx context.Context, id int64) error
}
