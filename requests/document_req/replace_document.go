package document_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

// NewRequestForReplaceDocument
// Replaces the document with key with the one in the body, provided there is such a document and no precondition is violated.
// If the document exists and can be updated, then a HTTP 201 or a HTTP 202 is returned (depending on waitForSync, see below), the Etag header field contains the new revision of the document and the Location header contains a URL under which the document can be queried.
//
//If silent is set to false, the body of the response contains a JSON object with the information about the handle and the revision.
//
//_id contains the known document-handle of each document.
//_key contains the key that uniquely identifies a document.
//_rev contains the new document revision.
//If the query parameter returnOld is true, then the previous revision of the document is returned under the old attribute in the result.
//
//If the query parameter returnNew is true, then the new document is returned under the new attribute in the result.
//
//If the document does not exist, then an HTTP 404 is returned and the body of the response contains an error document.
func NewRequestForReplaceDocument(fabric, collectionName, key string, jsonDocument []byte, parameters *ReplaceDocumentParameters) *RequestForReplaceDocument {
	return &RequestForReplaceDocument{
		payload: jsonDocument,
		path: fmt.Sprintf("_fabric/%v/_api/document/%v/%v",
			fabric, collectionName, key,
		),
		parameters: fmt.Sprintf("?ignoreRevs=%v&returnOld=%v&returnNew=%v&waitForSync=%v",
			parameters.ignoreRevs,
			parameters.returnOld,
			parameters.returnNew,
			parameters.waitForSync,
		),
	}
}

type RequestForReplaceDocument struct {
	path       string
	parameters string
	payload    []byte
}

func (req *RequestForReplaceDocument) Path() string {
	return req.path
}

func (req *RequestForReplaceDocument) Method() string {
	return http.MethodPut
}

func (req *RequestForReplaceDocument) Query() string {
	return ""
}

func (req *RequestForReplaceDocument) HasQueryParameter() bool {
	return true
}

func (req *RequestForReplaceDocument) GetQueryParameter() string {
	return req.parameters
}

func (req *RequestForReplaceDocument) Payload() []byte {
	return req.payload
}

func (req *RequestForReplaceDocument) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForReplaceDocument() *ResponseForReplaceDocument {
	return new(ResponseForReplaceDocument)
}

type ResponseForReplaceDocument DocumentResult

func (r *ResponseForReplaceDocument) IsResponse() {}

func (r ResponseForReplaceDocument) String() string {
	return fmt.Sprintf("ID: %v, Key: %v, Ref: %v, OldRev: %v",
		r.Id,
		r.Key,
		r.Rev,
		r.OldRev,
	)
}
