package config

import (
	"github.com/joho/godotenv"
)

// Load loads environment
func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}

// GRPCConfig interface for GRPCConfig
type GRPCConfig interface {
	Address() string
	TLSCertFile() string
	TLSKeyFile() string
}

// PGConfig interface for PGConfig
type PGConfig interface {
	DSN() string
}

// GRPCAuthConfig interface for GRPCAuthConfig
type GRPCAuthConfig interface {
	Address() string
}

// JaegerConfig interface for jaeger config
type JaegerConfig interface {
	LocalAgentAddress() string
	SamplerType() string
	SamplerParam() float64
	ServiceName() string
}
