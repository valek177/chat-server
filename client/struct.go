package client

import (
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/valek177/chat-server/grpc/pkg/chat_v1"
)

const address = "localhost:50071"

type ChatV1Client struct {
	conn *grpc.ClientConn
	C    chat_v1.ChatV1Client
}

func NewChatV1Client() (*ChatV1Client, error) {
	creds, err := credentials.NewClientTLSFromFile("/home/valek/microservices_course/auth/tls/service.pem", "")
	if err != nil {
		return nil, err
	}
	conn, err := grpc.NewClient(
		address,
		grpc.WithTransportCredentials(creds),
	)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &ChatV1Client{
		conn: conn,
		C:    chatClient(conn),
	}, nil
}

func (c *ChatV1Client) Close() {
	c.conn.Close()
}
