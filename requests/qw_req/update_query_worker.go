package qw_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForUpdateQueryWorker(fabric, workerName, queryString, bindVars string) *RequestForUpdateQueryWorker {
	return &RequestForUpdateQueryWorker{
		path:    fmt.Sprintf("_fabric/%v/_api/restql/%v", fabric, workerName),
		payload: getUpdateWorkerPayload(queryString, bindVars),
	}
}

func getUpdateWorkerPayload(queryString, parameterString string) []byte {
	str := fmt.Sprintf(`{
  "query": {
    "parameter": {%v},
    "value": "%v"
  }
}`, parameterString, queryString)
	return []byte(str)
}

type RequestForUpdateQueryWorker struct {
	path    string
	payload []byte
}

func (req *RequestForUpdateQueryWorker) Path() string {
	return req.path
}

func (req *RequestForUpdateQueryWorker) Method() string {
	return http.MethodPut
}

func (req *RequestForUpdateQueryWorker) Query() string {
	return ""
}

func (req *RequestForUpdateQueryWorker) HasQueryParameter() bool {
	return false
}

func (req *RequestForUpdateQueryWorker) GetQueryParameter() string {
	return ""
}

func (req *RequestForUpdateQueryWorker) Payload() []byte {
	return req.payload
}

func (req *RequestForUpdateQueryWorker) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

// Call NewResponseForQueryWorker()
