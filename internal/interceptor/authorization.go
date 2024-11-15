package interceptor

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	"github.com/valek177/chat-server/internal/client"
)

type Auth struct {
	authClient client.AuthClient
}

func NewAuthInterceptor(authClient client.AuthClient) *Auth {
	return &Auth{
		authClient: authClient,
	}
}

func (a *Auth) Interceptor(ctx context.Context) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp any, err error) {
		fmt.Println("unary interceptor ", info.FullMethod)
		_, err = a.authClient.Check(ctx, info.FullMethod)
		if err != nil {
			fmt.Println("err in interceptor ", err)
			return nil, err
		}

		return handler(ctx, req)
	}
}
