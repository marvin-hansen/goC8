package kv_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForTruncateKVCollection(fabric, collectionName string) *RequestForTruncateKVCollection {
	return &RequestForTruncateKVCollection{
		path: fmt.Sprintf("_fabric/%v/_api/kv/%v/truncate", fabric, collectionName),
	}
}

type RequestForTruncateKVCollection struct {
	path string
}

func (req *RequestForTruncateKVCollection) Path() string {
	return req.path
}

func (req *RequestForTruncateKVCollection) Method() string {
	return http.MethodPut
}

func (req *RequestForTruncateKVCollection) Query() string {
	return ""
}

func (req *RequestForTruncateKVCollection) HasQueryParameter() bool {
	return false
}

func (req *RequestForTruncateKVCollection) GetQueryParameter() string {
	return ""
}

func (req *RequestForTruncateKVCollection) Payload() []byte {
	return nil
}

func (req *RequestForTruncateKVCollection) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

// call NewKVResponse() instead
