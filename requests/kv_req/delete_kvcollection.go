package kv_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForDeleteKVCollection(fabric, collectionName string) *RequestForDeleteKVCollection {
	return &RequestForDeleteKVCollection{
		path: fmt.Sprintf("_fabric/%v/_api/kv/%v", fabric, collectionName),
	}
}

type RequestForDeleteKVCollection struct {
	path string
}

func (req *RequestForDeleteKVCollection) Path() string {
	return req.path
}

func (req *RequestForDeleteKVCollection) Method() string {
	return http.MethodDelete
}

func (req *RequestForDeleteKVCollection) Query() string {
	return ""
}

func (req *RequestForDeleteKVCollection) HasQueryParameter() bool {
	return false
}

func (req *RequestForDeleteKVCollection) GetQueryParameter() string {
	return ""
}

func (req *RequestForDeleteKVCollection) Payload() []byte {
	return nil
}

func (req *RequestForDeleteKVCollection) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

// call NewKVResponse() instead
