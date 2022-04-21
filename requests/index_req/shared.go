package index_req

import "fmt"

type IndexEntry struct {
	Fields              []string `json:"fields,omitempty"`
	Id                  string   `json:"id,omitempty"`
	Name                string   `json:"name,omitempty"`
	SelectivityEstimate int      `json:"selectivityEstimate,omitempty"`
	Sparse              bool     `json:"sparse,omitempty"`
	Type                string   `json:"type,omitempty"`
	Unique              bool     `json:"unique,omitempty"`
}

func (r IndexEntry) String() string {
	return fmt.Sprintf(" Fields: %v\n ID: %v\n Name: %v\n SelectivityEstimate: %v\n Sparse: %v\n  Type: %v\n  Unique: %v\n",
		r.Fields,
		r.Id,
		r.Name,
		r.SelectivityEstimate,
		r.Sparse,
		r.Type,
		r.Unique,
	)
}

type Identifiers = IndexEntry

type Index = IndexEntry
