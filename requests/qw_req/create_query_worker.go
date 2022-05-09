package qw_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForCreateQueryWorker(fabric, workerName, queryString, parameterString string) *RequestForCreateQueryWorker {
	return &RequestForCreateQueryWorker{
		path:    fmt.Sprintf("_fabric/%v/_api/restql", fabric),
		payload: getCreateWorkerPayload(workerName, queryString, parameterString),
	}
}

func getCreateWorkerPayload(workerName, queryString, parameterString string) []byte {
	str := fmt.Sprintf(`{
  "query": {
    "name": "%v",
    "parameter": {%v},
    "value": "%v"
  }
}`, workerName, parameterString, queryString)
	return []byte(str)
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

// Call NewResponseForQueryWorker()
