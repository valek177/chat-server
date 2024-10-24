package chat

import (
	"context"
)

// DeleteChat removes existing chat
func (s *serv) DeleteChat(ctx context.Context, id int64) error {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		errTx := s.chatRepository.DeleteChat(ctx, id)
		if errTx != nil {
			return errTx
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
