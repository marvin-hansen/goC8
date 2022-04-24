package index_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForCreatePersistentIndex(fabric, collectionName, field string, deduplicate, sparse, unique bool) *RequestForCreatePersistentIndex {
	indexType := "persistent"
	return &RequestForCreatePersistentIndex{
		path:       fmt.Sprintf("_fabric/%v/_api/index/%v", fabric, indexType),
		parameters: fmt.Sprintf("?collection=%v", collectionName),
		payload:    getIndexPayLoad(indexType, field, deduplicate, sparse, unique),
	}
}

type RequestForCreatePersistentIndex struct {
	path       string
	parameters string
	payload    []byte
}

func (req *RequestForCreatePersistentIndex) Path() string {
	return req.path
}

func (req *RequestForCreatePersistentIndex) Method() string {
	return http.MethodPost
}

func (req *RequestForCreatePersistentIndex) Query() string {
	return ""
}

func (req *RequestForCreatePersistentIndex) HasQueryParameter() bool {
	return true
}

func (req *RequestForCreatePersistentIndex) GetQueryParameter() string {
	return req.parameters
}

func (req *RequestForCreatePersistentIndex) Payload() []byte {
	return req.payload
}

func (req *RequestForCreatePersistentIndex) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForCreatePersistentIndex() *ResponseForCreatePersistentIndex {
	return new(ResponseForCreatePersistentIndex)
}

type ResponseForCreatePersistentIndex IndexEntry

func (r *ResponseForCreatePersistentIndex) IsResponse() {}

func (r ResponseForCreatePersistentIndex) String() string {
	return fmt.Sprintf(" Code: %v\n Error: %v\n GeoJson: %v\n Fields: %v\n MaxNumCoverCells: %v\n WorstIndexedLevel: %v\n Id: %v\n Name: %v\n SelectivityEstimate: %v\n Sparse: %v\n Type: %v\n, Unique: %v\n",
		r.Code,
		r.Error,
		r.GeoJson,
		r.Fields,
		r.MaxNumCoverCells,
		r.WorstIndexedLevel,
		r.Id,
		r.Name,
		r.SelectivityEstimate,
		r.Sparse,
		r.Type,
		r.Unique,
	)
}
