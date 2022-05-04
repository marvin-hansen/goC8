package document_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForUpdateDocument(
	fabric, collectionName, key string,
	jsonDocument []byte,
	silent bool,
	parameters *UpdateDocumentParameters) *RequestForUpdateDocument {
	return &RequestForUpdateDocument{
		payload: jsonDocument,
		path: fmt.Sprintf("_fabric/%v/_api/document/%v/%v",
			fabric, collectionName,
			key,
		),
		parameters: fmt.Sprintf("?keepNull=%v&mergeObjects=%v&ignoreRevs=%v&returnOld=%v&returnNew=%v&silent=%v&waitForSync=%v",
			parameters.keepNull,
			parameters.mergeObjects,
			parameters.ignoreRevs,
			parameters.returnOld,
			parameters.returnNew,
			silent,
			parameters.waitForSync,
		),
	}
}

type RequestForUpdateDocument struct {
	path       string
	parameters string
	payload    []byte
}

func (req *RequestForUpdateDocument) Path() string {
	return req.path
}

func (req *RequestForUpdateDocument) Method() string {
	return http.MethodPatch
}

func (req *RequestForUpdateDocument) Query() string {
	return ""
}

func (req *RequestForUpdateDocument) HasQueryParameter() bool {
	return true
}

func (req *RequestForUpdateDocument) GetQueryParameter() string {
	return req.parameters
}

func (req *RequestForUpdateDocument) Payload() []byte {
	return req.payload
}

func (req *RequestForUpdateDocument) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForUpdateDocument() *ResponseForUpdateDocument {
	return new(ResponseForUpdateDocument)
}

type ResponseForUpdateDocument DocumentResult

func (r *ResponseForUpdateDocument) IsResponse() {}

func (r ResponseForUpdateDocument) String() string {
	return fmt.Sprintf("ID: %v, Key: %v, Ref: %v, OldRev: %v",
		r.Id,
		r.Key,
		r.Rev,
		r.OldRev,
	)
}
