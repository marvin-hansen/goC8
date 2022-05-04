package index_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForDeleteIndex(fabric, collectionName, indexName string) *RequestForDeleteIndex {
	return &RequestForDeleteIndex{
		path: fmt.Sprintf("_fabric/%v/_api/index/%v/%v", fabric, collectionName, indexName),
	}
}

type RequestForDeleteIndex struct {
	path string
}

func (req *RequestForDeleteIndex) Path() string {
	return req.path
}

func (req *RequestForDeleteIndex) Method() string {
	return http.MethodGet
}

func (req *RequestForDeleteIndex) Query() string {
	return ""
}

func (req *RequestForDeleteIndex) HasQueryParameter() bool {
	return false
}

func (req *RequestForDeleteIndex) GetQueryParameter() string {
	return ""
}

func (req *RequestForDeleteIndex) Payload() []byte {
	return nil
}

func (req *RequestForDeleteIndex) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForDeleteIndex() *ResponseForDeleteIndex {
	return new(ResponseForDeleteIndex)
}

type ResponseForDeleteIndex struct {
	Code  int    `json:"code,omitempty"`
	Error bool   `json:"error,omitempty"`
	Id    string `json:"id,omitempty"`
}

func (r *ResponseForDeleteIndex) IsResponse() {}

func (r ResponseForDeleteIndex) String() string {
	return fmt.Sprintf("Code: %v Error: %v ID: %v",
		r.Code,
		r.Error,
		r.Id,
	)
}
