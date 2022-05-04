package kv_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForGetAllValues(fabric, collectionName string, keys KeyCollection, offset, limit int) *RequestForGetAllValues {
	return &RequestForGetAllValues{
		path:       fmt.Sprintf("_fabric/%v/_api/kv/%v/values", fabric, collectionName),
		parameters: fmt.Sprintf("?offset=%v&limit=%v", offset, limit),
		payload:    keys.Json(),
	}
}

type RequestForGetAllValues struct {
	path       string
	parameters string
	payload    []byte
}

func (req *RequestForGetAllValues) Path() string {
	return req.path
}

func (req *RequestForGetAllValues) Method() string {
	return http.MethodPost
}

func (req *RequestForGetAllValues) Query() string {
	return ""
}

func (req *RequestForGetAllValues) HasQueryParameter() bool {
	return true
}

func (req *RequestForGetAllValues) GetQueryParameter() string {
	return req.parameters
}

func (req *RequestForGetAllValues) Payload() []byte {
	return req.payload
}

func (req *RequestForGetAllValues) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForGetAllValues() *ResponseForGetAllValues {
	return new(ResponseForGetAllValues)
}

type ResponseForGetAllValues struct {
	Error  bool `json:"error"`
	Code   int  `json:"code"`
	Result []struct {
		Key      string `json:"_key"`
		Value    string `json:"value"`
		ExpireAt int    `json:"expireAt"`
	} `json:"result"`
}

func (r *ResponseForGetAllValues) IsResponse() {}

func (r ResponseForGetAllValues) String() string {
	return fmt.Sprintf(" Code: %v\n Error: %v\n Result: %v\n  ",
		r.Code,
		r.Error,
		r.Result,
	)
}
