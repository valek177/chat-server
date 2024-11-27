package chat

import (
	"sync"

	"github.com/valek177/chat-server/grpc/pkg/chat_v1"
	"github.com/valek177/chat-server/internal/service"
)

type Chat struct {
	streams map[int64]chat_v1.ChatV1_ConnectChatServer
	m       sync.RWMutex
}

// Implementation struct contains server
type Implementation struct {
	chat_v1.UnimplementedChatV1Server
	chatService service.ChatService

	chats  map[int64]*Chat
	mxChat sync.RWMutex

	channels  map[int64]chan *chat_v1.Message
	mxChannel sync.RWMutex
}

// NewImplementation returns implementation object
func NewImplementation(chatService service.ChatService) *Implementation {
	return &Implementation{
		chatService: chatService,
		chats:       make(map[int64]*Chat),
		channels:    make(map[int64]chan *chat_v1.Message),
	}
}
