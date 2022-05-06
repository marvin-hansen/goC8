package goC8

import (
	"github.com/valyala/fasthttp"
	"time"
)

const (
	debug     = false
	benchmark = false
	version   = "v1.4"
)

type Client struct {
	config      *ClientConfig
	Endpoint    string
	HTTPC       *fasthttp.Client
	HTTPTimeout time.Duration
	//
	Collection *CollectionManager
	Document   *DocumentManager
	Graph      *GraphManager
	Index      *IndexManager
	KV         *KVManager
	Query      *QueryManager
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

	// construct API managers
	c.Collection = NewCollectionManager(c)
	c.Document = NewDocumentManager(c)
	c.Graph = NewGraphManager(c)
	c.Index = NewIndexManager(c)
	c.KV = NewKVManager(c)
	c.Query = NewQueryManagerManager(c)

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
