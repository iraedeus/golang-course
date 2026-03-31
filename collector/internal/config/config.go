package config

import (
	"os"
	"strconv"
)

type Config struct {
	GRPCPort               string
	ShutdownTimeoutSeconds int
}

func Load() Config {
	port := os.Getenv("GRPC_PORT")
	if port == "" {
		port = "50051"
	}

	shutdownTimeout := 5
	if v := os.Getenv("SHUTDOWN_TIMEOUT_SECONDS"); v != "" {
		if parsed, err := strconv.Atoi(v); err == nil && parsed > 0 {
			shutdownTimeout = parsed
		}
	}

	return Config{
		GRPCPort:               port,
		ShutdownTimeoutSeconds: shutdownTimeout,
	}
}
