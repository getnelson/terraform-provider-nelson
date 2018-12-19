package nelson

import (
	"net/http"
	"time"

	cleanhttp "github.com/hashicorp/go-cleanhttp"
)

// Config is used to configure the creation of the client
type Config struct {

	// Address is the address of the Nelson server
	Address string

	// Token is the Github Auth token to use when communicating with Nelson
	Token   string
}

// DefaultConfig returns a default configuration for the client
func DefaultConfig() *Config {
	config := &Config{
		Address:    "https://127.0.0.1:9000",
		HttpClient: cleanhttp.DefaultPooledClient(),
	}

	return config
}

type Client struct {
	addr   *url.URL
	config *Config
}

func NewClient(c *Config) (*Client, error) {
	def := DefaultConfig()

	if c == nil {
		c = def
	}

	if c.HttpClient == nil {
		c.HttpClient = def.HttpClient
	}

	client := &Client{
		addr: u,
		config: c
	}

	transport := &http.Client{
		Timeout: time.Second * 30, // todo pull from terraform timeout config
	}

	return transport, nil
}
