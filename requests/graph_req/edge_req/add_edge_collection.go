package edge_req

import (
	"fmt"
	"github.com/marvin-hansen/goC8/requests/graph_req"
	"net/http"
)

//**// Request //**//

func NewRequestForAddEdgeCollection(fabric, graphName, edgeCollectionName, sourceVertex, destinationVertex string) *RequestForAddEdgeCollection {
	return &RequestForAddEdgeCollection{
		path:    fmt.Sprintf("_fabric/%v/_api/graph/%v/edge", fabric, graphName),
		payload: getEdgePayload(edgeCollectionName, sourceVertex, destinationVertex),
	}
}

func getEdgePayload(collection, sourceVertex, destinationVertex string) []byte {
	str := fmt.Sprintf(`{
  "collection": "%v",
  "from": [
    "%v"
  ],
  "to": [
    "%v"
  ]
}`, collection, sourceVertex, destinationVertex)
	return []byte(str)
}

type RequestForAddEdgeCollection struct {
	path    string
	payload []byte
}

func (req *RequestForAddEdgeCollection) Path() string {
	return req.path
}

func (req *RequestForAddEdgeCollection) Method() string {
	return http.MethodPost
}

func (req *RequestForAddEdgeCollection) Query() string {
	return ""
}

func (req *RequestForAddEdgeCollection) HasQueryParameter() bool {
	return false
}

func (req *RequestForAddEdgeCollection) GetQueryParameter() string {
	return ""
}

func (req *RequestForAddEdgeCollection) Payload() []byte {
	return req.payload
}

func (req *RequestForAddEdgeCollection) ResponseCode() int {
	return 201 // ok
}

//**// Response //**//

func NewResponseForAddEdgeCollection() *ResponseForAddEdgeCollection {
	return new(ResponseForAddEdgeCollection)
}

type ResponseForAddEdgeCollection struct {
	Code  int                `json:"code,omitempty"`
	Error bool               `json:"error,omitempty"`
	Graph graph_req.NewGraph `json:"graph,omitempty"`
}

func (r *ResponseForAddEdgeCollection) IsResponse() {}

func (r ResponseForAddEdgeCollection) String() string {
	return fmt.Sprintf(" Code: %v \n Error: %v \n Graph: %v",
		r.Code,
		r.Error,
		r.Graph,
	)
}
