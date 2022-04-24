package index_req

import "fmt"

func getIndexPayLoad(indexType, field string, deduplicate, sparse, unique bool) []byte {
	str := fmt.Sprintf(`{
  "deduplicate": %v,
  "fields": [
    "%v"
  ],
  "sparse": %v,
  "type": "%v",
  "unique": %v
}
`,
		deduplicate, field, sparse, indexType, unique)
	return []byte(str)
}

type IndexEntry struct {
	Code                int      `json:"code,omitempty"`
	Deduplicate         bool     `json:"deduplicate,omitempty"`
	Error               bool     `json:"error,omitempty"`
	Fields              []string `json:"fields,omitempty"`
	GeoJson             bool     `json:"geoJson,omitempty"`
	Id                  string   `json:"id,omitempty"`
	Name                string   `json:"name,omitempty"`
	MinLength           int      `json:"minLength,omitempty"`
	SelectivityEstimate int      `json:"selectivityEstimate,omitempty"`
	Sparse              bool     `json:"sparse,omitempty"`
	Type                string   `json:"type,omitempty"`
	Unique              bool     `json:"unique,omitempty"`
	MaxNumCoverCells    int      `json:"maxNumCoverCells,omitempty"`
	WorstIndexedLevel   int      `json:"worstIndexedLevel,omitempty"`
}

func (r IndexEntry) String() string {
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

type Identifiers = IndexEntry

type Index = IndexEntry
