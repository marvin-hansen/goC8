package index_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForCreateFulltextIndex(fabric, collectionName, field string, minLength int) *RequestForCreateFulltextIndex {
	return &RequestForCreateFulltextIndex{
		path:       fmt.Sprintf("_fabric/%v/_api/index/fulltext", fabric),
		parameters: fmt.Sprintf("?collection=%v", collectionName),
		payload:    getFulltextPayLoad(field, minLength),
	}
}

// getFulltextPayLoad constructs a JSON object with these properties:
// * fields: An array of attribute names. Currently, the array is limited to exactly one attribute.
// * type: Must be equal to "fulltext".
// * minLength: Minimum character length of words to index.Default take server-defined value if unspecified. Thus it is recommended to set this value explicitly when creating the index.
func getFulltextPayLoad(field string, minLength int) []byte {
	str := fmt.Sprintf(`{
	  "fields": [
		"%v"
	  ],
	  "minLength": %v,
	  "type": "fulltext"
}`,
		field, minLength)
	return []byte(str)
}

type RequestForCreateFulltextIndex struct {
	path       string
	parameters string
	payload    []byte
}

func (req *RequestForCreateFulltextIndex) Path() string {
	return req.path
}

func (req *RequestForCreateFulltextIndex) Method() string {
	return http.MethodPost
}

func (req *RequestForCreateFulltextIndex) Query() string {
	return ""
}

func (req *RequestForCreateFulltextIndex) HasQueryParameter() bool {
	return true
}

func (req *RequestForCreateFulltextIndex) GetQueryParameter() string {
	return req.parameters
}

func (req *RequestForCreateFulltextIndex) Payload() []byte {
	return req.payload
}

func (req *RequestForCreateFulltextIndex) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForCreateFulltextIndex() *ResponseForCreateFulltextIndex {
	return new(ResponseForCreateFulltextIndex)
}

type ResponseForCreateFulltextIndex struct {
	Code                int      `json:"code,omitempty"`
	Error               bool     `json:"error,omitempty"`
	Fields              []string `json:"fields,omitempty"`
	Id                  string   `json:"id,omitempty"`
	Name                string   `json:"name,omitempty"`
	SelectivityEstimate int      `json:"selectivityEstimate,omitempty"`
	Sparse              bool     `json:"sparse,omitempty"`
	Type                string   `json:"type,omitempty"`
	Unique              bool     `json:"unique,omitempty"`
}

func (r *ResponseForCreateFulltextIndex) IsResponse() {}

func (r ResponseForCreateFulltextIndex) String() string {
	return fmt.Sprintf(" Code: %v\n Error: %v\n Fields: %v\n Id: %v\n Name: %v\n SelectivityEstimate: %v\n Sparse: %v\n Type: %v\n, Unique: %v\n",
		r.Code,
		r.Error,
		r.Fields,
		r.Id,
		r.Name,
		r.SelectivityEstimate,
		r.Sparse,
		r.Type,
		r.Unique,
	)
}
