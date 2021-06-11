package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/taogilaaa/trace-sandbox/grpc/internal/database"
)

type ServerConfig struct {
	AppName        string
	AppNameVersion string
	AppVersion     string
	Database       database.DBConfig
}

func Load() ServerConfig {
	serviceName := "grpc"

	primaryDSN := fmt.Sprintf("%s application_name=%s", os.Getenv("DB_PRIMARY_DSN"), serviceName)
	replicaDSN := fmt.Sprintf("%s application_name=%s", os.Getenv("DB_REPLICA_DSN"), serviceName)
	maxIdleConn := 2
	if value, ok := os.LookupEnv("DB_MAX_IDLE_CONN"); ok {
		if value, err := strconv.Atoi(value); err == nil {
			maxIdleConn = value
		}
	}

	maxConn := 10
	if value, ok := os.LookupEnv("DB_MAX_CONN"); ok {
		if value, err := strconv.Atoi(value); err == nil {
			maxConn = value
		}
	}

	cml := 60
	if value, ok := os.LookupEnv("DB_CONN_MAX_LIFETIME"); ok {
		if value, err := strconv.Atoi(value); err == nil {
			cml = value
		}
	}
	connMaxLifetime := time.Duration(cml) * time.Second

	return ServerConfig{
		AppName:    serviceName,
		AppVersion: "1",
		Database: database.DBConfig{
			PrimaryDSN:      primaryDSN,
			ReplicaDSN:      replicaDSN,
			MaxConn:         maxConn,
			MaxIdleConn:     maxIdleConn,
			ConnMaxLifetime: connMaxLifetime,
		},
	}
}
