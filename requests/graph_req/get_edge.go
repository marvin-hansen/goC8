package graph_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForGetEdge(fabric, graphName, collectionName, edgeKey string) *RequestForGetEdge {
	return &RequestForGetEdge{
		path: fmt.Sprintf("_fabric/%v/_api/graph/%v/edge/%v/%v", fabric, graphName, collectionName, edgeKey),
	}
}

type RequestForGetEdge struct {
	path string
}

func (req *RequestForGetEdge) Path() string {
	return req.path
}

func (req *RequestForGetEdge) Method() string {
	return http.MethodGet
}

func (req *RequestForGetEdge) Query() string {
	return ""
}

func (req *RequestForGetEdge) HasQueryParameter() bool {
	return false
}

func (req *RequestForGetEdge) GetQueryParameter() string {
	return ""
}

func (req *RequestForGetEdge) Payload() []byte {
	return nil
}

func (req *RequestForGetEdge) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForGetEdge() *ResponseForGetEdge {
	return new(ResponseForGetEdge)
}

type ResponseForGetEdge struct {
	Code  int  `json:"code,omitempty"`
	Error bool `json:"error,omitempty"`
	Edge  Edge `json:"edge,omitempty"`
}

func (r *ResponseForGetEdge) IsResponse() {}

func (r ResponseForGetEdge) String() string {
	return fmt.Sprintf(" Code: %v \n Error: %v \n Edge: %v",
		r.Code,
		r.Error,
		r.Edge,
	)
}
