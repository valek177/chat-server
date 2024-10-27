package chat

import (
	"github.com/valek177/chat-server/grpc/pkg/chat_v1"
	"github.com/valek177/chat-server/internal/service"
)

// Implementation struct contains server
type Implementation struct {
	chat_v1.UnimplementedChatV1Server
	chatService service.ChatService
}

// NewImplementation returns implementation object
func NewImplementation(chatService service.ChatService) *Implementation {
	return &Implementation{
		chatService: chatService,
	}
}
