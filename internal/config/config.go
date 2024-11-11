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
	TlsCertFile() string
	TlsKeyFile() string
}

// PGConfig interface for PGConfig
type PGConfig interface {
	DSN() string
}
