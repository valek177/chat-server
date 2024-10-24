package chat

import (
	"context"

	"github.com/valek177/chat-server/grpc/pkg/chat_v1"
)

// CreateChat creates new chat
func (s *serv) CreateChat(ctx context.Context, req *chat_v1.CreateChatRequest) (int64, error) {
	var id int64

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.chatRepository.CreateChat(ctx, req)
		if errTx != nil {
			return errTx
		}

		return nil
	})
	if err != nil {
		return 0, err
	}

	return id, nil
}
