package goC8

import (
	collection_req2 "github.com/marvin-hansen/goC8/requests/collection_req"
	"strings"
)

type CollectionManager struct {
	client *Client
}

func NewCollectionManager(client *Client) *CollectionManager {
	return &CollectionManager{client: client}
}

// GetAllCollections
// Returns an object with an attribute collections containing an array of all collection descriptions.
// The same information is also available in the names as an object with the collection names as keys.
// By providing the optional query parameter excludeSystem with a value of true, all system collections will be excluded from the response.
// https://macrometa.com/docs/api#/operations/handleCommandGet
func (c CollectionManager) GetAllCollections(fabric string) (response *collection_req2.ResponseForGetAllCollections, err error) {
	req := collection_req2.NewRequestForGetAllCollections(fabric)
	response = collection_req2.NewResponseForGetAllCollections()
	if err = c.client.Request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

// CreateNewCollection
// Creates a new collection with a given name. The request must contain an object with the following attributes.
// name: The name of the collection.
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
// https://macrometa.com/docs/api#/operations/handleCommandPost:CreateCollection
func (c CollectionManager) CreateNewCollection(fabric, collectionName string, allowUserKeys bool, collectionType collection_req2.CollectionType) (err error) {
	req := collection_req2.NewRequestForCreateNewCollection(fabric, collectionName, allowUserKeys, collectionType)
	response := collection_req2.NewResponseForCreateNewCollection()
	if err = c.client.Request(req, response); err != nil {
		return err
	}
	return nil
}

// GetCollectionInfo
// Fetches the information about collection.
// https://macrometa.com/docs/api#/operations/handleCommandGet:collectionGetProperties
func (c CollectionManager) GetCollectionInfo(fabric, collectionName string) (response *collection_req2.ResponseForGetCollectionInfo, err error) {
	req := collection_req2.NewRequestForGetCollectionInfo(fabric, collectionName)
	response = collection_req2.NewResponseForGetCollectionInfo()
	if err = c.client.Request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

// CheckCollectionExists
// Returns true if the collection of the name exists. False otherwise.
func (c CollectionManager) CheckCollectionExists(fabric, collectionName string) (exists bool, err error) {
	req := collection_req2.NewRequestForGetCollectionInfo(fabric, collectionName)
	response := collection_req2.NewResponseForGetCollectionInfo()
	if err = c.client.Request(req, response); err != nil {
		if strings.Contains(err.Error(), "1203") { // Number=1203,  Error Message=collection or view not found
			return false, nil
		} else {
			return false, err // Any other error
		}
	}
	return true, nil
}

// CountCollection
// Returns the number of documents in the collection.
// Note: This always loads the collection into memory.
//
//The call returns a JSON object with at least the following attributes on success:
//error: false
//count: The number of documents inside the collection.
//id: The id of collection as string.
//name: The name of collection as string.
//isSystem: True if the collection is a system collection.
//searchEnabled: True if the collection is searchable.
//globallyUniqueId: Global unique identifier as string.
// https://macrometa.com/docs/api#/operations/handleCommandGet:getCollectionCount
func (c CollectionManager) CountCollection(fabric, collectionName string) (response *collection_req2.ResultFromCollection, err error) {
	req := collection_req2.NewRequestForCountCollection(fabric, collectionName)
	response = collection_req2.NewResultFromCollection()
	if err = c.client.Request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

// UpdateCollectionProperties updates collection properties.
// Note: except for waitForSync and hasStream, collection properties cannot be changed once a collection is created.
// Changes the properties of a collection.
//Requires a JSON object with these properties:
//
//hasStream: True if creating a live collection stream.
//waitForSync: True if waiting for documents to be synchronized to storage.
//The call returns a JSON object with at least the following attributes on success:
//id: The identifier of the collection.
//name: The name of the collection.
//status: The status of the collection as number.
//1: new born collection
//2: unloaded
//3: loaded
//4: in the process of being unloaded
//5: deleted
//6: loading
//Every other status indicates a corrupted collection.
//type: The type of the collection as number.
//2: document collection (normal case)
//3: edges collection
//isSystem: True then the collection is a system collection.
//stream: True if the collection has a local streams associate with it.
//Note: except for waitForSync and hasStream, collection properties cannot be changed once a collection is created.
// https://macrometa.com/docs/api#/operations/handleCommandPut:stream
func (c CollectionManager) UpdateCollectionProperties(fabric, collectionName string, properties *collection_req2.UpdateOptions) (response *collection_req2.ResponseForUpdateCollection, err error) {
	req := collection_req2.NewRequestForUpdateCollection(fabric, collectionName, properties)
	response = collection_req2.NewResponseForUpdateCollection()
	if err = c.client.Request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

// TruncateCollection
// Remove all documents from the collection but leaves the indexes intact.
// https://macrometa.com/docs/api#/operations/handleCommandPut:truncateCollection
func (c CollectionManager) TruncateCollection(fabric, collectionName string) (response *collection_req2.ResponseForTruncateCollection, err error) {
	req := collection_req2.NewRequestForTruncateCollection(fabric, collectionName)
	response = collection_req2.NewResponseForTruncateCollection()
	if err = c.client.Request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

// DeleteCollection
// Drops the collection identified by collection-name. If the collection was successfully dropped, an object is returned with the following attributes:
//error: false
//id: The identifier of the dropped collection.
// https://macrometa.com/docs/api#/operations/handleCommandDelete:collection
func (c CollectionManager) DeleteCollection(fabric, collectionName string, isSystem bool) (err error) {
	req := collection_req2.NewRequestForDeleteCollection(fabric, collectionName, isSystem)
	response := collection_req2.NewResponseForDeleteCollection()
	if err = c.client.Request(req, response); err != nil {
		return err
	}
	// valid response: Code: 200, Error: false, ID: 159XXXXX
	return nil
}
