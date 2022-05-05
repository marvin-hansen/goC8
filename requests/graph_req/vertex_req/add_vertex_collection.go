package vertex_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForAddVertexCollection(fabric, graphName, vertexCollectionName string) *RequestForAddVertexCollection {
	return &RequestForAddVertexCollection{
		path:    fmt.Sprintf("_fabric/%v/_api/graph/%v/vertex", fabric, graphName),
		payload: getVertexPayload(vertexCollectionName),
	}
}

func getVertexPayload(vertexCollectionName string) []byte {
	str := fmt.Sprintf(`{'collection': %v}`, vertexCollectionName)
	return []byte(str)
}

type RequestForAddVertexCollection struct {
	path    string
	payload []byte
}

func (req *RequestForAddVertexCollection) Path() string {
	return req.path
}

func (req *RequestForAddVertexCollection) Method() string {
	return http.MethodPost
}

func (req *RequestForAddVertexCollection) Query() string {
	return ""
}

func (req *RequestForAddVertexCollection) HasQueryParameter() bool {
	return false
}

func (req *RequestForAddVertexCollection) GetQueryParameter() string {
	return ""
}

func (req *RequestForAddVertexCollection) Payload() []byte {
	return req.payload
}

func (req *RequestForAddVertexCollection) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForAddVertexCollection() *ResponseForAddVertexCollection {
	return new(ResponseForAddVertexCollection)
}

type ResponseForAddVertexCollection struct {
	// @FIXME
	Field string
}

func (r *ResponseForAddVertexCollection) IsResponse() {}

func (r ResponseForAddVertexCollection) String() string {
	// @FIXME
	return fmt.Sprintf("Bootfile: %v", r.Field)
}
