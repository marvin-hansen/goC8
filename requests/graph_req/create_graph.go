package graph_req

import (
	"fmt"
	"github.com/marvin-hansen/goC8/requests/graph_req/shared"
	"net/http"
)

//**// Request //**//

func NewRequestForCreateGraph(fabric string, jsonGraph []byte) *RequestForCreateGraph {
	return &RequestForCreateGraph{
		payload: jsonGraph,
		path:    fmt.Sprintf("_fabric/%v/_api/graph", fabric),
	}
}

type RequestForCreateGraph struct {
	path    string
	payload []byte
}

func (req *RequestForCreateGraph) Path() string {
	return req.path
}

func (req *RequestForCreateGraph) Method() string {
	return http.MethodPost
}

func (req *RequestForCreateGraph) Query() string {
	return ""
}

func (req *RequestForCreateGraph) HasQueryParameter() bool {
	return false
}

func (req *RequestForCreateGraph) GetQueryParameter() string {
	return ""
}

func (req *RequestForCreateGraph) Payload() []byte {
	return req.payload
}

func (req *RequestForCreateGraph) ResponseCode() int {
	return 201 // ok
}

//**// Response //**//

func NewResponseForCreateGraph() *ResponseForCreateGraph {
	return new(ResponseForCreateGraph)
}

type ResponseForCreateGraph struct {
	Code  int      `json:"code"`
	Error bool     `json:"error"`
	Graph NewGraph `json:"graph"`
}

func (r *ResponseForCreateGraph) IsResponse() {}

func (r ResponseForCreateGraph) String() string {
	return fmt.Sprintf(" Code: %v \n Error: %v \n Graph: %v",
		r.Code,
		r.Error,
		r.Graph,
	)
}

type NewGraph struct {
	Id                string                   `json:"_id"`
	Key               string                   `json:"_key"`
	Rev               string                   `json:"_rev"`
	Edgedefinitions   []shared.EdgeDefinitions `json:"edgedefinitions"`
	Name              string                   `json:"name"`
	OrphanCollections []string                 `json:"orphanCollections"`
}

func (r NewGraph) String() string {
	return fmt.Sprintf("ID: %v \n Key: %v \n Rev: %v \n Edgedefinitions: %v \n  Name: %v \n  OrphanCollections: %v \n ",
		r.Id,
		r.Key,
		r.Rev,
		r.Edgedefinitions,
		r.Name,
		r.OrphanCollections,
	)
}
