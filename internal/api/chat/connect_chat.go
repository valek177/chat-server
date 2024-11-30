package chat

import (
	"log"

	"github.com/valek177/chat-server/grpc/pkg/chat_v1"
)

func (i *Implementation) ConnectChat(req *chat_v1.ConnectChatRequest,
	stream chat_v1.ChatV1_ConnectChatServer,
) error {
	// check user can connect to this chat (permissions)
	log.Print("we are connecting to chat", req.Chatname, req.Username)

	chatID, err := i.chatService.GetChatIdByName(stream.Context(), req.Chatname)
	if err != nil {
		return err
	}

	err = i.chatService.ConnectChat(stream.Context(), chatID, req.Username, stream)
	if err != nil {
		return err
	}

	return nil
}
