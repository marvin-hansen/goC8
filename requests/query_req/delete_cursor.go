package query_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForDeleteCursor(fabric, cursorID string) *RequestForDeleteCursor {
	return &RequestForDeleteCursor{
		path: fmt.Sprintf("_fabric/%v/_api/cursor/%v", fabric, cursorID),
	}
}

type RequestForDeleteCursor struct {
	path string
}

func (req *RequestForDeleteCursor) Path() string {
	return req.path
}

func (req *RequestForDeleteCursor) Method() string {
	return http.MethodDelete
}

func (req *RequestForDeleteCursor) Query() string {
	return ""
}

func (req *RequestForDeleteCursor) HasQueryParameter() bool {
	return false
}

func (req *RequestForDeleteCursor) GetQueryParameter() string {
	return ""
}

func (req *RequestForDeleteCursor) Payload() []byte {
	return nil
}

func (req *RequestForDeleteCursor) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForDeleteCursor() *ResponseForDeleteCursor {
	return new(ResponseForDeleteCursor)
}

type ResponseForDeleteCursor struct {
	Code  int    `json:"code"`
	Error bool   `json:"error"`
	Id    string `json:"id"`
}

func (r *ResponseForDeleteCursor) IsResponse() {}

func (r ResponseForDeleteCursor) String() string {
	return fmt.Sprintf("Code: %v, Error: %v, ID: %v",
		r.Code,
		r.Error,
		r.Id,
	)
}
