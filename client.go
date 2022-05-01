package goC8

import (
	"github.com/valyala/fasthttp"
	"time"
)

const benchmark = false
const version = "v1.2"

type Client struct {
	config      *ClientConfig
	Endpoint    string
	HTTPC       *fasthttp.Client
	HTTPTimeout time.Duration
}

func NewClient(config *ClientConfig) *Client {

	if config == nil {
		panic("No config provided!")
	}

	timeOut := time.Duration(config.Timeout) * time.Second

	return &Client{
		config:      config,
		Endpoint:    config.GetConnectionString(),
		HTTPTimeout: timeOut,
		HTTPC:       new(fasthttp.Client),
	}
}

// getApiKey is used internally to pass key to the request handler
func (c Client) getApiKey() string {
	return c.config.GetApiKey()
}

func (c Client) getQueryTTL() int {
	return c.config.GetQueryTTL()
}

func (c Client) Info() {
	println("Golang driver for the Macrometa global data platform (GDN). "+
		"\n Version: %v", version)
}

func (c Client) Version() string {
	return version
}
