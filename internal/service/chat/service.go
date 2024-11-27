package chat

import (
	"sync"

	"github.com/valek177/chat-server/grpc/pkg/chat_v1"
	"github.com/valek177/chat-server/internal/repository"
	"github.com/valek177/chat-server/internal/service"
	"github.com/valek177/platform-common/pkg/client/db"
)

type Chat struct {
	streams map[string]chat_v1.ChatV1_ConnectChatServer
	m       sync.RWMutex
}

type serv struct {
	chatRepository repository.ChatRepository
	logRepository  repository.LogRepository
	txManager      db.TxManager

	chats  map[int64]*Chat
	mxChat sync.RWMutex

	channels  map[int64]chan *chat_v1.Message
	mxChannel sync.RWMutex
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
		chats:          make(map[int64]*Chat),
		channels:       make(map[int64]chan *chat_v1.Message),
	}
}
