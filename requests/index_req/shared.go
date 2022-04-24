package index_req

import "fmt"

type IndexEntry struct {
	Error               bool     `json:"error,omitempty"`
	Code                int      `json:"code,omitempty"`
	Fields              []string `json:"fields,omitempty"`
	Id                  string   `json:"id,omitempty"`
	Name                string   `json:"name,omitempty"`
	MinLength           int      `json:"minLength,omitempty"`
	SelectivityEstimate int      `json:"selectivityEstimate,omitempty"`
	Sparse              bool     `json:"sparse,omitempty"`
	Type                string   `json:"type,omitempty"`
	Unique              bool     `json:"unique,omitempty"`
}

func (r IndexEntry) String() string {
	return fmt.Sprintf(" Error: %v\n Code: %v\n Fields: %v\n ID: %v\n Name: %v\n MinLength: %v\n SelectivityEstimate: %v\n Sparse: %v\n  Type: %v\n  Unique: %v\n",
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
