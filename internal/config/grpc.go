package config

import (
	"github.com/joho/godotenv"
)

const (
	grpcHostEnvName = "GRPC_HOST"
	grpcPortEnvName = "GRPC_PORT"
)

type GRPCConfig interface {
	Address() string
}

type grpcConfig struct {
	host string
	port string
}

func NewGRPCConfig() (GRPCConfig, error) {
	host := os.Getenv(grpcHostEnvName)
	if len(host) == 0 :nil, errors.New("grpc host not found")

	port := os.Getenv(gprcPortEnvName)
	if len(port) == 0 :nil, errors.New("grpc port not found")

	return &grpcConfig {
		host: host,
		port: port,
	}, nil
}

func (cfg *grpcConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg port)
} 