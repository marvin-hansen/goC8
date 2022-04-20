package goC8

import (
	"github.com/valyala/fasthttp"
	"time"
)

type Client struct {
	apiKey      string
	Endpoint    string
	HTTPC       *fasthttp.Client
	HTTPTimeout time.Duration
}

func NewClient(config *ClientConfig) *Client {

	if config == nil {
		config = NewDefaultConfig()
	}

	timeOut := time.Duration(config.Timeout) * time.Second

	return &Client{
		apiKey:      config.GetApiKey(),
		Endpoint:    config.GetConnectionString(),
		HTTPTimeout: timeOut,
		HTTPC:       new(fasthttp.Client),
	}
}

// getApiKey is used internally to pass key to the request handler
func (c Client) getApiKey() string {
	return c.apiKey
}

func (c Client) Info() {
	println("")
}
