package config

import (
	"os"
)

type ServerConfig struct {
	AppName              string
	NATSStreamingUrl     string
	NATSStreamingCluster string
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

	return ServerConfig{
		AppName:              serviceName,
		NATSStreamingUrl:     natsStreamingUrl,
		NATSStreamingCluster: natsStreamingCluster,
	}
}
