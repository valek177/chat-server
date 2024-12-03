package chat

import (
	"log"

	"github.com/valek177/chat-server/grpc/pkg/chat_v1"
)

// ConnectChat connects user to chat
func (i *Implementation) ConnectChat(req *chat_v1.ConnectChatRequest,
	stream chat_v1.ChatV1_ConnectChatServer,
) error {
	log.Printf("User %s connecting to chat %s", req.Username, req.Chatname)

	chatID, err := i.chatService.GetChatIDByName(stream.Context(), req.Chatname)
	if err != nil {
		return err
	}

	err = i.chatService.ConnectChat(stream.Context(), chatID, req.Username, stream)
	if err != nil {
		return err
	}

	return nil
}
