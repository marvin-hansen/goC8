package graph_req

import "fmt"

type Graph struct {
	Id                   string            `json:"_id"`
	Key                  string            `json:"_key"`
	Rev                  string            `json:"_rev"`
	EdgeDefinitions      []EdgeDefinitions `json:"edgeDefinitions"`
	IsSmart              bool              `json:"isSmart"`
	MinReplicationFactor int               `json:"minReplicationFactor"`
	Name                 string            `json:"name"`
	NumberOfShards       int               `json:"numberOfShards"`
	OrphanCollections    []interface{}     `json:"orphanCollections"`
	ReplicationFactor    int               `json:"replicationFactor"`
}

func (r Graph) String() string {
	return fmt.Sprintf("Name: %v, ID: %v, Key: %v, Rec: %v, EdgeDefinitions: %v, IsSmart: %v, MinReplicationFactor: %v  Name: %v,  NumberOfShards: %v, OrphanCollections: %v,  ReplicationFactor: %v",
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
	Collection string   `json:"collection"`
	From       []string `json:"from"`
	To         []string `json:"to"`
}

func (r EdgeDefinitions) String() string {
	return fmt.Sprintf("Collection: %v, From: %v, To: %v",
		r.Collection,
		r.From,
		r.To,
	)
}
