package src

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
		config = NewClientDefaultConfig()
	}

	timeOut := time.Duration(config.Timeout) * time.Second

	return &Client{
		apiKey:      apiKEY,
		Endpoint:    config.GetConnectionString(),
		HTTPTimeout: timeOut,
		HTTPC:       new(fasthttp.Client),
	}
}

func (c Client) getApiKey() string {
	return c.apiKey
}
