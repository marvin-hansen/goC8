package index_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForCreateHashIndex(fabric string) *RequestForCreateHashIndex {
  // @FIXME: Add correct API path
	return &RequestForCreateHashIndex{
			path: fmt.Sprintf("_fabric/%v/_api/NAME", fabric),
	}
}

type RequestForCreateHashIndex struct {
	path string
}

func (req *RequestForCreateHashIndex) Path() string {
	return req.path
}

func (req *RequestForCreateHashIndex) Method() string {
	return http.MethodGet
}

func (req *RequestForCreateHashIndex) Query() string {
	return ""
}

func (req *RequestForCreateHashIndex) HasQueryParameter() bool {
	return false
}

func (req *RequestForCreateHashIndex) GetQueryParameter() string {
	return ""
}

func (req *RequestForCreateHashIndex) Payload() []byte {
	return nil
}

func (req *RequestForCreateHashIndex) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForCreateHashIndex() *ResponseForCreateHashIndex {
	return new(ResponseForCreateHashIndex)
}

type ResponseForCreateHashIndex struct {
  // @FIXME
	Field string 
}

func (r *ResponseForCreateHashIndex) IsResponse() {}

func (r ResponseForCreateHashIndex) String() string {
  // @FIXME
	return fmt.Sprintf("Bootfile: %v", r.Field)
}


