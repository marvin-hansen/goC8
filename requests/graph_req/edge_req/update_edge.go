package edge_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForUpdateEdge(fabric, graphName, collectionName, edgeKey string, jsonUpdate []byte, keepNull, returnOld, returnNew bool) *RequestForUpdateEdge {
	return &RequestForUpdateEdge{
		path:       fmt.Sprintf("_fabric/%v/_api/graph/%v/edge/%v/%v", fabric, graphName, collectionName, edgeKey),
		parameters: fmt.Sprintf("?keepNull=%v&returnOld=%v&returnNew=%v", keepNull, returnOld, returnNew),
		payload:    jsonUpdate,
	}
}

type RequestForUpdateEdge struct {
	path       string
	parameters string
	payload    []byte
}

func (req *RequestForUpdateEdge) Path() string {
	return req.path
}

func (req *RequestForUpdateEdge) Method() string {
	return http.MethodPatch
}

func (req *RequestForUpdateEdge) Query() string {
	return ""
}

func (req *RequestForUpdateEdge) HasQueryParameter() bool {
	return true
}

func (req *RequestForUpdateEdge) GetQueryParameter() string {
	return req.parameters
}

func (req *RequestForUpdateEdge) Payload() []byte {
	return req.payload
}

func (req *RequestForUpdateEdge) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

// Call NewResponseForEdge()
