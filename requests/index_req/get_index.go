package index_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForGetIndex(fabric, collectionName, indexName string) *RequestForGetIndex {
	return &RequestForGetIndex{
		path: fmt.Sprintf("_fabric/%v/_api/index/%v/%v", fabric, collectionName, indexName),
	}
}

type RequestForGetIndex struct {
	path string
}

func (req *RequestForGetIndex) Path() string {
	return req.path
}

func (req *RequestForGetIndex) Method() string {
	return http.MethodGet
}

func (req *RequestForGetIndex) Query() string {
	return ""
}

func (req *RequestForGetIndex) HasQueryParameter() bool {
	return false
}

func (req *RequestForGetIndex) GetQueryParameter() string {
	return ""
}

func (req *RequestForGetIndex) Payload() []byte {
	return nil
}

func (req *RequestForGetIndex) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForGetIndex() *ResponseForGetIndex {
	return new(ResponseForGetIndex)
}

type ResponseForGetIndex IndexEntry

func (r *ResponseForGetIndex) IsResponse() {}

func (r ResponseForGetIndex) String() string {
	return fmt.Sprintf(" Error: %v\n Code: %v\n Fields: %v\n ID: %v\n Name: %v\n MinLength: %v\n SelectivityEstimate: %v\n Sparse: %v\n  Type: %v\n  Unique: %v\n",
		r.Error,
		r.Code,
		r.Fields,
		r.Id,
		r.Name,
		r.MinLength,
		r.SelectivityEstimate,
		r.Sparse,
		r.Type,
		r.Unique,
	)
}
