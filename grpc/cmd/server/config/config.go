package config

type ServerConfig struct {
	AppName        string
	AppNameVersion string
	AppVersion     string
}

func Load() ServerConfig {
	return ServerConfig{
		AppName:    "grpc",
		AppVersion: "1",
	}
}
