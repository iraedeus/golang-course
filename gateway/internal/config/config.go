package config

import (
	"os"
	"strconv"
)

type Config struct {
	HTTPPort               string
	CollectorAddr          string
	ShutdownTimeoutSeconds int
}

func Load() Config {
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	collectorAddr := os.Getenv("COLLECTOR_ADDR")
	if collectorAddr == "" {
		collectorAddr = "localhost:50051"
	}

	shutdownTimeout := 5
	if v := os.Getenv("SHUTDOWN_TIMEOUT_SECONDS"); v != "" {
		if parsed, err := strconv.Atoi(v); err == nil && parsed > 0 {
			shutdownTimeout = parsed
		}
	}

	return Config{
		HTTPPort:               httpPort,
		CollectorAddr:          collectorAddr,
		ShutdownTimeoutSeconds: shutdownTimeout,
	}
}
