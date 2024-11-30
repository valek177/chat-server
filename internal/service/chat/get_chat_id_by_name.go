package chat

import (
	"context"
	"log"
)

func (s *serv) GetChatIdByName(ctx context.Context, chatname string) (int64, error) {
	var id int64
	log.Printf("get chat id service")

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.chatRepository.GetChatIdByName(ctx, chatname)
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
