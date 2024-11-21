package access

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/metadata"

	"github.com/valek177/chat-server/internal/service"
)

const (
	authPrefix     = "Bearer "
	authHeaderName = "authorization"
)

type serv struct{}

// NewService creates new service with settings
func NewService() service.AccessService {
	return &serv{}
}

// TokenCtx returns new context with access token
func (s *serv) TokenCtx(ctx context.Context) (context.Context, error) {
	mdin, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("unable to get metadata from incoming context")
	}

	authHeader := mdin.Get(authHeaderName)
	if len(authHeader) == 0 {
		return nil, fmt.Errorf("header not provided")
	}

	if !strings.HasPrefix(authHeader[0], authPrefix) {
		return nil, fmt.Errorf("invalid header format")
	}
	accessToken := strings.TrimPrefix(authHeader[0], authPrefix)

	md := metadata.New(map[string]string{authHeaderName: authPrefix + accessToken})

	return metadata.NewOutgoingContext(ctx, md), nil
}
