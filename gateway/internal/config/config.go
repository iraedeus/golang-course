package config

import "os"

type Config struct {
	HTTPPort      string
	CollectorAddr string
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

	return Config{
		HTTPPort:      httpPort,
		CollectorAddr: collectorAddr,
	}
}
