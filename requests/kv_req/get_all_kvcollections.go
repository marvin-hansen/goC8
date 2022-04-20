package kv_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForGetAllKVCollections(fabric string) *RequestForGetAllKVCollections {
	return &RequestForGetAllKVCollections{
		path: fmt.Sprintf("_fabric/%v/_api/kv", fabric),
	}
}

type RequestForGetAllKVCollections struct {
	path string
}

func (req *RequestForGetAllKVCollections) Path() string {
	return req.path
}

func (req *RequestForGetAllKVCollections) Method() string {
	return http.MethodGet
}

func (req *RequestForGetAllKVCollections) Query() string {
	return ""
}

func (req *RequestForGetAllKVCollections) HasQueryParameter() bool {
	return false
}

func (req *RequestForGetAllKVCollections) GetQueryParameter() string {
	return ""
}

func (req *RequestForGetAllKVCollections) Payload() []byte {
	return nil
}

func (req *RequestForGetAllKVCollections) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForGetAllKVCollections() *ResponseForGetAllKVCollections {
	return new(ResponseForGetAllKVCollections)
}

type ResponseForGetAllKVCollections struct {
	// @FIXME
	Field string
}

func (r *ResponseForGetAllKVCollections) IsResponse() {}

func (r ResponseForGetAllKVCollections) String() string {
	// @FIXME
	return fmt.Sprintf("Bootfile: %v", r.Field)
}
