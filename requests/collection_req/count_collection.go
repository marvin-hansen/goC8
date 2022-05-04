package collection_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForCountCollection(fabric, collectionName string) *RequestForCountCollection {
	return &RequestForCountCollection{
		path: fmt.Sprintf("_fabric/%v/_api/collection/%v/count", fabric, collectionName),
	}
}

type RequestForCountCollection struct {
	path string
}

func (req *RequestForCountCollection) Path() string {
	return req.path
}

func (req *RequestForCountCollection) Method() string {
	return http.MethodGet
}

func (req *RequestForCountCollection) Query() string {
	return ""
}

func (req *RequestForCountCollection) HasQueryParameter() bool {
	return false
}

func (req *RequestForCountCollection) GetQueryParameter() string {
	return ""
}

func (req *RequestForCountCollection) Payload() []byte {
	return nil
}

func (req *RequestForCountCollection) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

// call  NewResultFromCollection() instead
