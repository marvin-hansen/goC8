package document_req

import (
	"bytes"
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForDeleteManyDocuments(fabric, collectionName string, keysToDelete []byte, parameters *DeleteDocumentParameters) *RequestForDeleteManyDocuments {
	return &RequestForDeleteManyDocuments{
		payload: keysToDelete,
		path: fmt.Sprintf("_fabric/%v/_api/document/%v",
			fabric, collectionName,
		),
		parameters: fmt.Sprintf("?returnOld=%v&ignoreRevs=%v&waitForSync=%v",
			parameters.returnOld,
			parameters.ignoreRevs,
			parameters.waitForSync,
		),
	}
}

type RequestForDeleteManyDocuments struct {
	path       string
	parameters string
	payload    []byte
}

func (req *RequestForDeleteManyDocuments) Path() string {
	return req.path
}

func (req *RequestForDeleteManyDocuments) Method() string {
	return http.MethodDelete
}

func (req *RequestForDeleteManyDocuments) Query() string {
	return ""
}

func (req *RequestForDeleteManyDocuments) HasQueryParameter() bool {
	return true
}

func (req *RequestForDeleteManyDocuments) GetQueryParameter() string {
	return req.parameters
}

func (req *RequestForDeleteManyDocuments) Payload() []byte {
	return req.payload
}

func (req *RequestForDeleteManyDocuments) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForDeleteManyDocuments() *ResponseForDeleteManyDocuments {
	return new(ResponseForDeleteManyDocuments)
}

type ResponseForDeleteManyDocuments []DocumentResult

func (r *ResponseForDeleteManyDocuments) IsResponse() {}

func (r ResponseForDeleteManyDocuments) String() string {
	var s bytes.Buffer
	for _, v := range r {
		s.WriteString(v.String())
		s.WriteString("/n")
	}
	return s.String()
}
