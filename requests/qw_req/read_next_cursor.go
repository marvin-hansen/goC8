package qw_req

import (
	"fmt"
	"github.com/marvin-hansen/goC8/requests/query_req"
	"net/http"
)

//**// Request //**//

func NewRequestForReadNextCursor(fabric, cursorID string) *RequestForReadNextCursor {
	return &RequestForReadNextCursor{
		path: fmt.Sprintf("_fabric/%v/_api/restql/fetch/%v", fabric, cursorID),
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
	return "" //"?excludeSystem=true"
}

func (req *RequestForReadNextCursor) Payload() []byte {
	return nil
}

func (req *RequestForReadNextCursor) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForReadNextCursor() *query_req.Cursor {
	return new(query_req.Cursor)
}
