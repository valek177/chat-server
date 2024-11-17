package env

import (
	"net"
	"os"

	"github.com/pkg/errors"

	"github.com/valek177/chat-server/internal/config"
)

var _ config.GRPCConfig = (*grpcConfig)(nil)

const (
	grpcHostEnvName    = "GRPC_HOST"
	grpcPortEnvName    = "GRPC_PORT"
	serviceTLSCertFile = "GRPC_TLS_CERT_FILE"
	serviceTLSKeyFile  = "GRPC_TLS_KEY_FILE"
)

type grpcConfig struct {
	host        string
	port        string
	tlsCertFile string
	tlsKeyFile  string
}

// NewGRPCConfig creates new grpcConfig
func NewGRPCConfig() (*grpcConfig, error) {
	host := os.Getenv(grpcHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("grpc host not found")
	}

	port := os.Getenv(grpcPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("grpc port not found")
	}

	tlsServiceCertFile := os.Getenv(serviceTLSCertFile)
	if tlsServiceCertFile == "" {
		return nil, errors.New("grpc tls cert file not found")
	}

	tlsServiceKeyFile := os.Getenv(serviceTLSKeyFile)
	if tlsServiceKeyFile == "" {
		return nil, errors.New("grpc tls key file not found")
	}

	return &grpcConfig{
		host:        host,
		port:        port,
		tlsCertFile: tlsServiceCertFile,
		tlsKeyFile:  tlsServiceKeyFile,
	}, nil
}

// Address returns address from config
func (cfg *grpcConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}

// TLSCertFile returns path to TLS cert file from config
func (cfg *grpcConfig) TLSCertFile() string {
	return cfg.tlsCertFile
}

// TLSKeyFile returns path to TLS key file from config
func (cfg *grpcConfig) TLSKeyFile() string {
	return cfg.tlsKeyFile
}
