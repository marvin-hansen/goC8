package index_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForGetIndex(fabric, collectionName, indexName string) *RequestForGetIndex {
	return &RequestForGetIndex{
		path: fmt.Sprintf("_fabric/%v/_api/index/%v/%v", fabric, collectionName, indexName),
	}
}

type RequestForGetIndex struct {
	path string
}

func (req *RequestForGetIndex) Path() string {
	return req.path
}

func (req *RequestForGetIndex) Method() string {
	return http.MethodGet
}

func (req *RequestForGetIndex) Query() string {
	return ""
}

func (req *RequestForGetIndex) HasQueryParameter() bool {
	return false
}

func (req *RequestForGetIndex) GetQueryParameter() string {
	return ""
}

func (req *RequestForGetIndex) Payload() []byte {
	return nil
}

func (req *RequestForGetIndex) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForGetIndex() *IndexEntry {
	return new(IndexEntry)
}
