package interceptor

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	"github.com/valek177/chat-server/internal/service/access"
)

// Auth is a struct for auth client
type Auth struct {
	authClient access.AuthClient
}

// NewAuthInterceptor is interceptor for AuthClient
func NewAuthInterceptor(authClient access.AuthClient) *Auth {
	return &Auth{
		authClient: authClient,
	}
}

// Interceptor returns interceptor for AuthClient
func (a *Auth) Interceptor(_ context.Context) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp any, err error) {
		accessAllowed, err := a.authClient.IsAccessGranted(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}
		if !accessAllowed {
			return nil, fmt.Errorf("access denied")
		}

		return handler(ctx, req)
	}
}
