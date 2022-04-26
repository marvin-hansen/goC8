package edge_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForDeleteEdge(fabric string) *RequestForDeleteEdge {
  // @FIXME: Add correct API path
	return &RequestForDeleteEdge{
			path: fmt.Sprintf("_fabric/%v/_api/NAME", fabric),
	}
}

type RequestForDeleteEdge struct {
	path string
}

func (req *RequestForDeleteEdge) Path() string {
	return req.path
}

func (req *RequestForDeleteEdge) Method() string {
	return http.MethodGet
}

func (req *RequestForDeleteEdge) Query() string {
	return ""
}

func (req *RequestForDeleteEdge) HasQueryParameter() bool {
	return false
}

func (req *RequestForDeleteEdge) GetQueryParameter() string {
	return ""
}

func (req *RequestForDeleteEdge) Payload() []byte {
	return nil
}

func (req *RequestForDeleteEdge) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForDeleteEdge() *ResponseForDeleteEdge {
	return new(ResponseForDeleteEdge)
}

type ResponseForDeleteEdge struct {
  // @FIXME
	Field string 
}

func (r *ResponseForDeleteEdge) IsResponse() {}

func (r ResponseForDeleteEdge) String() string {
  // @FIXME
	return fmt.Sprintf("Bootfile: %v", r.Field)
}


