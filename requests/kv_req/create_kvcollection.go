package kv_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForCreateKVCollection(fabric, collectionName string, expiration bool, options *CreateKVOptions) *RequestForCreateKVCollection {
	return &RequestForCreateKVCollection{
		path:       fmt.Sprintf("_fabric/%v/_api/kv/%v", fabric, collectionName),
		parameters: fmt.Sprintf("?expiration=%v", expiration),
		payload:    getCreateKVPayload(options),
	}
}

func getCreateKVPayload(opt *CreateKVOptions) []byte {
	str := fmt.Sprintf(`{
  "stream": %v,
  "enableShards": %v,
  "waitForSync": %v,
  "shardKeys": %v
}`, opt.Stream, opt.EnableShards, opt.WaitForSync, opt.ShardKeys)

	return []byte(str)
}

type RequestForCreateKVCollection struct {
	path       string
	parameters string
	payload    []byte
}

func (req *RequestForCreateKVCollection) Path() string {
	return req.path
}

func (req *RequestForCreateKVCollection) Method() string {
	return http.MethodPost
}

func (req *RequestForCreateKVCollection) Query() string {
	return ""
}

func (req *RequestForCreateKVCollection) HasQueryParameter() bool {
	return true
}

func (req *RequestForCreateKVCollection) GetQueryParameter() string {
	return req.parameters
}

func (req *RequestForCreateKVCollection) Payload() []byte {
	return req.payload
}

func (req *RequestForCreateKVCollection) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

// call NewKVResponse() instead
