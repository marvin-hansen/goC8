package vertex_req

import (
	"fmt"
	"github.com/marvin-hansen/goC8/requests/graph_req/shared"
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

func NewResponseForGetVertex() *ResponseForGetVertex {
	return new(ResponseForGetVertex)
}

type ResponseForGetVertex struct {
	Code   int           `json:"code"`
	Error  bool          `json:"error"`
	Vertex shared.Vertex `json:"vertex"`
}

func (r *ResponseForGetVertex) IsResponse() {}

func (r ResponseForGetVertex) String() string {
	return fmt.Sprintf(" Code: %v \n Error: %v \n Vertex: %v",
		r.Code,
		r.Error,
		r.Vertex,
	)
}
