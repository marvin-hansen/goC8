package query_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForExplainQuery(fabric, queryString string) *RequestForExplainQuery {
	return &RequestForExplainQuery{
		path:    fmt.Sprintf("_fabric/%v/_api/query/explain", fabric),
		payload: getQueryPayload(queryString),
	}
}

type RequestForExplainQuery struct {
	path    string
	payload []byte
}

func (req *RequestForExplainQuery) Path() string {
	return req.path
}

func (req *RequestForExplainQuery) Method() string {
	return http.MethodPost
}

func (req *RequestForExplainQuery) Query() string {
	return ""
}

func (req *RequestForExplainQuery) HasQueryParameter() bool {
	return false
}

func (req *RequestForExplainQuery) GetQueryParameter() string {
	return ""
}

func (req *RequestForExplainQuery) Payload() []byte {
	return req.payload
}

func (req *RequestForExplainQuery) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForExplainQuery() *ResponseForExplainQuery {
	return new(ResponseForExplainQuery)
}

type ResponseForExplainQuery struct {
	Code      int          `json:"code,omitempty"`
	Error     bool         `json:"error,omitempty"`
	Plan      Plan         `json:"plan,omitempty"`
	Cacheable bool         `json:"cacheable,omitempty"`
	Warnings  Warnings     `json:"warnings,omitempty"`
	Stats     ExplainStats `json:"stats,omitempty"`
}

func (r *ResponseForExplainQuery) IsResponse() {}

func (r ResponseForExplainQuery) String() string {
	return fmt.Sprintf("Code: %v\n Error: %v\n Plan: %v\n Cacheable: %v\n Warnings %v\n Stats: %v\n",
		r.Code,
		r.Error,
		r.Plan,
		r.Cacheable,
		r.Warnings,
		r.Stats,
	)
}
