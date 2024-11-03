package chat

import (
	"github.com/valek177/chat-server/internal/repository"
	"github.com/valek177/chat-server/internal/service"
	"github.com/valek177/platform-common/pkg/client/db"
)

type serv struct {
	chatRepository repository.ChatRepository
	logRepository  repository.LogRepository
	txManager      db.TxManager
}

// NewService creates new service with settings
func NewService(
	chatRepository repository.ChatRepository,
	logRepository repository.LogRepository,
	txManager db.TxManager,
) service.ChatService {
	return &serv{
		chatRepository: chatRepository,
		logRepository:  logRepository,
		txManager:      txManager,
	}
}
