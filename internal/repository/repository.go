package repository

import (
	"context"

	"github.com/valek177/chat-server/grpc/pkg/chat_v1"
	"github.com/valek177/chat-server/internal/model"
)

// ChatRepository is interface for chat logic
type ChatRepository interface {
	GetChatIdByName(ctx context.Context, chatname string) (int64, error)
	CreateChat(ctx context.Context, req *chat_v1.CreateChatRequest) (int64, error)
	DeleteChat(ctx context.Context, id int64) error
}

// LogRepository is interface for logging user actions
type LogRepository interface {
	CreateRecord(ctx context.Context, record *model.Record) (int64, error)
}
