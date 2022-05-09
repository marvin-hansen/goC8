package qw_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForCreateQueryWorker(fabric, workerName, queryString, parameterString string) *RequestForCreateQueryWorker {
	return &RequestForCreateQueryWorker{
		path:    fmt.Sprintf("_fabric/%v/_api/restql", fabric),
		payload: getQueryWorkerPayload(workerName, queryString, parameterString),
	}
}

type RequestForCreateQueryWorker struct {
	path    string
	payload []byte
}

func (req *RequestForCreateQueryWorker) Path() string {
	return req.path
}

func (req *RequestForCreateQueryWorker) Method() string {
	return http.MethodPost
}

func (req *RequestForCreateQueryWorker) Query() string {
	return ""
}

func (req *RequestForCreateQueryWorker) HasQueryParameter() bool {
	return false
}

func (req *RequestForCreateQueryWorker) GetQueryParameter() string {
	return "" //"?excludeSystem=true"
}

func (req *RequestForCreateQueryWorker) Payload() []byte {
	return req.payload
}

func (req *RequestForCreateQueryWorker) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForCreateQueryWorker() *ResponseForCreateQueryWorker {
	return new(ResponseForCreateQueryWorker)
}

type ResponseForCreateQueryWorker struct {
	Code   int         `json:"code"`
	Error  bool        `json:"error"`
	Result QueryWorker `json:"result"`
}

func (r *ResponseForCreateQueryWorker) IsResponse() {}

func (r ResponseForCreateQueryWorker) String() string {
	return fmt.Sprintf("Code: %v\n Error: %v\n Result: %v\n",
		r.Code,
		r.Error,
		r.Result,
	)
}
