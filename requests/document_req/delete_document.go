package document_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForDeleteDocument(fabric, collectionName, key string, parameters *DeleteDocumentParameters) *RequestForDeleteDocument {

	if parameters == nil {
		parameters = GetDefaultDeleteDocumentParameters()
	}

	return &RequestForDeleteDocument{
		path: fmt.Sprintf("_fabric/%v/_api/document/%v/%v",
			fabric, collectionName,
			key,
		),
		parameters: fmt.Sprintf("?returnOld=%v&silent=%v&waitForSync=%v",
			parameters.returnOld,
			parameters.silent,
			parameters.waitForSync,
		),
	}
}

type RequestForDeleteDocument struct {
	path       string
	parameters string
}

func (req *RequestForDeleteDocument) Path() string {
	return req.path
}

func (req *RequestForDeleteDocument) Method() string {
	return http.MethodDelete
}

func (req *RequestForDeleteDocument) Query() string {
	return ""
}

func (req *RequestForDeleteDocument) HasQueryParameter() bool {
	return true
}

func (req *RequestForDeleteDocument) GetQueryParameter() string {
	return req.parameters
}

func (req *RequestForDeleteDocument) Payload() []byte {
	return nil
}

func (req *RequestForDeleteDocument) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForDeleteDocument() *ResponseForDeleteDocument {
	return new(ResponseForDeleteDocument)
}

type ResponseForDeleteDocument DocumentResult

func (r *ResponseForDeleteDocument) IsResponse() {}

func (r ResponseForDeleteDocument) String() string {
	return fmt.Sprintf("ID: %v, Key: %v, Ref: %v",
		r.Id,
		r.Key,
		r.Rev,
	)
}
