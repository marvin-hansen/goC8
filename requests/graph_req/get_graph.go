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

// call  NewResponseForGraph()
