package vertex_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForUpdateVertex(fabric, graphName, collectionName, vertexKey string, jsonUpdate []byte, keepNull, returnOld, returnNew bool) *RequestForUpdateVertex {
	return &RequestForUpdateVertex{
		path:       fmt.Sprintf("_fabric/%v/_api/graph/%v/vertex/%v/%v", fabric, graphName, collectionName, vertexKey),
		parameters: fmt.Sprintf("?keepNull=%v&returnOld=%v&returnNew=%v", keepNull, returnOld, returnNew),
		payload:    jsonUpdate,
	}
}

type RequestForUpdateVertex struct {
	path       string
	parameters string
	payload    []byte
}

func (req *RequestForUpdateVertex) Path() string {
	return req.path
}

func (req *RequestForUpdateVertex) Method() string {
	return http.MethodPatch
}

func (req *RequestForUpdateVertex) Query() string {
	return ""
}

func (req *RequestForUpdateVertex) HasQueryParameter() bool {
	return true
}

func (req *RequestForUpdateVertex) GetQueryParameter() string {
	return req.parameters
}

func (req *RequestForUpdateVertex) Payload() []byte {
	return req.payload
}

func (req *RequestForUpdateVertex) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

// call NewResponseForVertex()
