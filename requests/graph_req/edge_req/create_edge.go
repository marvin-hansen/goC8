package edge_req

import (
	"fmt"
	"github.com/marvin-hansen/goC8/requests/graph_req"
	"net/http"
)

//**// Request //**//

func NewRequestForCreateEdge(fabric, graphName, edgeCollectionName, sourceVertex, destinationVertex string, jsonDef []byte, returnNew bool) *RequestForCreateEdge {

	return &RequestForCreateEdge{
		path:       fmt.Sprintf("_fabric/%v/_api/graph/%v/edge/%v", fabric, graphName, edgeCollectionName),
		parameters: fmt.Sprintf("?returnNew=%v", returnNew),
		payload:    getCreateEdgePayload(sourceVertex, destinationVertex, jsonDef),
	}
}

func getCreateEdgePayload(sourceVertex, destinationVertex string, jsonDef []byte) []byte {
	if jsonDef == nil {
		str := fmt.Sprintf(`{
  "_from": "%v",
  "_to": "%v"
}`, sourceVertex, destinationVertex)
		return []byte(str)
	} else {
		return jsonDef
	}
}

type RequestForCreateEdge struct {
	path       string
	parameters string
	payload    []byte
}

func (req *RequestForCreateEdge) Path() string {
	return req.path
}

func (req *RequestForCreateEdge) Method() string {
	return http.MethodPost
}

func (req *RequestForCreateEdge) Query() string {
	return ""
}

func (req *RequestForCreateEdge) HasQueryParameter() bool {
	return true
}

func (req *RequestForCreateEdge) GetQueryParameter() string {
	return req.parameters
}

func (req *RequestForCreateEdge) Payload() []byte {
	return req.payload
}

func (req *RequestForCreateEdge) ResponseCode() int {
	return 201 // ok
}

//**// Response //**//

func NewResponseForCreateEdge() *ResponseForCreateEdge {
	return new(ResponseForCreateEdge)
}

type ResponseForCreateEdge struct {
	Code  int            `json:"code,omitempty"`
	Error bool           `json:"error,omitempty"`
	Edge  graph_req.Edge `json:"edge,omitempty"`
	New   graph_req.Edge `json:"new,omitempty"`
}

func (r *ResponseForCreateEdge) IsResponse() {}

func (r ResponseForCreateEdge) String() string {
	return fmt.Sprintf("Code: %v\n, Error: %v\n Edge: %v\n New: %v\n",
		r.Code,
		r.Error,
		r.Edge,
		r.New,
	)
}
