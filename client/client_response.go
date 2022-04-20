package client

import jsoniter "github.com/json-iterator/go"

type Response struct {
	Result     interface{} `json:"result,omitempty"`
	RawMessage jsoniter.RawMessage
	Error      string `json:"error,omitempty"`
	Success    bool   `json:"success"`
}

func (c *Client) newResponse() *Response {
	return &Response{}
}
