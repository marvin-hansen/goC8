package graph_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForCreateEdge(fabric, graphName, edgeCollectionName, sourceVertex, destinationVertex string, returnNew bool) *RequestForCreateEdge {
	return &RequestForCreateEdge{
		path:       fmt.Sprintf("_fabric/%v/_api/graph/%v/edge/%v", fabric, graphName, edgeCollectionName),
		parameters: fmt.Sprintf("?returnNew=%v", returnNew),
		payload:    getCreateEdgePayload(sourceVertex, destinationVertex),
	}
}

func getCreateEdgePayload(sourceVertex, destinationVertex string) []byte {
	str := fmt.Sprintf(`{
  "_from": "%v",
  "_to": "%v"
}`, sourceVertex, destinationVertex)
	return []byte(str)
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
	Code  int  `json:"code,omitempty"`
	Error bool `json:"error,omitempty"`
	Edge  Edge `json:"edge,omitempty"`
	New   Edge `json:"new,omitempty"`
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
