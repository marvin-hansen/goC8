package kv_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForDeleteKeyValue(fabric, collectionName string, keys KeyCollection) *RequestForDeleteKeyValue {
	return &RequestForDeleteKeyValue{
		path:    fmt.Sprintf("_fabric/%v/_api/kv/%v/values", fabric, collectionName),
		payload: keys.Json(),
	}
}

type RequestForDeleteKeyValue struct {
	path    string
	payload []byte
}

func (req *RequestForDeleteKeyValue) Path() string {
	return req.path
}

func (req *RequestForDeleteKeyValue) Method() string {
	return http.MethodDelete
}

func (req *RequestForDeleteKeyValue) Query() string {
	return ""
}

func (req *RequestForDeleteKeyValue) HasQueryParameter() bool {
	return false
}

func (req *RequestForDeleteKeyValue) GetQueryParameter() string {
	return ""
}

func (req *RequestForDeleteKeyValue) Payload() []byte {
	return req.payload
}

func (req *RequestForDeleteKeyValue) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//
