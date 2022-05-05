package vertex_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForDeleteVertex(fabric, graphName, collectionName, vertexKey string, returnOld bool) *RequestForDeleteVertex {
	return &RequestForDeleteVertex{
		path:       fmt.Sprintf("_fabric/%v/_api/graph/%v/vertex/%v/%v", fabric, graphName, collectionName, vertexKey),
		parameters: fmt.Sprintf("?returnOld=%v", returnOld),
	}
}

type RequestForDeleteVertex struct {
	path       string
	parameters string
}

func (req *RequestForDeleteVertex) Path() string {
	return req.path
}

func (req *RequestForDeleteVertex) Method() string {
	return http.MethodDelete
}

func (req *RequestForDeleteVertex) Query() string {
	return ""
}

func (req *RequestForDeleteVertex) HasQueryParameter() bool {
	return true
}

func (req *RequestForDeleteVertex) GetQueryParameter() string {
	return req.parameters
}

func (req *RequestForDeleteVertex) Payload() []byte {
	return nil
}

func (req *RequestForDeleteVertex) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

// call NewResponseForVertex()
