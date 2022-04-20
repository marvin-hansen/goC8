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
	IsLocal      bool       `json:"isLocal"`
	IsSystem     bool       `json:"isSystem"`
	KeyOptions   KeyOptions `json:"keyOptions"`
	Name         string     `json:"name"`
	ShardKeys    []string   `json:"shardKeys"`
	EnableShards bool       `json:"enableShards"`
	Type         int        `json:"type"`
	Stream       bool       `json:"stream"`
}

type KeyOptions struct {
	AllowUserKeys bool   `json:"allowUserKeys"`
	Increment     int    `json:"increment"`
	Offset        int    `json:"offset"`
	Type          string `json:"type"`
}

type ResulFromCollections struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	Status           int    `json:"status"`
	Type             int    `json:"type"`
	CollectionModel  string `json:"collectionModel"`
	IsSpot           bool   `json:"isSpot"`
	IsLocal          bool   `json:"isLocal"`
	HasStream        bool   `json:"hasStream"`
	WaitForSync      bool   `json:"waitForSync"`
	IsSystem         bool   `json:"isSystem"`
	GloballyUniqueId string `json:"globallyUniqueId"`
	SearchEnabled    bool   `json:"searchEnabled"`
}

func (r ResulFromCollections) String() string {
	return fmt.Sprintf("ID: %v, Name: %v,  Status: %v, Type: %v, CollectionModel: %v, IsSpot: %v, IsLocal: %v, HasStream: %v, WaitForSync: %v, IsSystem: %v, GloballyUniqueId: %v, SearchEnabled: %v",
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
