package query_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForReadNextCursor(fabric string) *RequestForReadNextCursor {
  // @FIXME: Add correct API path
	return &RequestForReadNextCursor{
			path: fmt.Sprintf("_fabric/%v/_api/NAME", fabric),
	}
}

type RequestForReadNextCursor struct {
	path string
}

func (req *RequestForReadNextCursor) Path() string {
	return req.path
}

func (req *RequestForReadNextCursor) Method() string {
	return http.MethodGet
}

func (req *RequestForReadNextCursor) Query() string {
	return ""
}

func (req *RequestForReadNextCursor) HasQueryParameter() bool {
	return false
}

func (req *RequestForReadNextCursor) GetQueryParameter() string {
	return "" //"?excludeSystem=true"
}

func (req *RequestForReadNextCursor) Payload() []byte {
	return nil
}

func (req *RequestForReadNextCursor) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForReadNextCursor() *ResponseForReadNextCursor {
	return new(ResponseForReadNextCursor)
}

type ResponseForReadNextCursor struct {
  // @FIXME
	Field string 
}

func (r *ResponseForReadNextCursor) IsResponse() {}

func (r ResponseForReadNextCursor) String() string {
  // @FIXME
	return fmt.Sprintf("Bootfile: %v", r.Field)
}


