package config

import (
	"os"
)

type ServerConfig struct {
	AppName              string
	AppVersion           string
	GRPCUrl              string
	NATSStreamingUrl     string
	NATSStreamingCluster string
}

func Load() ServerConfig {
	serviceName := "http"

	grpcUrl := "localhost:50041"
	if value, ok := os.LookupEnv("GRPC_URL"); ok {
		grpcUrl = value
	}

	natsStreamingUrl := "nats://localhost:4222"
	if value, ok := os.LookupEnv("NATS_STREAMING_URL"); ok {
		natsStreamingUrl = value
	}

	natsStreamingCluster := "test-cluster"
	if value, ok := os.LookupEnv("NATS_CLUSTER"); ok {
		natsStreamingCluster = value
	}

	return ServerConfig{
		AppName:              serviceName,
		AppVersion:           "1",
		GRPCUrl:              grpcUrl,
		NATSStreamingUrl:     natsStreamingUrl,
		NATSStreamingCluster: natsStreamingCluster,
	}
}
