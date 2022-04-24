package index_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForCreatePersistentIndex(fabric string) *RequestForCreatePersistentIndex {
	// @FIXME: Add correct API path
	return &RequestForCreatePersistentIndex{
		path: fmt.Sprintf("_fabric/%v/_api/NAME", fabric),
	}
}

type RequestForCreatePersistentIndex struct {
	path string
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
	return false
}

func (req *RequestForCreatePersistentIndex) GetQueryParameter() string {
	return ""
}

func (req *RequestForCreatePersistentIndex) Payload() []byte {
	return nil
}

func (req *RequestForCreatePersistentIndex) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForCreatePersistentIndex() *ResponseForCreatePersistentIndex {
	return new(ResponseForCreatePersistentIndex)
}

type ResponseForCreatePersistentIndex struct {
	// @FIXME
	Field string
}

func (r *ResponseForCreatePersistentIndex) IsResponse() {}

func (r ResponseForCreatePersistentIndex) String() string {
	// @FIXME
	return fmt.Sprintf("Bootfile: %v", r.Field)
}
