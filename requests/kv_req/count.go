package kv_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForCount(fabric, collectionName string) *RequestForCount {
	return &RequestForCount{
		path: fmt.Sprintf("_fabric/%v/_api/kv/%v/count", fabric, collectionName),
	}
}

type RequestForCount struct {
	path string
}

func (req *RequestForCount) Path() string {
	return req.path
}

func (req *RequestForCount) Method() string {
	return http.MethodGet
}

func (req *RequestForCount) Query() string {
	return ""
}

func (req *RequestForCount) HasQueryParameter() bool {
	return false
}

func (req *RequestForCount) GetQueryParameter() string {
	return ""
}

func (req *RequestForCount) Payload() []byte {
	return nil
}

func (req *RequestForCount) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

// call NewKVResponse() instead
