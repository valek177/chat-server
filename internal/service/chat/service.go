package chat

import (
	"github.com/valek177/chat-server/internal/client/db"
	"github.com/valek177/chat-server/internal/repository"
	"github.com/valek177/chat-server/internal/service"
)

type serv struct {
	chatRepository repository.ChatRepository
	txManager      db.TxManager
}

// NewService creates new service with settings
func NewService(
	chatRepository repository.ChatRepository,
	txManager db.TxManager,
) service.ChatService {
	return &serv{
		chatRepository: chatRepository,
		txManager:      txManager,
	}
}
