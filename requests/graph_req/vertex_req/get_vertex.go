package vertex_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForGetVertex(fabric, graphName, collectionName, vertexKey string) *RequestForGetVertex {
	return &RequestForGetVertex{
		path: fmt.Sprintf("_fabric/%v/_api/graph/%v/vertex/%v/%v", fabric, graphName, collectionName, vertexKey),
	}
}

type RequestForGetVertex struct {
	path string
}

func (req *RequestForGetVertex) Path() string {
	return req.path
}

func (req *RequestForGetVertex) Method() string {
	return http.MethodGet
}

func (req *RequestForGetVertex) Query() string {
	return ""
}

func (req *RequestForGetVertex) HasQueryParameter() bool {
	return false
}

func (req *RequestForGetVertex) GetQueryParameter() string {
	return ""
}

func (req *RequestForGetVertex) Payload() []byte {
	return nil
}

func (req *RequestForGetVertex) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

// call NewResponseForVertex()
