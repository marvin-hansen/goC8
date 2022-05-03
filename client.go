package goC8

import (
	"github.com/valyala/fasthttp"
	"time"
)

const (
	debug     = false
	benchmark = false
	version   = "v1.2"
)

type Client struct {
	config      *ClientConfig
	Endpoint    string
	HTTPC       *fasthttp.Client
	HTTPTimeout time.Duration
	//
	KV *KVManager
}

func NewClient(config *ClientConfig) *Client {

	if config == nil {
		panic("No config provided!")
	}

	timeOut := time.Duration(config.Timeout) * time.Second

	c := &Client{
		config:      config,
		Endpoint:    config.GetConnectionString(),
		HTTPTimeout: timeOut,
		HTTPC:       new(fasthttp.Client),
	}

	// construct KV manager
	c.KV = NewKVManager(c)

	return c
}

// getApiKey is used internally to pass key to the Request handler
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
