package vault

import (
	"fmt"
	"os"
	"strings"

	"github.com/hashicorp/vault/api"
)

// Provider ...
type Provider struct {
	path    string
	client  *api.Logical
	results map[string]map[string]string
}

// New ...
func New(token, addr, path string) (*Provider, error) {
	config := &api.Config{
		Address: addr,
	}

	client, err := api.NewClient(config)
	if err != nil {
		return nil, fmt.Errorf("api.NewClient")
	}

	client.SetToken(token)

	return &Provider{
		path:    path,
		client:  client.Logical(),
		results: make(map[string]map[string]string),
	}, nil
}

// Get retrieves a value from vault using the KV engine. The actual key selected is determined by the value
// separated by the colon. For example "database:password" will retrieve the key "password" from the path
// "database".
func (p *Provider) Get(v string) (string, error) {
	// <path>/data/<path-secret>:key
	split := strings.Split(v, ":")
	if len(split) == 1 {
		return "", fmt.Errorf("missing key value")
	}

	pathSecret := split[0]
	key := split[1]

	res, ok := p.results[pathSecret]
	if ok {
		val, ok := res[key]
		if !ok {
			return "", fmt.Errorf("key not found in cached data")
		}

		return val, nil
	}

	secret, err := p.client.Read(fmt.Sprintf("%s/data/%s", p.path, pathSecret))
	if err != nil {
		return "", fmt.Errorf("reading")
	}

	if secret == nil {
		return "", fmt.Errorf("secret not found")
	}

	data, ok := secret.Data["data"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("invalid data in secret")
	}

	secrets := make(map[string]string)

	for k, v := range data {
		val, ok := v.(string)
		if !ok {
			return "", fmt.Errorf("secret value in data is not string")
		}

		secrets[k] = val
	}

	val, ok := secrets[key]
	if !ok {
		return "", fmt.Errorf("key not found in retrieved data")
	}

	p.results[pathSecret] = secrets

	return val, nil
}

// NewVaultProvider instantiates the Vault client using configuration defined in environment variables.
func NewVaultProvider() (*Provider, error) {
	vaultPath := os.Getenv("VAULT_PATH")
	vaultToken := os.Getenv("VAULT_TOKEN")
	vaultAddress := os.Getenv("VAULT_ADDRESS")

	provider, err := New(vaultToken, vaultAddress, vaultPath)
	if err != nil {
		return nil, fmt.Errorf("vault.New ")
	}

	return provider, nil
}
