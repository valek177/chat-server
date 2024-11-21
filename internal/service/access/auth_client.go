package access

import (
	"context"

	"google.golang.org/grpc"

	"github.com/valek177/auth/grpc/pkg/access_v1"
	"github.com/valek177/chat-server/internal/service"
)

// AuthClient is interface for auth client
type AuthClient interface {
	IsAccessGranted(ctx context.Context, endpoint string) (bool, error)
}

type authClient struct {
	client        access_v1.AccessV1Client
	accessService service.AccessService
}

// NewAuthClient returns new AuthClient
func NewAuthClient(conn *grpc.ClientConn, accessService service.AccessService) *authClient {
	return &authClient{
		client:        access_v1.NewAccessV1Client(conn),
		accessService: accessService,
	}
}

// IsAccessGranted checks access is granted by access token
func (c *authClient) IsAccessGranted(ctx context.Context, endpoint string) (bool, error) {
	ctx, err := c.accessService.TokenCtx(ctx)
	if err != nil {
		return false, err
	}

	_, err = c.client.Check(ctx, &access_v1.CheckRequest{
		EndpointAddress: endpoint,
	})
	if err != nil {
		return false, err
	}

	return true, nil
}
