package collection_req

import (
	"fmt"
	"net/http"
	"strconv"
)

//**// Request //**//

func NewRequestForDeleteCollection(fabric, collectionName string, isSystem bool) *RequestForDeleteCollection {
	return &RequestForDeleteCollection{
		path:       fmt.Sprintf("_fabric/%v/_api/collection/%v", fabric, collectionName),
		parameters: "?isSystem=" + strconv.FormatBool(isSystem),
	}
}

type RequestForDeleteCollection struct {
	path       string
	parameters string
}

func (req *RequestForDeleteCollection) Path() string {
	return req.path
}

func (req *RequestForDeleteCollection) Method() string {
	return http.MethodDelete
}

func (req *RequestForDeleteCollection) Query() string {
	return ""
}

func (req *RequestForDeleteCollection) HasQueryParameter() bool {
	return true
}

func (req *RequestForDeleteCollection) GetQueryParameter() string {
	return req.parameters
}

func (req *RequestForDeleteCollection) Payload() []byte {
	return nil
}

func (req *RequestForDeleteCollection) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForDeleteCollection() *ResponseForDeleteCollection {
	return new(ResponseForDeleteCollection)
}

type ResponseForDeleteCollection struct {
	Error bool   `json:"error"`
	Code  int    `json:"code"`
	Id    string `json:"id"`
}

func (r *ResponseForDeleteCollection) IsResponse() {}

func (r ResponseForDeleteCollection) String() string {
	return fmt.Sprintf("Code: %v, Error: %v, ID: %v",
		r.Code,
		r.Error,
		r.Id,
	)
}
