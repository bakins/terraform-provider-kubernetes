package kubernetes

import (
	"errors"
	"log"
	"net/http"
)

type (
	Config struct {
		Endpoint string `mapstructure:"endpoint"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
	}

	Client struct {
		http.Client
		Config Config
	}
)

// Client() returns a new client for accessing kubernetes.
func (c *Config) Client() (*Client, error) {
	if c.Endpoint == "" {
		return nil, errors.New("endpoint is required")
	}

	client := &Client{
		Client: http.Client{},
		Config: *c,
	}

	log.Printf("[INFO] Kubernetes Client configured with endpoint: '%s'", c.Endpoint)

	return client, nil
}
