package chat

import (
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/valek177/chat-server/grpc/pkg/chat_v1"
)

func (i *Implementation) ConnectChat(req *chat_v1.ConnectChatRequest, stream chat_v1.ChatV1_ConnectChatServer) error {
	log.Print("we are connecting to chat", req.ChatId, req.Username)
	i.mxChannel.RLock()
	chatChan, ok := i.channels[req.GetChatId()]
	i.mxChannel.RUnlock()

	if !ok {
		return status.Errorf(codes.NotFound, "chat not found")
	}

	i.mxChat.Lock()
	// Create new chat in map if it doesnt exist
	if _, okChat := i.chats[req.GetChatId()]; !okChat {
		i.chats[req.GetChatId()] = &Chat{
			streams: make(map[int64]chat_v1.ChatV1_ConnectChatServer),
		}
	}
	i.mxChat.Unlock()

	i.chats[req.GetChatId()].m.Lock()
	i.chats[req.GetChatId()].streams[1] = stream // req.GetUsername()] = stream
	i.chats[req.GetChatId()].m.Unlock()

	for {
		select {
		case msg, okCh := <-chatChan:
			if !okCh {
				return nil
			}

			for _, st := range i.chats[req.GetChatId()].streams {
				if err := st.Send(msg); err != nil {
					return err
				}
			}

		case <-stream.Context().Done():
			i.chats[req.GetChatId()].m.Lock()
			delete(i.chats[req.GetChatId()].streams, 1) // req.GetUsername())
			i.chats[req.GetChatId()].m.Unlock()

			log.Printf("context in connect chat is done")
			return nil
		}
	}
}
