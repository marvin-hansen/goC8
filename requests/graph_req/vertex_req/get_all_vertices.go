package vertex_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForGetAllVertices(fabric, graphName string) *RequestForGetAllVertices {
	return &RequestForGetAllVertices{
		path: fmt.Sprintf("_fabric/%v/_api/graph/%v/vertex", fabric, graphName),
	}
}

type RequestForGetAllVertices struct {
	path string
}

func (req *RequestForGetAllVertices) Path() string {
	return req.path
}

func (req *RequestForGetAllVertices) Method() string {
	return http.MethodGet
}

func (req *RequestForGetAllVertices) Query() string {
	return ""
}

func (req *RequestForGetAllVertices) HasQueryParameter() bool {
	return false
}

func (req *RequestForGetAllVertices) GetQueryParameter() string {
	return ""
}

func (req *RequestForGetAllVertices) Payload() []byte {
	return nil
}

func (req *RequestForGetAllVertices) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForGetAllVertices() *ResponseForGetAllVertices {
	return new(ResponseForGetAllVertices)
}

type ResponseForGetAllVertices struct {
	Code        int      `json:"code"`
	Error       bool     `json:"error"`
	Collections []string `json:"collections"`
}

func (r *ResponseForGetAllVertices) IsResponse() {}

func (r ResponseForGetAllVertices) String() string {
	return fmt.Sprintf(" Code: %v \n Error: %v \n Vertex Collections: %v",
		r.Code,
		r.Error,
		r.Collections,
	)
}
