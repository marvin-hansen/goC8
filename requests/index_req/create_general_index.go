package index_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForCreateGeneralIndex(fabric, collectionName, field string) *RequestForCreateGeneralIndex {
	return &RequestForCreateGeneralIndex{
		path:       fmt.Sprintf("_fabric/%v/_api/index/geo", fabric),
		parameters: fmt.Sprintf("?collection=%v", collectionName),
		//payload:    getGeneralIndexPayload(field, geoJson)}
	}
}

type RequestForCreateGeneralIndex struct {
	path       string
	parameters string
	payload    []byte
}

func (req *RequestForCreateGeneralIndex) Path() string {
	return req.path
}

func (req *RequestForCreateGeneralIndex) Method() string {
	return http.MethodGet
}

func (req *RequestForCreateGeneralIndex) Query() string {
	return ""
}

func (req *RequestForCreateGeneralIndex) HasQueryParameter() bool {
	return false
}

func (req *RequestForCreateGeneralIndex) GetQueryParameter() string {
	return "" //"?excludeSystem=true"
}

func (req *RequestForCreateGeneralIndex) Payload() []byte {
	return nil
}

func (req *RequestForCreateGeneralIndex) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForCreateGeneralIndex() *ResponseForCreateGeneralIndex {
	return new(ResponseForCreateGeneralIndex)
}

type ResponseForCreateGeneralIndex struct {
	// @FIXME
	Field string
}

func (r *ResponseForCreateGeneralIndex) IsResponse() {}

func (r ResponseForCreateGeneralIndex) String() string {
	// @FIXME
	return fmt.Sprintf("Bootfile: %v", r.Field)
}
