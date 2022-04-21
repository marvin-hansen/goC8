package index_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForGetAllIndices(fabric, collectionName string) *RequestForGetAllIndices {
	return &RequestForGetAllIndices{
		path:       fmt.Sprintf("_fabric/%v/_api/index", fabric),
		parameters: fmt.Sprintf("?collection=%v", collectionName),
	}
}

type RequestForGetAllIndices struct {
	path       string
	parameters string
}

func (req *RequestForGetAllIndices) Path() string {
	return req.path
}

func (req *RequestForGetAllIndices) Method() string {
	return http.MethodGet
}

func (req *RequestForGetAllIndices) Query() string {
	return ""
}

func (req *RequestForGetAllIndices) HasQueryParameter() bool {
	return true
}

func (req *RequestForGetAllIndices) GetQueryParameter() string {
	return req.parameters
}

func (req *RequestForGetAllIndices) Payload() []byte {
	return nil
}

func (req *RequestForGetAllIndices) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForGetAllIndices() *ResponseForGetAllIndices {
	return new(ResponseForGetAllIndices)
}

type ResponseForGetAllIndices struct {
	Error   bool    `json:"error,omitempty"`
	Code    int     `json:"code,omitempty"`
	Indexes []Index `json:"indexes,omitempty"`
	//Identifiers Identifiers `json:"identifiers,omitempty"`
}

func (r *ResponseForGetAllIndices) IsResponse() {}

func (r ResponseForGetAllIndices) String() string {
	return fmt.Sprintf(" Error: %v\n Code: %v\n Indexes: %v\n",
		r.Error,
		r.Code,
		r.Indexes,
	)
}
