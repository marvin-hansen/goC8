package kv_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForDeleteValue(fabric, collectionName, key string) *RequestForDeleteValue {
	return &RequestForDeleteValue{
		path: fmt.Sprintf("_fabric/%v/_api/kv/%v/value/%v", fabric, collectionName, key),
	}
}

type RequestForDeleteValue struct {
	path string
}

func (req *RequestForDeleteValue) Path() string {
	return req.path
}

func (req *RequestForDeleteValue) Method() string {
	return http.MethodDelete
}

func (req *RequestForDeleteValue) Query() string {
	return ""
}

func (req *RequestForDeleteValue) HasQueryParameter() bool {
	return false
}

func (req *RequestForDeleteValue) GetQueryParameter() string {
	return ""
}

func (req *RequestForDeleteValue) Payload() []byte {
	return nil
}

func (req *RequestForDeleteValue) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

// call NewEmptyKVPair() instead
