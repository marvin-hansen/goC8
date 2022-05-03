package document_req

import (
	"bytes"
	"fmt"
	"net/http"
)

//**// Request //**//

// NewRequestForReplaceManyDocument
// Replaces multiple documents in the specified collection with the ones in the body. The replaced documents are specified by the _key attributes in the body documents.
// If ignoreRevs is false, a _rev attribute in each document body must match the revision of the corresponding document in the database. Otherwise, the call fails
// In case of an error or violated precondition, an error object with the attribute error set to true and the attribute errorCode set to the error code.
//
//If the query parameter returnOld is true, for each generated document the previous revision of the document is returned under the old attribute in the result.
//If the query parameter returnNew is true, for each generated document the new document is returned under the new attribute in the result.

func NewRequestForReplaceManyDocument(fabric, collectionName string, jsonDocument []byte, parameters *ReplaceDocumentParameters) *RequestForReplaceManyDocument {
	return &RequestForReplaceManyDocument{
		payload: jsonDocument,
		path: fmt.Sprintf("_fabric/%v/_api/document/%v",
			fabric, collectionName,
		),
		parameters: fmt.Sprintf("?ignoreRevs=%v&returnOld=%v&returnNew=%v&waitForSync=%v",
			parameters.ignoreRevs,
			parameters.returnOld,
			parameters.returnNew,
			parameters.waitForSync,
		),
	}
}

type RequestForReplaceManyDocument struct {
	path       string
	parameters string
	payload    []byte
}

func (req *RequestForReplaceManyDocument) Path() string {
	return req.path
}

func (req *RequestForReplaceManyDocument) Method() string {
	return http.MethodPut
}

func (req *RequestForReplaceManyDocument) Query() string {
	return ""
}

func (req *RequestForReplaceManyDocument) HasQueryParameter() bool {
	return true
}

func (req *RequestForReplaceManyDocument) GetQueryParameter() string {
	return req.parameters
}

func (req *RequestForReplaceManyDocument) Payload() []byte {
	return req.payload
}

func (req *RequestForReplaceManyDocument) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForReplaceManyDocument() *ResponseForReplaceManyDocument {
	return new(ResponseForReplaceManyDocument)
}

type ResponseForReplaceManyDocument []DocumentResult

func (r *ResponseForReplaceManyDocument) IsResponse() {}

func (r ResponseForReplaceManyDocument) String() string {
	var s bytes.Buffer
	for _, v := range r {
		s.WriteString(v.String())
		s.WriteString("/n")
	}
	return s.String()
}
