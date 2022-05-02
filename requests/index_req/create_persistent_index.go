package index_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForCreatePersistentIndex(fabric, collectionName, field string, deduplicate, sparse, unique bool) *RequestForCreatePersistentIndex {
	indexType := "persistent"
	return &RequestForCreatePersistentIndex{
		path:       fmt.Sprintf("_fabric/%v/_api/index/%v", fabric, indexType),
		parameters: fmt.Sprintf("?collection=%v", collectionName),
		payload:    getIndexPayLoad(indexType, field, deduplicate, sparse, unique),
	}
}

type RequestForCreatePersistentIndex struct {
	path       string
	parameters string
	payload    []byte
}

func (req *RequestForCreatePersistentIndex) Path() string {
	return req.path
}

func (req *RequestForCreatePersistentIndex) Method() string {
	return http.MethodPost
}

func (req *RequestForCreatePersistentIndex) Query() string {
	return ""
}

func (req *RequestForCreatePersistentIndex) HasQueryParameter() bool {
	return true
}

func (req *RequestForCreatePersistentIndex) GetQueryParameter() string {
	return req.parameters
}

func (req *RequestForCreatePersistentIndex) Payload() []byte {
	return req.payload
}

func (req *RequestForCreatePersistentIndex) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForCreatePersistentIndex() *IndexEntry {
	return new(IndexEntry)
}
