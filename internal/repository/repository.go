package repository

import (
	"context"

	"github.com/valek177/chat-server/grpc/pkg/chat_v1"
)

// ChatRepository is interface for chat logic
type ChatRepository interface {
	CreateChat(ctx context.Context, req *chat_v1.CreateChatRequest) (int64, error)
	DeleteChat(ctx context.Context, id int64) error
}
