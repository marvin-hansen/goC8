package vertex_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForDeleteVertexCollection(fabric, graphName, collectionName string, dropCollections bool) *RequestForDeleteVertexCollection {
	return &RequestForDeleteVertexCollection{
		path:       fmt.Sprintf("_fabric/%v/_api/graph/%v/vertex/%v", fabric, graphName, collectionName),
		parameters: fmt.Sprintf("?dropCollections=%v", dropCollections)}
}

type RequestForDeleteVertexCollection struct {
	path       string
	parameters string
}

func (req *RequestForDeleteVertexCollection) Path() string {
	return req.path
}

func (req *RequestForDeleteVertexCollection) Method() string {
	return http.MethodDelete
}

func (req *RequestForDeleteVertexCollection) Query() string {
	return ""
}

func (req *RequestForDeleteVertexCollection) HasQueryParameter() bool {
	return true
}

func (req *RequestForDeleteVertexCollection) GetQueryParameter() string {
	return req.parameters
}

func (req *RequestForDeleteVertexCollection) Payload() []byte {
	return nil
}

func (req *RequestForDeleteVertexCollection) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

// call  NewResponseForGraph()
