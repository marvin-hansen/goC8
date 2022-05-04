package graph_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForDeleteEdge(fabric, graphName, collectionName, edgeKey string, returnOld bool) *RequestForDeleteEdge {
	return &RequestForDeleteEdge{
		path:       fmt.Sprintf("_fabric/%v/_api/graph/%v/edge/%v/%v", fabric, graphName, collectionName, edgeKey),
		parameters: fmt.Sprintf("?returnOld=%v", returnOld),
	}
}

type RequestForDeleteEdge struct {
	path       string
	parameters string
}

func (req *RequestForDeleteEdge) Path() string {
	return req.path
}

func (req *RequestForDeleteEdge) Method() string {
	return http.MethodDelete
}

func (req *RequestForDeleteEdge) Query() string {
	return ""
}

func (req *RequestForDeleteEdge) HasQueryParameter() bool {
	return true
}

func (req *RequestForDeleteEdge) GetQueryParameter() string {
	return req.parameters
}

func (req *RequestForDeleteEdge) Payload() []byte {
	return nil
}

func (req *RequestForDeleteEdge) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

// call NewEdgeResponse() instead
