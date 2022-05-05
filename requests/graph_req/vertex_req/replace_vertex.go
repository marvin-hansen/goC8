package vertex_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForReplaceVertex(fabric string) *RequestForReplaceVertex {
	// @FIXME: Add correct API path
	return &RequestForReplaceVertex{
		path: fmt.Sprintf("_fabric/%v/_api/NAME", fabric),
	}
}

type RequestForReplaceVertex struct {
	path       string
	parameters string
	payload    []byte
}

func (req *RequestForReplaceVertex) Path() string {
	return req.path
}

func (req *RequestForReplaceVertex) Method() string {
	return http.MethodPut
}

func (req *RequestForReplaceVertex) Query() string {
	return ""
}

func (req *RequestForReplaceVertex) HasQueryParameter() bool {
	return true
}

func (req *RequestForReplaceVertex) GetQueryParameter() string {
	return req.parameters
}

func (req *RequestForReplaceVertex) Payload() []byte {
	return req.payload
}

func (req *RequestForReplaceVertex) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForReplaceVertex() *ResponseForReplaceVertex {
	return new(ResponseForReplaceVertex)
}

type ResponseForReplaceVertex struct {
	// @FIXME
	Field string
}

func (r *ResponseForReplaceVertex) IsResponse() {}

func (r ResponseForReplaceVertex) String() string {
	// @FIXME
	return fmt.Sprintf("Bootfile: %v", r.Field)
}
