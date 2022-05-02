package kv_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForSetKeyValue(fabric, collectionName string, kvPairs KVPairCollection) *RequestForSetKeyValue {
	return &RequestForSetKeyValue{
		path:    fmt.Sprintf("_fabric/%v/_api/kv/%v", fabric, collectionName),
		payload: kvPairs.Json(),
	}
}

type RequestForSetKeyValue struct {
	path    string
	payload []byte
}

func (req *RequestForSetKeyValue) Path() string {
	return req.path
}

func (req *RequestForSetKeyValue) Method() string {
	return http.MethodPut
}

func (req *RequestForSetKeyValue) Query() string {
	return ""
}

func (req *RequestForSetKeyValue) HasQueryParameter() bool {
	return false
}

func (req *RequestForSetKeyValue) GetQueryParameter() string {
	return ""
}

func (req *RequestForSetKeyValue) Payload() []byte {
	return req.payload
}

func (req *RequestForSetKeyValue) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

// call NewKVResponse() instead
