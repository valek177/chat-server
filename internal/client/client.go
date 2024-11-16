package client

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/valek177/auth/grpc/pkg/access_v1"
)

type AuthClient interface {
	Check(ctx context.Context, endpoint string) (bool, error)
}

type authClient struct {
	client access_v1.AccessV1Client
}

func NewAuthClient(conn *grpc.ClientConn) *authClient {
	return &authClient{
		client: access_v1.NewAccessV1Client(conn),
	}
}

func (c *authClient) Check(ctx context.Context, endpoint string) (bool, error) {
	mdin, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("err in in")
	}
	authHeader := mdin.Get("authorization")
	if len(authHeader) == 0 {
		fmt.Println("err in in header")
		return false, fmt.Errorf("header not provided")
	}

	if !strings.HasPrefix(authHeader[0], "Bearer ") {
		fmt.Println("err in prefix")
		return false, fmt.Errorf("invalid header format")
	}
	accessToken := strings.TrimPrefix(authHeader[0], "Bearer ")

	fmt.Println("access token", accessToken)

	md := metadata.New(map[string]string{"Authorization": "Bearer " + accessToken})
	ctx = metadata.NewOutgoingContext(ctx, md)

	// endpoint = "test"
	fmt.Println("send request ", endpoint, md)
	_, err := c.client.Check(ctx, &access_v1.CheckRequest{
		EndpointAddress: endpoint,
	})
	if err != nil {
		fmt.Println("err in INTERC", err)
		return false, err
	}

	return true, nil
}
