package graph_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForGetAllEdges(fabric, graphName string) *RequestForGetAllEdges {
	return &RequestForGetAllEdges{
		path: fmt.Sprintf("_fabric/%v/_api/graph/%v/edge", fabric, graphName),
	}
}

type RequestForGetAllEdges struct {
	path string
}

func (req *RequestForGetAllEdges) Path() string {
	return req.path
}

func (req *RequestForGetAllEdges) Method() string {
	return http.MethodGet
}

func (req *RequestForGetAllEdges) Query() string {
	return ""
}

func (req *RequestForGetAllEdges) HasQueryParameter() bool {
	return false
}

func (req *RequestForGetAllEdges) GetQueryParameter() string {
	return ""
}

func (req *RequestForGetAllEdges) Payload() []byte {
	return nil
}

func (req *RequestForGetAllEdges) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForGetAllEdges() *ResponseForGetAllEdges {
	return new(ResponseForGetAllEdges)
}

type ResponseForGetAllEdges struct {
	Code        int      `json:"code"`
	Error       bool     `json:"error"`
	Collections []string `json:"collections"`
}

func (r *ResponseForGetAllEdges) IsResponse() {}

func (r ResponseForGetAllEdges) String() string {
	return fmt.Sprintf(" Code: %v \n Error: %v \n Edge Collections: %v",
		r.Code,
		r.Error,
		r.Collections,
	)
}
