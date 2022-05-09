package qw_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForDeleteQueryWorker(fabric, workerName string) *RequestForDeleteQueryWorker {
	return &RequestForDeleteQueryWorker{
		path: fmt.Sprintf("_fabric/%v/_api/restql/%v", fabric, workerName),
	}
}

type RequestForDeleteQueryWorker struct {
	path string
}

func (req *RequestForDeleteQueryWorker) Path() string {
	return req.path
}

func (req *RequestForDeleteQueryWorker) Method() string {
	return http.MethodDelete
}

func (req *RequestForDeleteQueryWorker) Query() string {
	return ""
}

func (req *RequestForDeleteQueryWorker) HasQueryParameter() bool {
	return false
}

func (req *RequestForDeleteQueryWorker) GetQueryParameter() string {
	return "" //"?excludeSystem=true"
}

func (req *RequestForDeleteQueryWorker) Payload() []byte {
	return nil
}

func (req *RequestForDeleteQueryWorker) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForDeleteQueryWorker() *ResponseForDeleteQueryWorker {
	return new(ResponseForDeleteQueryWorker)
}

type ResponseForDeleteQueryWorker struct {
	Code  int  `json:"code"`
	Error bool `json:"error"`
}

func (r *ResponseForDeleteQueryWorker) IsResponse() {}

func (r ResponseForDeleteQueryWorker) String() string {
	return fmt.Sprintf("Code: %v Error: %v", r.Code, r.Error)
}
