package chat

import (
	"context"
	"log"

	"github.com/valek177/chat-server/grpc/pkg/chat_v1"
	"github.com/valek177/chat-server/internal/converter"
)

// CreateChat creates new chat
func (s *serv) CreateChat(ctx context.Context, req *chat_v1.CreateChatRequest) (int64, error) {
	var id int64
	log.Printf("create chat service")

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.chatRepository.CreateChat(ctx, req)
		if errTx != nil {
			return errTx
		}

		_, errTx = s.logRepository.CreateRecord(ctx,
			converter.ToRecordRepoFromService(id, "create"))
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
