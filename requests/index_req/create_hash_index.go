package index_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForCreateHashIndex(fabric, collectionName, field string, deduplicate, sparse, unique bool) *RequestForCreateHashIndex {
	indexType := "hash"
	return &RequestForCreateHashIndex{
		path:       fmt.Sprintf("_fabric/%v/_api/index/%v", fabric, indexType),
		parameters: fmt.Sprintf("?collection=%v", collectionName),
		payload:    getIndexPayLoad(indexType, field, deduplicate, sparse, unique),
	}
}

type RequestForCreateHashIndex struct {
	path       string
	parameters string
	payload    []byte
}

func (req *RequestForCreateHashIndex) Path() string {
	return req.path
}

func (req *RequestForCreateHashIndex) Method() string {
	return http.MethodPost
}

func (req *RequestForCreateHashIndex) Query() string {
	return ""
}

func (req *RequestForCreateHashIndex) HasQueryParameter() bool {
	return true
}

func (req *RequestForCreateHashIndex) GetQueryParameter() string {
	return req.parameters
}

func (req *RequestForCreateHashIndex) Payload() []byte {
	return req.payload
}

func (req *RequestForCreateHashIndex) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForCreateHashIndex() *IndexEntry {
	return new(IndexEntry)
}
