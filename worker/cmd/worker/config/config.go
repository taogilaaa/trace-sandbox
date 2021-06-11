package config

import (
	"os"
)

type ServerConfig struct {
	AppName              string
	AppVersion           string
	NATSStreamingUrl     string
	NATSStreamingCluster string
	GRPCUrl              string
}

func Load() ServerConfig {
	// TODO: Load from env injected on docker image tag ?
	serviceName := "worker"

	natsStreamingUrl := "nats://localhost:4222"
	if value, ok := os.LookupEnv("NATS_STREAMING_URL"); ok {
		natsStreamingUrl = value
	}

	natsStreamingCluster := "test-cluster"
	if value, ok := os.LookupEnv("NATS_CLUSTER"); ok {
		natsStreamingCluster = value
	}

	grpcUrl := "localhost:50041"
	if value, ok := os.LookupEnv("GRPC_URL"); ok {
		natsStreamingUrl = value
	}

	return ServerConfig{
		AppName:              serviceName,
		AppVersion:           "1",
		NATSStreamingUrl:     natsStreamingUrl,
		NATSStreamingCluster: natsStreamingCluster,
		GRPCUrl:              grpcUrl,
	}
}
