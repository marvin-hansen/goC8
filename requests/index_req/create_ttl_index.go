package index_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForCreateTTLIndex(fabric, collectionName, field string, expireAfter int) *RequestForCreateTTLIndex {
	indexType := "ttl"
	return &RequestForCreateTTLIndex{
		path:       fmt.Sprintf("_fabric/%v/_api/index/%v", fabric, indexType),
		parameters: fmt.Sprintf("?collection=%v", collectionName),
		payload:    getTTLPayload(indexType, field, expireAfter),
	}
}

func getTTLPayload(indexType, field string, expireAfter int) []byte {
	str := fmt.Sprintf(`
{
  "expireAfter": %v,
  "fields": [
    "%v"
  ],
  "type": "%v"
}
`, expireAfter, field, indexType)
	return []byte(str)
}

type RequestForCreateTTLIndex struct {
	path       string
	parameters string
	payload    []byte
}

func (req *RequestForCreateTTLIndex) Path() string {
	return req.path
}

func (req *RequestForCreateTTLIndex) Method() string {
	return http.MethodPost
}

func (req *RequestForCreateTTLIndex) Query() string {
	return ""
}

func (req *RequestForCreateTTLIndex) HasQueryParameter() bool {
	return true
}

func (req *RequestForCreateTTLIndex) GetQueryParameter() string {
	return req.parameters
}

func (req *RequestForCreateTTLIndex) Payload() []byte {
	return req.payload
}

func (req *RequestForCreateTTLIndex) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForCreateTTLIndex() *ResponseForCreateTTLIndex {
	return new(ResponseForCreateTTLIndex)
}

type ResponseForCreateTTLIndex IndexEntry

func (r *ResponseForCreateTTLIndex) IsResponse() {}

func (r ResponseForCreateTTLIndex) String() string {
	return fmt.Sprintf(" Code: %v\n Deduplicate: %v\n Error: %v\n ExpireAfter: %v\n Fields: %v\n  GeoJson: %v\n ID: %v\n Name: %v\n MinLength: %v\n SelectivityEstimate: %v\n Sparse: %v\n  Type: %v\n  Unique: %v\n MaxNumCoverCells: %v\n WorstIndexedLevel: %v\n",
		r.Code,
		r.Deduplicate,
		r.Error,
		r.ExpireAfter,
		r.Fields,
		r.GeoJson,
		r.Id,
		r.Name,
		r.MinLength,
		r.SelectivityEstimate,
		r.Sparse,
		r.Type,
		r.Unique,
		r.MaxNumCoverCells,
		r.WorstIndexedLevel,
	)
}
