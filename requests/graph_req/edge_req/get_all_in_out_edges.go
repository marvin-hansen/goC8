package edge_req

import (
	"fmt"
	"github.com/marvin-hansen/goC8/requests/graph_req"
	"github.com/marvin-hansen/goC8/types"
	"net/http"
)

//**// Request //**//

func NewRequestForGetAllInOutEdges(fabric, edgeCollectionName, vertexKey string, direction types.EdgeDirection) *RequestForGetAllInOutEdges {
	return &RequestForGetAllInOutEdges{
		path:       fmt.Sprintf("_fabric/%v/_api/edges/%v", fabric, edgeCollectionName),
		parameters: getParameters(vertexKey, direction),
	}
}

func getParameters(vertexKey string, direction types.EdgeDirection) string {
	if direction != types.ANY {
		return fmt.Sprintf("?vertex=%v&direction=%v", vertexKey, direction.String())

	} else {
		return fmt.Sprintf("?vertex=%v", vertexKey)
	}
}

type RequestForGetAllInOutEdges struct {
	path       string
	parameters string
}

func (req *RequestForGetAllInOutEdges) Path() string {
	return req.path
}

func (req *RequestForGetAllInOutEdges) Method() string {
	return http.MethodGet
}

func (req *RequestForGetAllInOutEdges) Query() string {
	return ""
}

func (req *RequestForGetAllInOutEdges) HasQueryParameter() bool {
	return true
}

func (req *RequestForGetAllInOutEdges) GetQueryParameter() string {
	return req.parameters
}

func (req *RequestForGetAllInOutEdges) Payload() []byte {
	return nil
}

func (req *RequestForGetAllInOutEdges) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForGetAllInOutEdges() *ResponseForGetAllInOutEdges {
	return new(ResponseForGetAllInOutEdges)
}

type ResponseForGetAllInOutEdges struct {
	Code  int              `json:"code,omitempty"`
	Error bool             `json:"error,omitempty"`
	Edges []graph_req.Edge `json:"edges,omitempty"`
}

func (r *ResponseForGetAllInOutEdges) IsResponse() {}

func (r ResponseForGetAllInOutEdges) String() string {
	return fmt.Sprintf(" Code: %v \n Error: %v \n Edges: %v",
		r.Code,
		r.Error,
		r.Edges,
	)
}
