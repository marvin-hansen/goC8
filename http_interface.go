package goC8

type Requester interface {
	Path() string              // path relative to URI i.e. /data/function/
	Method() string            // http method i.e. get, post, put, delete etc
	Query() string             // query to append after ? i.e. query?"my query"
	Payload() []byte           // payload i.e. file to upload
	ResponseCode() int         // expected http status code the server should normally return.
	HasQueryParameter() bool   // returns true if query parameters needed to be added to the path
	GetQueryParameter() string // returns of key / value pairs of query parameters
}

type Responder interface {
	IsResponse()
	String() string
}

type JsonResponder interface {
	IsJsonResponse()
	GetJsonMessage() []byte
	SetJsonMessage(raw []byte)
	String() string
}
