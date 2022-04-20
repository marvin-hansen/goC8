package document_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

// NewRequestForGetDocument
// Returns the document identified by key. The returned document contains three special attributes:
//_id containing the document handle.
//_key containing key which uniquely identifies a document.
//_rev containing the revision.
func NewRequestForGetDocument(fabric, collectionName, key string) *RequestForGetDocument {
	return &RequestForGetDocument{
		path: fmt.Sprintf("_fabric/%v/_api/document/%v/%v",
			fabric, collectionName,
			key,
		),
	}
}

type RequestForGetDocument struct {
	path string
}

func (req *RequestForGetDocument) Path() string {
	return req.path
}

func (req *RequestForGetDocument) Method() string {
	return http.MethodGet
}

func (req *RequestForGetDocument) Query() string {
	return ""
}

func (req *RequestForGetDocument) HasQueryParameter() bool {
	return false
}

func (req *RequestForGetDocument) GetQueryParameter() string {
	return ""
}

func (req *RequestForGetDocument) Payload() []byte {
	return nil
}

func (req *RequestForGetDocument) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForGetJsonDocument() *ResponseForGetJsonDocument {
	return new(ResponseForGetJsonDocument)
}

type ResponseForGetJsonDocument struct {
	json []byte
}

func (r *ResponseForGetJsonDocument) IsJsonResponse() {}

func (r *ResponseForGetJsonDocument) SetJsonMessage(rawJson []byte) {
	r.json = rawJson
}

func (r *ResponseForGetJsonDocument) GetJsonMessage() []byte {
	return r.json
}

func (r *ResponseForGetJsonDocument) String() string {
	return string(r.json)
}
