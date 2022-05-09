package qw_req

import (
	"fmt"
	"github.com/marvin-hansen/goC8/requests/query_req"
	"net/http"
)

//**// Request //**//

func NewRequestForRunQueryWorker(fabric, workerName, bindVars string) *RequestForRunQueryWorker {
	return &RequestForRunQueryWorker{
		path:    fmt.Sprintf("_fabric/%v/_api/restql/execute/%v", fabric, workerName),
		payload: getRunWorkerPayload(bindVars),
	}
}

func getRunWorkerPayload(parameterString string) []byte {
	str := fmt.Sprintf(`{
  "bindVars": {%v}
}`, parameterString)
	return []byte(str)
}

type RequestForRunQueryWorker struct {
	path    string
	payload []byte
}

func (req *RequestForRunQueryWorker) Path() string {
	return req.path
}

func (req *RequestForRunQueryWorker) Method() string {
	return http.MethodPost
}

func (req *RequestForRunQueryWorker) Query() string {
	return ""
}

func (req *RequestForRunQueryWorker) HasQueryParameter() bool {
	return false
}

func (req *RequestForRunQueryWorker) GetQueryParameter() string {
	return ""
}

func (req *RequestForRunQueryWorker) Payload() []byte {
	return req.payload
}

func (req *RequestForRunQueryWorker) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForRunQueryWorker() *query_req.Cursor {
	return new(query_req.Cursor)
}
