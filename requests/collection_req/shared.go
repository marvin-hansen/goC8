package collection_req

import (
	"fmt"
)

type UpdateOptions struct {
	HasStream   bool `json:"hasStream"`
	WaitForSync bool `json:"waitForSync"`
}

func NewCollectionOption(name string, allowUserKeys bool, collectionType CollectionType) *CollectionOption {
	return &CollectionOption{
		Name:     name,
		Type:     collectionType.ToInt(),
		IsLocal:  false,
		IsSystem: false,
		KeyOptions: KeyOptions{
			AllowUserKeys: allowUserKeys,
			Increment:     1,
			Offset:        1,
			Type:          "autoincrement",
		},
		ShardKeys:    []string{"_key"},
		EnableShards: false,
		Stream:       false,
	}
}

// CollectionOption
//name: The name of the collection.
//keyOptions:
//allowUserKeys: If set to true, you can supply a key value in the _key attribute of a document. If set to false, the key generator creates keys, and an error occurs if you add a _key value.
//type: Choose the type of key generator:
//traditional: generates numerical keys in ascending order.
//autoincrement: generates numerical keys in ascending order with configurable initial offset and spacing.
//padded: generates keys of a fixed length (16 bytes) in ascending lexicographical sort order.
//uuid64: generates universally unique keys.
//increment: increment value for autoincrement key generator. Not used for other key generator types.
//offset: Initial offset value for autoincrement key generator. Not used for other key generator types.
//isSystem: If true, create a system collection. The collection-name must start with an underscore. Not Applicable for end-users. (The default is false)
//isLocal: If true, create a local collection. For a local collection data is not replicated across regions. (The default is false)
//type: The type of the collection to be created.
//The following values for type are valid (The default is 2):
//2: document collection
//3: edge collection.
//stream: If true, create a local stream for collection. (The default is false)
//shardKeys: The specified shard key determines in which shard a given document is to be stored. Choosing the right shard key can have significant impact on your performance can reduce network traffic and increase performance.
type CollectionOption struct {
	IsLocal      bool       `json:"isLocal,omitempty"`
	IsSystem     bool       `json:"isSystem,omitempty"`
	KeyOptions   KeyOptions `json:"keyOptions,omitempty"`
	Name         string     `json:"name,omitempty"`
	ShardKeys    []string   `json:"shardKeys,omitempty"`
	EnableShards bool       `json:"enableShards,omitempty"`
	Type         int        `json:"type,omitempty"`
	Stream       bool       `json:"stream,omitempty"`
}

type KeyOptions struct {
	AllowUserKeys bool   `json:"allowUserKeys,omitempty"`
	Increment     int    `json:"increment,omitempty"`
	Offset        int    `json:"offset,omitempty"`
	Type          string `json:"type,omitempty"`
}

func NewResultFromCollection() *ResultFromCollection {
	return new(ResultFromCollection)
}

type ResultFromCollection struct {
	Code             int    `json:"code,omitempty"`
	Error            bool   `json:"error,omitempty"`
	Id               string `json:"id,omitempty"`
	Name             string `json:"name,omitempty"`
	Status           int    `json:"status,omitempty"`
	Type             int    `json:"type,omitempty"`
	CollectionModel  string `json:"collectionModel,omitempty"`
	IsSpot           bool   `json:"isSpot,omitempty"`
	IsLocal          bool   `json:"isLocal,omitempty"`
	HasStream        bool   `json:"hasStream,omitempty"`
	WaitForSync      bool   `json:"waitForSync,omitempty"`
	IsSystem         bool   `json:"isSystem,omitempty"`
	GloballyUniqueId string `json:"globallyUniqueId,omitempty"`
	SearchEnabled    bool   `json:"searchEnabled,omitempty"`
	Count            int    `json:"count,omitempty"`
}

func (r *ResultFromCollection) IsResponse() {}

func (r ResultFromCollection) String() string {
	return fmt.Sprintf(" ID: %v\n, Name: %v\n,  Status: %v\n, Type: %v\n, CollectionModel: %v\n, IsSpot: %v\n, IsLocal: %v\n, HasStream: %v\n, WaitForSync: %v\n, IsSystem: %v\n, GloballyUniqueId: %v\n, SearchEnabled: %v\n Count: %v",
		r.Id,
		r.Name,
		r.Status,
		r.Type,
		r.CollectionModel,
		r.IsSpot,
		r.IsLocal,
		r.HasStream,
		r.WaitForSync,
		r.IsSystem,
		r.GloballyUniqueId,
		r.SearchEnabled,
		r.Count,
	)
}

type Response200 struct {
	Code  int    `json:"code"`
	Error bool   `json:"error"`
	Id    string `json:"id"`
}

func (r Response200) String() string {
	return fmt.Sprintf("Code: %v, Error: %v, ID: %v",
		r.Code,
		r.Error,
		r.Id,
	)
}
