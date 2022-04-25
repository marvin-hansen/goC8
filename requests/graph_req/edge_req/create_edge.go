package edge_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForCreateEdge(fabric string) *RequestForCreateEdge {
	// @FIXME: Add correct API path
	return &RequestForCreateEdge{
		path: fmt.Sprintf("_fabric/%v/_api/NAME", fabric),
	}
}

type RequestForCreateEdge struct {
	path string
}

func (req *RequestForCreateEdge) Path() string {
	return req.path
}

func (req *RequestForCreateEdge) Method() string {
	return http.MethodGet
}

func (req *RequestForCreateEdge) Query() string {
	return ""
}

func (req *RequestForCreateEdge) HasQueryParameter() bool {
	return false
}

func (req *RequestForCreateEdge) GetQueryParameter() string {
	return ""
}

func (req *RequestForCreateEdge) Payload() []byte {
	return nil
}

func (req *RequestForCreateEdge) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForCreateEdge() *ResponseForCreateEdge {
	return new(ResponseForCreateEdge)
}

type ResponseForCreateEdge struct {
	// @FIXME
	Field string
}

func (r *ResponseForCreateEdge) IsResponse() {}

func (r ResponseForCreateEdge) String() string {
	// @FIXME
	return fmt.Sprintf("Bootfile: %v", r.Field)
}
