package kv_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForGetAllValues(fabric string) *RequestForGetAllValues {
  // @FIXME: Add correct API path
	return &RequestForGetAllValues{
			path: fmt.Sprintf("_fabric/%v/_api/NAME", fabric),
	}
}

type RequestForGetAllValues struct {
	path string
}

func (req *RequestForGetAllValues) Path() string {
	return req.path
}

func (req *RequestForGetAllValues) Method() string {
	return http.MethodGet
}

func (req *RequestForGetAllValues) Query() string {
	return ""
}

func (req *RequestForGetAllValues) HasQueryParameter() bool {
	return false
}

func (req *RequestForGetAllValues) GetQueryParameter() string {
	return ""
}

func (req *RequestForGetAllValues) Payload() []byte {
	return nil
}

func (req *RequestForGetAllValues) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForGetAllValues() *ResponseForGetAllValues {
	return new(ResponseForGetAllValues)
}

type ResponseForGetAllValues struct {
  // @FIXME
	Field string 
}

func (r *ResponseForGetAllValues) IsResponse() {}

func (r ResponseForGetAllValues) String() string {
  // @FIXME
	return fmt.Sprintf("Bootfile: %v", r.Field)
}


