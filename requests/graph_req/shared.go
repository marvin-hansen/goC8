package graph_req

import "fmt"

func NewResponseForGraph() *ResponseForGraph {
	return new(ResponseForGraph)
}

type ResponseForGraph struct {
	Code   int   `json:"code,omitempty"`
	Error  bool  `json:"error,omitempty"`
	Graphs Graph `json:"graph,omitempty"`
}

func (r *ResponseForGraph) IsResponse() {}

func (r ResponseForGraph) String() string {
	return fmt.Sprintf(" Code: %v \n Error: %v \n Graphs: %v",
		r.Code,
		r.Error,
		r.Graphs,
	)
}

type Graph struct {
	Id                   string            `json:"_id,omitempty"`
	Key                  string            `json:"_key,omitempty"`
	Rev                  string            `json:"_rev,omitempty"`
	EdgeDefinitions      []EdgeDefinitions `json:"edgeDefinitions,omitempty"`
	IsSmart              bool              `json:"isSmart,omitempty"`
	MinReplicationFactor int               `json:"minReplicationFactor,omitempty"`
	Name                 string            `json:"name,omitempty"`
	NumberOfShards       int               `json:"numberOfShards,omitempty"`
	OrphanCollections    []interface{}     `json:"orphanCollections,omitempty"`
	ReplicationFactor    int               `json:"replicationFactor,omitempty"`
}

func (r Graph) String() string {
	return fmt.Sprintf("ID: %v \n Key: %v \n Rec: %v \n EdgeDefinitions: %v \n IsSmart: %v \n MinReplicationFactor: %v \n Name: %v \n NumberOfShards: %v \n OrphanCollections: %v \n ReplicationFactor: %v \n",
		r.Id,
		r.Key,
		r.Rev,
		r.EdgeDefinitions,
		r.IsSmart,
		r.MinReplicationFactor,
		r.Name,
		r.NumberOfShards,
		r.OrphanCollections,
		r.ReplicationFactor,
	)
}

type EdgeDefinitions struct {
	Collection string   `json:"collection,omitempty"`
	From       []string `json:"from,omitempty"`
	To         []string `json:"to,omitempty"`
}

func (r EdgeDefinitions) String() string {
	return fmt.Sprintf("Collection: %v \n From: %v \n To: %v \n",
		r.Collection,
		r.From,
		r.To,
	)
}

func NewEdgeResponse() *EdgeResponse {
	return new(EdgeResponse)
}

type EdgeResponse struct {
	Code    int  `json:"code"`
	Error   bool `json:"error"`
	Edge    Edge `json:"edge,omitempty"`
	Old     Edge `json:"old"`
	Removed bool `json:"removed"`
}

func (r *EdgeResponse) IsResponse() {}

func (r EdgeResponse) String() string {
	return fmt.Sprintf(" Code: %v \n Error: %v \n Edge: %v",
		r.Code,
		r.Error,
		r.Edge,
	)
}

type Edge struct {
	Id     string `json:"_id,omitempty"`
	Key    string `json:"_key,omitempty"`
	Rev    string `json:"_rev,omitempty"`
	From   string `json:"_from,omitempty"`
	To     string `json:"_to,omitempty"`
	Online bool   `json:"online,omitempty"`
}

func (r Edge) String() string {
	return fmt.Sprintf("ID: %v \n Key: %v \n Ref: %v \n From: %v \n To: %v \n",
		r.Id,
		r.Key,
		r.Rev,
		r.From,
		r.To,
	)
}

type Vertex struct {
	Id  string `json:"_id"`
	Key string `json:"_key"`
	Rev string `json:"_rev"`
}

func (r Vertex) String() string {
	return fmt.Sprintf("ID: %v \n Key: %v \n Ref: %v",
		r.Id,
		r.Key,
		r.Rev,
	)
}
