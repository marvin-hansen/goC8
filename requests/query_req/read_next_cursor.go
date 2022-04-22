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

func NewResponseForReadNextCursor() *ResponseForReadNextCursor {
	return new(ResponseForReadNextCursor)
}

type ResponseForReadNextCursor CursorResponse

func (r *ResponseForReadNextCursor) IsResponse() {}

func (r ResponseForReadNextCursor) String() string {
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
