package graph_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForGetGraph(fabric, graphName string) *RequestForGetGraph {
	return &RequestForGetGraph{
		path: fmt.Sprintf("_fabric/%v/_api/graph/%v", fabric, graphName),
	}
}

type RequestForGetGraph struct {
	path string
}

func (req *RequestForGetGraph) Path() string {
	return req.path
}

func (req *RequestForGetGraph) Method() string {
	return http.MethodGet
}

func (req *RequestForGetGraph) Query() string {
	return ""
}

func (req *RequestForGetGraph) HasQueryParameter() bool {
	return false
}

func (req *RequestForGetGraph) GetQueryParameter() string {
	return ""
}

func (req *RequestForGetGraph) Payload() []byte {
	return nil
}

func (req *RequestForGetGraph) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForGetGraph() *ResponseForGetGraph {
	return new(ResponseForGetGraph)
}

type ResponseForGetGraph struct {
	Code   int   `json:"code,omitempty"`
	Error  bool  `json:"error,omitempty"`
	Graphs Graph `json:"graph,omitempty"`
}

func (r *ResponseForGetGraph) IsResponse() {}

func (r ResponseForGetGraph) String() string {
	return fmt.Sprintf(" Code: %v \n Error: %v \n Graphs: %v",
		r.Code,
		r.Error,
		r.Graphs,
	)
}
