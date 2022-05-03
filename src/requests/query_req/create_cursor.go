package query_req

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForCreateCursor(fabric, query string, bindVars map[string]interface{}, options *CursorOptions, ttl int) *RequestForCreateCursor {
	return &RequestForCreateCursor{
		path:    fmt.Sprintf("_fabric/%v/_api/cursor", fabric),
		payload: constructQuery(query, bindVars, options, ttl),
	}
}

type RequestForCreateCursor struct {
	path    string
	payload []byte
}

func constructQuery(query string, bindVars map[string]interface{}, options *CursorOptions, ttl int) []byte {

	q := Query{
		BatchSize: 100,
		BindVars:  bindVars,
		Count:     false,
		Options:   *options,
		Query:     query,
		TTL:       ttl,
	}

	b, err := json.Marshal(q)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return b

}

type Query struct {
	BatchSize int                    `json:"batchSize,omitempty"`
	BindVars  map[string]interface{} `json:"bindVars,omitempty"`
	Count     bool                   `json:"count,omitempty"`
	Options   CursorOptions          `json:"options,omitempty"`
	Query     string                 `json:"query,omitempty"`
	TTL       int                    `json:"ttl,omitempty"`
}

func (req *RequestForCreateCursor) Path() string {
	return req.path
}

func (req *RequestForCreateCursor) Method() string {
	return http.MethodPost
}

func (req *RequestForCreateCursor) Query() string {
	return ""
}

func (req *RequestForCreateCursor) HasQueryParameter() bool {
	return false
}

func (req *RequestForCreateCursor) GetQueryParameter() string {
	return ""
}

func (req *RequestForCreateCursor) Payload() []byte {
	return req.payload
}

func (req *RequestForCreateCursor) ResponseCode() int {
	return 201 // ok
}

//**// Response //**//

func NewResponseForCreateCursor() *ResponseForCreateCursor {
	return new(ResponseForCreateCursor)
}

// ResponseForCreateCursor see query_req/shared/Cursor for spec
type ResponseForCreateCursor Cursor

func (r *ResponseForCreateCursor) IsResponse() {}

func (r ResponseForCreateCursor) String() string {
	return fmt.Sprintf("Code: %v\n Error: %v\n Count: %v\n Extra: %v\n HasMore: %v\n Id: %v\n Cached: %v\n Result: %v\n",
		r.Code,
		r.Error,
		r.Count,
		r.Extra,
		r.HasMore,
		r.Id,
		r.Cached,
		r.Result,
	)
}
