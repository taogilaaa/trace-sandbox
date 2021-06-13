package config

import (
	"os"
)

type ServerConfig struct {
	AppName    string
	AppVersion string
	GRPCUrl    string
}

func Load() ServerConfig {
	serviceName := "http"

	grpcUrl := "localhost:50041"
	if value, ok := os.LookupEnv("GRPC_URL"); ok {
		grpcUrl = value
	}

	return ServerConfig{
		AppName:    serviceName,
		AppVersion: "1",
		GRPCUrl:    grpcUrl,
	}
}
