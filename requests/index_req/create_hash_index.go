package index_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForCreateHashIndex(fabric, collectionName, field string, deduplicate, sparse, unique bool) *RequestForCreateHashIndex {
	return &RequestForCreateHashIndex{
		path:       fmt.Sprintf("_fabric/%v/_api/index/hash", fabric),
		parameters: fmt.Sprintf("?collection=%v", collectionName),
		payload:    getHashPayLoad(field, deduplicate, sparse, unique),
	}
}

type RequestForCreateHashIndex struct {
	path       string
	parameters string
	payload    []byte
}

func getHashPayLoad(field string, deduplicate, sparse, unique bool) []byte {
	str := fmt.Sprintf(`{
  "deduplicate": %v,
  "fields": [
    "%v"
  ],
  "sparse": %v,
  "type": "hash",
  "unique": %v
}
`,
		deduplicate, field, sparse, unique)
	return []byte(str)
}

func (req *RequestForCreateHashIndex) Path() string {
	return req.path
}

func (req *RequestForCreateHashIndex) Method() string {
	return http.MethodPost
}

func (req *RequestForCreateHashIndex) Query() string {
	return ""
}

func (req *RequestForCreateHashIndex) HasQueryParameter() bool {
	return true
}

func (req *RequestForCreateHashIndex) GetQueryParameter() string {
	return req.parameters
}

func (req *RequestForCreateHashIndex) Payload() []byte {
	return req.payload
}

func (req *RequestForCreateHashIndex) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForCreateHashIndex() *ResponseForCreateHashIndex {
	return new(ResponseForCreateHashIndex)
}

type ResponseForCreateHashIndex IndexEntry

func (r *ResponseForCreateHashIndex) IsResponse() {}

func (r ResponseForCreateHashIndex) String() string {
	return fmt.Sprintf(" Deduplicate: %v\n Error: %v\n Code: %v\n Fields: %v\n ID: %v\n Name: %v\n MinLength: %v\n SelectivityEstimate: %v\n Sparse: %v\n  Type: %v\n  Unique: %v\n",
		r.Deduplicate,
		r.Error,
		r.Code,
		r.Fields,
		r.Id,
		r.Name,
		r.MinLength,
		r.SelectivityEstimate,
		r.Sparse,
		r.Type,
		r.Unique,
	)
}
