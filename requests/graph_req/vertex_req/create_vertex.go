package vertex_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForCreateVertex(fabric, graphName, vertexCollectionName string, jsonDef []byte, returnNew bool) *RequestForCreateVertex {
	return &RequestForCreateVertex{
		path:       fmt.Sprintf("_fabric/%v/_api/graph/%v/vertex/%v", fabric, graphName, vertexCollectionName),
		parameters: fmt.Sprintf("?returnNew=%v", returnNew),
		payload:    jsonDef,
	}
}

type RequestForCreateVertex struct {
	path       string
	parameters string
	payload    []byte
}

func (req *RequestForCreateVertex) Path() string {
	return req.path
}

func (req *RequestForCreateVertex) Method() string {
	return http.MethodPost
}

func (req *RequestForCreateVertex) Query() string {
	return ""
}

func (req *RequestForCreateVertex) HasQueryParameter() bool {
	return true
}

func (req *RequestForCreateVertex) GetQueryParameter() string {
	return req.parameters
}

func (req *RequestForCreateVertex) Payload() []byte {
	return req.payload
}

func (req *RequestForCreateVertex) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

// call NewResponseForVertex()
