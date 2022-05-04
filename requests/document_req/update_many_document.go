package document_req

import (
	"bytes"
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForUpdateManyDocuments(fabric, collectionName string, jsonDocument []byte, parameters *UpdateDocumentParameters) *RequestForUpdateManyDocument {
	return &RequestForUpdateManyDocument{
		payload: jsonDocument,
		path: fmt.Sprintf("_fabric/%v/_api/document/%v",
			fabric, collectionName,
		),
		parameters: fmt.Sprintf("?keepNull=%v&mergeObjects=%v&ignoreRevs=%v&returnOld=%v&returnNew=%v&waitForSync=%v",
			parameters.keepNull,
			parameters.mergeObjects,
			parameters.ignoreRevs,
			parameters.returnOld,
			parameters.returnNew,
			parameters.waitForSync,
		),
	}
}

type RequestForUpdateManyDocument struct {
	path       string
	parameters string
	payload    []byte
}

func (req *RequestForUpdateManyDocument) Path() string {
	return req.path
}

func (req *RequestForUpdateManyDocument) Method() string {
	return http.MethodPatch
}

func (req *RequestForUpdateManyDocument) Query() string {
	return ""
}

func (req *RequestForUpdateManyDocument) HasQueryParameter() bool {
	return true
}

func (req *RequestForUpdateManyDocument) GetQueryParameter() string {
	return req.parameters
}

func (req *RequestForUpdateManyDocument) Payload() []byte {
	return req.payload
}

func (req *RequestForUpdateManyDocument) ResponseCode() int {
	return 201 // ok
}

//**// Response //**//

func NewResponseForUpdateManyDocuments() *ResponseForUpdateManyDocument {
	return new(ResponseForUpdateManyDocument)
}

type ResponseForUpdateManyDocument []DocumentResult

func (r *ResponseForUpdateManyDocument) IsResponse() {}

func (r ResponseForUpdateManyDocument) String() string {
	var s bytes.Buffer
	for _, v := range r {
		s.WriteString(v.String())
		s.WriteString("/n")
	}
	return s.String()
}
