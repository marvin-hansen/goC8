package qw_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForReadAllQueryWorkers(fabric string) *RequestForReadAllQueryWorkers {
	return &RequestForReadAllQueryWorkers{
		path: fmt.Sprintf("_fabric/%v/_api/restql/user", fabric),
	}
}

type RequestForReadAllQueryWorkers struct {
	path string
}

func (req *RequestForReadAllQueryWorkers) Path() string {
	return req.path
}

func (req *RequestForReadAllQueryWorkers) Method() string {
	return http.MethodGet
}

func (req *RequestForReadAllQueryWorkers) Query() string {
	return ""
}

func (req *RequestForReadAllQueryWorkers) HasQueryParameter() bool {
	return false
}

func (req *RequestForReadAllQueryWorkers) GetQueryParameter() string {
	return ""
}

func (req *RequestForReadAllQueryWorkers) Payload() []byte {
	return nil
}

func (req *RequestForReadAllQueryWorkers) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForReadAllQueryWorkers() *ResponseForReadAllQueryWorkers {
	return new(ResponseForReadAllQueryWorkers)
}

type ResponseForReadAllQueryWorkers struct {
	Code   int           `json:"code"`
	Error  bool          `json:"error"`
	Result []QueryWorker `json:"result"`
}

func (r *ResponseForReadAllQueryWorkers) IsResponse() {}

func (r ResponseForReadAllQueryWorkers) String() string {
	return fmt.Sprintf("Code: %v\n Error: %v\n Result: %v\n",
		r.Code,
		r.Error,
		r.Result,
	)
}
