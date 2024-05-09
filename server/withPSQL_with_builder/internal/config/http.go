package config

import (
	"errors"
	"os"
)

const (
	httpPortEnvName = "HTTP_PORT"
)

type httpConfig struct {
	port string
}

func NewHTTPConfig() (HTTPConfig, error) {
	port := os.Getenv(httpPortEnvName)

	if len(port) == 0 {
		return nil, errors.New("http port not found")
	}

	return &httpConfig{
		port: port,
	}, nil
}

func (cfg *httpConfig) Port() string {
	return cfg.port
}
