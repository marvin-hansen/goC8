package query_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForDeleteCursor(fabric string) *RequestForDeleteCursor {
  // @FIXME: Add correct API path
	return &RequestForDeleteCursor{
			path: fmt.Sprintf("_fabric/%v/_api/NAME", fabric),
	}
}

type RequestForDeleteCursor struct {
	path string
}

func (req *RequestForDeleteCursor) Path() string {
	return req.path
}

func (req *RequestForDeleteCursor) Method() string {
	return http.MethodGet
}

func (req *RequestForDeleteCursor) Query() string {
	return ""
}

func (req *RequestForDeleteCursor) HasQueryParameter() bool {
	return false
}

func (req *RequestForDeleteCursor) GetQueryParameter() string {
	return "" //"?excludeSystem=true"
}

func (req *RequestForDeleteCursor) Payload() []byte {
	return nil
}

func (req *RequestForDeleteCursor) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForDeleteCursor() *ResponseForDeleteCursor {
	return new(ResponseForDeleteCursor)
}

type ResponseForDeleteCursor struct {
  // @FIXME
	Field string 
}

func (r *ResponseForDeleteCursor) IsResponse() {}

func (r ResponseForDeleteCursor) String() string {
  // @FIXME
	return fmt.Sprintf("Bootfile: %v", r.Field)
}


