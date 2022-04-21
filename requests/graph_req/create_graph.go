package graph_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForCreateGraph(fabric, graphName string) *RequestForCreateGraph {
	// @FIXME: Add correct API path
	return &RequestForCreateGraph{
		path: fmt.Sprintf("_fabric/%v/_api/NAME", fabric),
	}
}

type RequestForCreateGraph struct {
	path string
}

func (req *RequestForCreateGraph) Path() string {
	return req.path
}

func (req *RequestForCreateGraph) Method() string {
	return http.MethodGet
}

func (req *RequestForCreateGraph) Query() string {
	return ""
}

func (req *RequestForCreateGraph) HasQueryParameter() bool {
	return false
}

func (req *RequestForCreateGraph) GetQueryParameter() string {
	return ""
}

func (req *RequestForCreateGraph) Payload() []byte {
	return nil
}

func (req *RequestForCreateGraph) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForCreateGraph() *ResponseForCreateGraph {
	return new(ResponseForCreateGraph)
}

type ResponseForCreateGraph struct {
	// @FIXME
	Field string
}

func (r *ResponseForCreateGraph) IsResponse() {}

func (r ResponseForCreateGraph) String() string {
	// @FIXME
	return fmt.Sprintf("Bootfile: %v", r.Field)
}
