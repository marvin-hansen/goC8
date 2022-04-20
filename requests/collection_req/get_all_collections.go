package collection_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForGetAllCollections(fabric string) *RequestForGetAllCollections {
	return &RequestForGetAllCollections{
		path: fmt.Sprintf("_fabric/%v/_api/collection", fabric),
	}
}

type RequestForGetAllCollections struct {
	path string
}

func (req *RequestForGetAllCollections) Path() string {
	return req.path
}

func (req *RequestForGetAllCollections) Method() string {
	return http.MethodGet
}

func (req *RequestForGetAllCollections) Query() string {
	return ""
}

func (req *RequestForGetAllCollections) HasQueryParameter() bool {
	return true
}
func (req *RequestForGetAllCollections) GetQueryParameter() string {
	return "?excludeSystem=true"
}

func (req *RequestForGetAllCollections) Payload() []byte {
	return nil
}

func (req *RequestForGetAllCollections) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForGetAllCollections() *ResponseForGetAllCollections {
	return new(ResponseForGetAllCollections)
}

type ResponseForGetAllCollections struct {
	Result []ResulFromCollections `json:"result"`
}

func (r *ResponseForGetAllCollections) IsResponse() {}

func (r ResponseForGetAllCollections) String() string {
	return fmt.Sprintf("Result: %v", r.Result)
}
