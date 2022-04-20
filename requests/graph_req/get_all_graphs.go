package graph_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForGetAllGraphs(fabric string) *RequestForGetAllGraphs {
	return &RequestForGetAllGraphs{
		path: fmt.Sprintf("_fabric/%v/_api/graph", fabric),
	}
}

type RequestForGetAllGraphs struct {
	path string
}

func (req *RequestForGetAllGraphs) Path() string {
	return req.path
}

func (req *RequestForGetAllGraphs) Method() string {
	return http.MethodGet
}

func (req *RequestForGetAllGraphs) Query() string {
	return ""
}

func (req *RequestForGetAllGraphs) HasQueryParameter() bool {
	return false
}

func (req *RequestForGetAllGraphs) GetQueryParameter() string {
	return ""
}

func (req *RequestForGetAllGraphs) Payload() []byte {
	return nil
}

func (req *RequestForGetAllGraphs) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForGetAllGraphs() *ResponseForGetAllGraphs {
	return new(ResponseForGetAllGraphs)
}

type ResponseForGetAllGraphs struct {
	// @FIXME
	Field string
}

func (r *ResponseForGetAllGraphs) IsResponse() {}

func (r ResponseForGetAllGraphs) String() string {
	// @FIXME
	return fmt.Sprintf("Bootfile: %v", r.Field)
}
