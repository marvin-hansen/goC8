package vertex_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForReplaceVertex(fabric, graphName, collectionName, vertexKey string, jsonReplace []byte, keepNull, returnOld, returnNew bool) *RequestForReplaceVertex {
	return &RequestForReplaceVertex{
		path:       fmt.Sprintf("_fabric/%v/_api/graph/%v/vertex/%v/%v", fabric, graphName, collectionName, vertexKey),
		parameters: fmt.Sprintf("?keepNull=%v&returnOld=%v&returnNew=%v", keepNull, returnOld, returnNew),
		payload:    jsonReplace,
	}
}

type RequestForReplaceVertex struct {
	path       string
	parameters string
	payload    []byte
}

func (req *RequestForReplaceVertex) Path() string {
	return req.path
}

func (req *RequestForReplaceVertex) Method() string {
	return http.MethodPut
}

func (req *RequestForReplaceVertex) Query() string {
	return ""
}

func (req *RequestForReplaceVertex) HasQueryParameter() bool {
	return true
}

func (req *RequestForReplaceVertex) GetQueryParameter() string {
	return req.parameters
}

func (req *RequestForReplaceVertex) Payload() []byte {
	return req.payload
}

func (req *RequestForReplaceVertex) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

// call NewResponseForVertex()
