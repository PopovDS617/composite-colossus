package env

import (
	"fmt"
	"os"
)

type Provider interface {
	Get(key string) (string, error)
}

type Configuration struct {
	provider Provider
}

// New ...
func New(provider Provider) *Configuration {
	return &Configuration{
		provider: provider,
	}
}

// Get returns the value from environment variable `<key>`. When an environment variable `<key>_SECURE` exists
// the provider is used for getting the value.
func (c *Configuration) Get(key string) (string, error) {
	res := os.Getenv(key)
	valSecret := os.Getenv(key + "_SECURE")

	if valSecret != "" {
		valSecretRes, err := c.provider.Get(valSecret)
		if err != nil {
			return "", fmt.Errorf("provider.Get")
		}

		res = valSecretRes
	}

	return res, nil
}
