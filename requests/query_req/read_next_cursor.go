package query_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForReadNextCursor(fabric, cursorID string) *RequestForReadNextCursor {
	return &RequestForReadNextCursor{
		path: fmt.Sprintf("_fabric/%v/_api/cursor/%v", fabric, cursorID),
	}
}

type RequestForReadNextCursor struct {
	path string
}

func (req *RequestForReadNextCursor) Path() string {
	return req.path
}

func (req *RequestForReadNextCursor) Method() string {
	return http.MethodPut
}

func (req *RequestForReadNextCursor) Query() string {
	return ""
}

func (req *RequestForReadNextCursor) HasQueryParameter() bool {
	return false
}

func (req *RequestForReadNextCursor) GetQueryParameter() string {
	return ""
}

func (req *RequestForReadNextCursor) Payload() []byte {
	return nil
}

func (req *RequestForReadNextCursor) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForReadNextCursor() *Cursor {
	return new(Cursor)
}
