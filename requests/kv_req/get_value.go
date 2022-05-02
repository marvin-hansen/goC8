package kv_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForGetValue(fabric, collectionName, key string) *RequestForGetValue {
	return &RequestForGetValue{
		path: fmt.Sprintf("_fabric/%v/_api/kv/%v/value/%v", fabric, collectionName, key),
	}
}

type RequestForGetValue struct {
	path string
}

func (req *RequestForGetValue) Path() string {
	return req.path
}

func (req *RequestForGetValue) Method() string {
	return http.MethodGet
}

func (req *RequestForGetValue) Query() string {
	return ""
}

func (req *RequestForGetValue) HasQueryParameter() bool {
	return false
}

func (req *RequestForGetValue) GetQueryParameter() string {
	return ""
}

func (req *RequestForGetValue) Payload() []byte {
	return nil
}

func (req *RequestForGetValue) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

// call NewKVPair() instead
