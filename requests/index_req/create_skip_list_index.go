package index_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForCreateSkipListIndex(fabric, collectionName, field string, deduplicate, sparse, unique bool) *RequestForCreateSkipListIndex {
	indexType := "skiplist"

	return &RequestForCreateSkipListIndex{
		path:       fmt.Sprintf("_fabric/%v/_api/index/%v", fabric, indexType),
		parameters: fmt.Sprintf("?collection=%v", collectionName),
		payload:    getIndexPayLoad(indexType, field, deduplicate, sparse, unique),
	}
}

type RequestForCreateSkipListIndex struct {
	path       string
	parameters string
	payload    []byte
}

func (req *RequestForCreateSkipListIndex) Path() string {
	return req.path
}

func (req *RequestForCreateSkipListIndex) Method() string {
	return http.MethodPost
}

func (req *RequestForCreateSkipListIndex) Query() string {
	return ""
}

func (req *RequestForCreateSkipListIndex) HasQueryParameter() bool {
	return true
}

func (req *RequestForCreateSkipListIndex) GetQueryParameter() string {
	return req.parameters
}

func (req *RequestForCreateSkipListIndex) Payload() []byte {
	return req.payload
}

func (req *RequestForCreateSkipListIndex) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForCreateSkipListIndex() *ResponseForCreateSkipListIndex {
	return new(ResponseForCreateSkipListIndex)
}

type ResponseForCreateSkipListIndex IndexEntry

func (r *ResponseForCreateSkipListIndex) IsResponse() {}

func (r ResponseForCreateSkipListIndex) String() string {
	return fmt.Sprintf(" Code: %v\n Deduplicate: %v\n Error: %v\n Fields: %v\n  GeoJson: %v\n ID: %v\n Name: %v\n MinLength: %v\n SelectivityEstimate: %v\n Sparse: %v\n  Type: %v\n  Unique: %v\n MaxNumCoverCells: %v\n WorstIndexedLevel: %v\n",
		r.Code,
		r.Deduplicate,
		r.Error,
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
