package env

import (
	"net"
	"os"

	"github.com/pkg/errors"

	"github.com/valek177/chat-server/internal/config"
)

var _ config.GRPCAuthConfig = (*grpcAuthConfig)(nil)

const (
	grpcAuthHostEnvName = "GRPC_AUTH_HOST"
	grpcAuthPortEnvName = "GRPC_AUTH_PORT"
)

type grpcAuthConfig struct {
	host string
	port string
}

// NewGRPCAuthConfig creates new grpcAuthConfig
func NewGRPCAuthConfig() (*grpcAuthConfig, error) {
	host := os.Getenv(grpcAuthHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("grpc auth host not found")
	}

	port := os.Getenv(grpcAuthPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("grpc auth port not found")
	}

	return &grpcAuthConfig{
		host: host,
		port: port,
	}, nil
}

// Address returns address from config
func (cfg *grpcAuthConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
