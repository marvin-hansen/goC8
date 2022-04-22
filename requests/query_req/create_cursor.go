package query_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForCreateCursor(fabric string) *RequestForCreateCursor {
  // @FIXME: Add correct API path
	return &RequestForCreateCursor{
			path: fmt.Sprintf("_fabric/%v/_api/NAME", fabric),
	}
}

type RequestForCreateCursor struct {
	path string
}

func (req *RequestForCreateCursor) Path() string {
	return req.path
}

func (req *RequestForCreateCursor) Method() string {
	return http.MethodGet
}

func (req *RequestForCreateCursor) Query() string {
	return ""
}

func (req *RequestForCreateCursor) HasQueryParameter() bool {
	return false
}

func (req *RequestForCreateCursor) GetQueryParameter() string {
	return "" //"?excludeSystem=true"
}

func (req *RequestForCreateCursor) Payload() []byte {
	return nil
}

func (req *RequestForCreateCursor) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForCreateCursor() *ResponseForCreateCursor {
	return new(ResponseForCreateCursor)
}

type ResponseForCreateCursor struct {
  // @FIXME
	Field string 
}

func (r *ResponseForCreateCursor) IsResponse() {}

func (r ResponseForCreateCursor) String() string {
  // @FIXME
	return fmt.Sprintf("Bootfile: %v", r.Field)
}


