package document_req

import (
	"bytes"
	"fmt"
	"net/http"
)

//**// Request //**//

// NewRequestForCreateDocument
// silent - If set to false, the primary key of the new doc is returned. If set to true, an empty object is returned as response. No meta-data is returned for the created document. This option can be used to save some network traffic. True by default
// parameters - additional query parameters for non-standard cases.
// jsonDocument the document to store in the collection
func NewRequestForCreateDocument(fabric, collectionName string, silent bool, jsonDocument []byte, parameters *CreateDocumentParameters) *RequestForCreateDocument {
	return &RequestForCreateDocument{
		path:    fmt.Sprintf("_fabric/%v/_api/document/%v", fabric, collectionName),
		payload: jsonDocument,
		parameter: fmt.Sprintf("?returnNew=%v&returnOld=%v&silent=%v&overwrite=%v&waitForSync=%v",
			parameters.returnNew,
			parameters.returnOld,
			silent,
			parameters.overwrite,
			parameters.waitForSync,
		),
	}
}

type RequestForCreateDocument struct {
	path      string
	payload   []byte
	parameter string
}

func (req *RequestForCreateDocument) Path() string {
	return req.path
}

func (req *RequestForCreateDocument) Method() string {
	return http.MethodPost
}

func (req *RequestForCreateDocument) Query() string {
	return ""
}

func (req *RequestForCreateDocument) HasQueryParameter() bool {
	return true
}

func (req *RequestForCreateDocument) GetQueryParameter() string {
	return req.parameter
}

func (req *RequestForCreateDocument) Payload() []byte {
	return req.payload
}

func (req *RequestForCreateDocument) ResponseCode() int {
	return 202 // ok
}

//**// Response //**//

func NewResponseForCreateDocument() *ResponseForCreateDocument {
	return new(ResponseForCreateDocument)
}

type ResponseForCreateDocument []DocumentResult

func (r *ResponseForCreateDocument) IsResponse() {}

func (r ResponseForCreateDocument) String() string {
	var s bytes.Buffer
	for _, v := range r {
		s.WriteString(v.String())
		s.WriteString("/n")
	}
	return s.String()

}
