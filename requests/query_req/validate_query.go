package query_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForValidateQuery(fabric, queryString string) *RequestForValidateQuery {
	return &RequestForValidateQuery{
		path:    fmt.Sprintf("_fabric/%v/_api/query", fabric),
		payload: getValidatePayload(queryString),
	}
}

func getValidatePayload(query string) []byte {
	return []byte(fmt.Sprintf(`{"query": "%v"}`, query))
}

type RequestForValidateQuery struct {
	path    string
	payload []byte
}

func (req *RequestForValidateQuery) Path() string {
	return req.path
}

func (req *RequestForValidateQuery) Method() string {
	return http.MethodPost
}

func (req *RequestForValidateQuery) Query() string {
	return ""
}

func (req *RequestForValidateQuery) HasQueryParameter() bool {
	return false
}

func (req *RequestForValidateQuery) GetQueryParameter() string {
	return "" //"?excludeSystem=true"
}

func (req *RequestForValidateQuery) Payload() []byte {
	return req.payload
}

func (req *RequestForValidateQuery) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForValidateQuery() *ResponseForValidateQuery {
	return new(ResponseForValidateQuery)
}

type ResponseForValidateQuery struct {
	Error       bool          `json:"error"`
	Code        int           `json:"code"`
	Parsed      bool          `json:"parsed"`
	Collections []string      `json:"collections"`
	BindVars    []interface{} `json:"bindVars"`
}

func (r *ResponseForValidateQuery) IsResponse() {}

func (r ResponseForValidateQuery) String() string {
	return fmt.Sprintf("Error: %v\n Code: %v\n Parsed: %v\n Collections: %v\n BindVars: %v\n",
		r.Error,
		r.Code,
		r.Parsed,
		r.Collections,
		r.BindVars,
	)
}
