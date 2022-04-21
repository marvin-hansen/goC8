package graph_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForDeleteGraph(fabric, graphName string, dropCollections bool) *RequestForDeleteGraph {
	return &RequestForDeleteGraph{
		path:       fmt.Sprintf("_fabric/%v/_api/graph/%v", fabric, graphName),
		parameters: fmt.Sprintf("?dropCollections=%v", dropCollections),
	}
}

type RequestForDeleteGraph struct {
	path       string
	parameters string
}

func (req *RequestForDeleteGraph) Path() string {
	return req.path
}

func (req *RequestForDeleteGraph) Method() string {
	return http.MethodDelete
}

func (req *RequestForDeleteGraph) Query() string {
	return ""
}

func (req *RequestForDeleteGraph) HasQueryParameter() bool {
	return true
}

func (req *RequestForDeleteGraph) GetQueryParameter() string {
	return req.parameters
}

func (req *RequestForDeleteGraph) Payload() []byte {
	return nil
}

func (req *RequestForDeleteGraph) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForDeleteGraph() *ResponseForDeleteGraph {
	return new(ResponseForDeleteGraph)
}

type ResponseForDeleteGraph struct {
	Code    int  `json:"code"`
	Error   bool `json:"error"`
	Removed bool `json:"removed"`
}

func (r *ResponseForDeleteGraph) IsResponse() {}

func (r ResponseForDeleteGraph) String() string {
	return fmt.Sprintf("Code: %v, Error: %v, Removed: %v,",
		r.Code,
		r.Error,
		r.Removed,
	)
}
