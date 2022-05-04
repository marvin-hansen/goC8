package graph_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForDeleteEdgeCollection(fabric, graphName, collectionName string, dropCollections bool) *RequestForDeleteEdgeCollection {
	return &RequestForDeleteEdgeCollection{
		path:       fmt.Sprintf("_fabric/%v/_api/graph/%v/edge/%v", fabric, graphName, collectionName),
		parameters: fmt.Sprintf("?dropCollections=%v", dropCollections),
	}
}

type RequestForDeleteEdgeCollection struct {
	path       string
	parameters string
}

func (req *RequestForDeleteEdgeCollection) Path() string {
	return req.path
}

func (req *RequestForDeleteEdgeCollection) Method() string {
	return http.MethodDelete
}

func (req *RequestForDeleteEdgeCollection) Query() string {
	return ""
}

func (req *RequestForDeleteEdgeCollection) HasQueryParameter() bool {
	return true
}

func (req *RequestForDeleteEdgeCollection) GetQueryParameter() string {
	return req.parameters
}

func (req *RequestForDeleteEdgeCollection) Payload() []byte {
	return nil
}

func (req *RequestForDeleteEdgeCollection) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

// call  NewResponseForGraph()
