package graph_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForReplaceEdge(fabric, graphName, edgeCollectionName, sourceVertex, destinationVertex string, dropCollections bool) *RequestForReplaceEdge {
	return &RequestForReplaceEdge{
		path:       fmt.Sprintf("_fabric/%v/_api/graph/%v/edge", fabric, graphName),
		parameters: fmt.Sprintf("?dropCollections=%v", dropCollections),
		payload:    getEdgePayload(edgeCollectionName, sourceVertex, destinationVertex),
	}
}

type RequestForReplaceEdge struct {
	path       string
	parameters string
	payload    []byte
}

func (req *RequestForReplaceEdge) Path() string {
	return req.path
}

func (req *RequestForReplaceEdge) Method() string {
	return http.MethodPut
}

func (req *RequestForReplaceEdge) Query() string {
	return ""
}

func (req *RequestForReplaceEdge) HasQueryParameter() bool {
	return true
}

func (req *RequestForReplaceEdge) GetQueryParameter() string {
	return req.parameters
}

func (req *RequestForReplaceEdge) Payload() []byte {
	return req.payload
}

func (req *RequestForReplaceEdge) ResponseCode() int {
	return 201 // ok
}

//**// Response //**//

func NewResponseForReplaceEdge() *ResponseForReplaceEdge {
	return new(ResponseForReplaceEdge)
}

type ResponseForReplaceEdge struct {
	Code   int   `json:"code,omitempty"`
	Error  bool  `json:"error,omitempty"`
	Graphs Graph `json:"graph,omitempty"`
}

func (r *ResponseForReplaceEdge) IsResponse() {}

func (r ResponseForReplaceEdge) String() string {
	return fmt.Sprintf(" Code: %v \n Error: %v \n Graphs: %v",
		r.Code,
		r.Error,
		r.Graphs,
	)
}