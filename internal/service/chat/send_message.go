package chat

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/valek177/chat-server/grpc/pkg/chat_v1"
)

// SendMessage sends message to chat
func (s *serv) SendMessage(_ context.Context, chatID int64, message *chat_v1.Message) error {
	s.mxChannel.RLock()
	chatChan, ok := s.channels[chatID]
	s.mxChannel.RUnlock()

	if !ok {
		return status.Errorf(codes.NotFound, "chat not found")
	}

	chatChan <- message

	return nil
}
