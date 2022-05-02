package kv_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForGetAllKeys(fabric, collectionName string, offset, limit int, order Order) *RequestForGetAllKeys {
	return &RequestForGetAllKeys{
		path:       fmt.Sprintf("_fabric/%v/_api/kv/%v/keys", fabric, collectionName),
		parameters: fmt.Sprintf("?offset=%v&limit=%v&order=%v", offset, limit, order.String()),
	}
}

type RequestForGetAllKeys struct {
	path       string
	parameters string
	payload    []byte
}

func (req *RequestForGetAllKeys) Path() string {
	return req.path
}

func (req *RequestForGetAllKeys) Method() string {
	return http.MethodGet
}

func (req *RequestForGetAllKeys) Query() string {
	return ""
}

func (req *RequestForGetAllKeys) HasQueryParameter() bool {
	return true
}

func (req *RequestForGetAllKeys) GetQueryParameter() string {
	return ""
}

func (req *RequestForGetAllKeys) Payload() []byte {
	return nil
}

func (req *RequestForGetAllKeys) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForGetAllKeys() *ResponseForGetAllKeys {
	return new(ResponseForGetAllKeys)
}

type ResponseForGetAllKeys struct {
	Error  bool     `json:"error,omitempty"`
	Code   int      `json:"code,omitempty"`
	Result []string `json:"result,omitempty"`
}

func (r *ResponseForGetAllKeys) IsResponse() {}

func (r ResponseForGetAllKeys) String() string {
	return fmt.Sprintf(" Error: %v\n Code: %v\n Result: %v\n",
		r.Error,
		r.Code,
		r.Result,
	)
}
