package goC8

import (
	"github.com/marvin-hansen/goC8/requests/index_req"
	"github.com/marvin-hansen/goC8/utils"
	"time"
)

type IndexManager struct {
	client *Client
}

func NewIndexManager(client *Client) *IndexManager {
	return &IndexManager{client: client}
}

// GetIndexes
// Fetches the list of all indexes of a collection.
// https://macrometa.com/docs/api#/operations/getIndexes
func (c IndexManager) GetIndexes(fabric, collectionName string) (response *index_req.ResponseForGetAllIndices, err error) {
	if benchmark {
		defer utils.TimeTrack(time.Now(), "GetIndexes")
	}

	req := index_req.NewRequestForGetAllIndices(fabric, collectionName)
	response = index_req.NewResponseForGetAllIndices()
	if err = c.client.Request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetIndex
// Fetches the information about index.
// https://macrometa.com/docs/api#/operations/getIndexes:handle
func (c IndexManager) GetIndex(fabric, collectionName, indexName string) (response *index_req.IndexEntry, err error) {
	if benchmark {
		defer utils.TimeTrack(time.Now(), "GetIndex")
	}

	req := index_req.NewRequestForGetIndex(fabric, collectionName, indexName)
	response = index_req.NewResponseForGetIndex()
	if err = c.client.Request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

// CreateFulltextIndex
// Creates fulltext index, if it does not already exist.
// field: Attribute to index. Must be of text format.
// minLength: Minimum character length of words to index.
func (c IndexManager) CreateFulltextIndex(fabric, collectionName, field string, minLength int) (response *index_req.IndexEntry, err error) {
	if benchmark {
		defer utils.TimeTrack(time.Now(), "CreateFulltextIndex")
	}

	req := index_req.NewRequestForCreateFulltextIndex(fabric, collectionName, field, minLength)
	response = index_req.NewResponseForCreateFulltextIndex()
	if err = c.client.Request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

// CreateHashIndex
// * fields (string): an array of attribute paths.
// * unique: if true, then create a unique index.
// * sparse: if true, then create a sparse index.
// * deduplicate: if false, the deduplication of array values is turned off.
// Creates a hash index for the collection collection-name if it does not already exist. The call expects an object containing the index details.
//In a sparse index all documents will be excluded from the index that do not contain at least one of the specified index attributes (i.e. fields) or that have a value of null in any of the specified index attributes. Such documents will not be indexed, and not be taken into account for uniqueness checks if the unique flag is set.
//In a non-sparse index, these documents will be indexed (for non-present indexed attributes, a value of null will be used) and will be taken into account for uniqueness checks if the unique flag is set.
//Note: unique indexes on non-shard keys are not supported in a cluster.
// https://macrometa.com/docs/api#/operations/createIndex:hash
func (c IndexManager) CreateHashIndex(fabric, collectionName, field string, deduplicate, sparse, unique bool) (response *index_req.IndexEntry, err error) {
	if benchmark {
		defer utils.TimeTrack(time.Now(), "CreateHashIndex")
	}

	req := index_req.NewRequestForCreateHashIndex(fabric, collectionName, field, deduplicate, sparse, unique)
	response = index_req.NewResponseForCreateHashIndex()
	if err = c.client.Request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

// CreateGeoIndex
// Creates a geo index.
// Note: Geo indexes are always sparse, meaning that documents that do not contain the index attributes or have non-numeric values in the index attributes will not be indexed.
func (c IndexManager) CreateGeoIndex(fabric, collectionName, field string, geoJson bool) (response *index_req.IndexEntry, err error) {
	if benchmark {
		defer utils.TimeTrack(time.Now(), "CreateGeoIndex")
	}

	req := index_req.NewRequestForCreateGeoIndex(fabric, collectionName, field, geoJson)
	response = index_req.NewResponseForCreateGeoIndex()
	if err = c.client.Request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

// CreatePersistentIndex
// Creates a persistent index for the collection collection-name, if it does not already exist. The call expects an object containing the index details.
// field (string): An array of attribute paths.
// unique: True if the index is unique.
// type: Must be equal to "persistent".
// sparse: True if the index is sparse type.
// deduplicate: It controls whether inserting duplicate index values from the same document into a unique array index will lead to a unique constraint error or not. (The default is true)
//Note:
//In a sparse index all documents are excluded from the index that do not contain at least one of the specified index attributes (i.e. fields) or that have a value of null in any of the specified index attributes. Such documents are not indexed and are not taken into account for uniqueness checks if the unique flag is set.
//In a non-sparse index, these documents are indexed (for non-present indexed attributes, a value of null is used) and are taken into account for uniqueness checks if the unique flag is set.
//unique indexes on non-shard keys are not supported in a cluster.
// https://macrometa.com/docs/api#/operations/createIndex:persistent
func (c IndexManager) CreatePersistentIndex(fabric, collectionName, field string, deduplicate, sparse, unique bool) (response *index_req.IndexEntry, err error) {
	if benchmark {
		defer utils.TimeTrack(time.Now(), "CreatePersistentIndex")
	}

	req := index_req.NewRequestForCreatePersistentIndex(fabric, collectionName, field, deduplicate, sparse, unique)
	response = index_req.NewResponseForCreatePersistentIndex()
	if err = c.client.Request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

// CreateSkipListIndex
// Creates a skiplist index for the collection collection-name, if it does not already exist. The call expects an object containing the index details.
// fields (string): an array of attribute paths.
//unique: if true, then create a unique index.
//type: must be equal to "skiplist".
//sparse: if true, then create a sparse index.
//deduplicate: if false, the deduplication of array values is turned off.
//
// In a sparse index all documents will be excluded from the index that do not contain at least one of the specified index attributes (i.e. fields) or that have a value of null in any of the specified index attributes. Such documents will not be indexed, and not be taken into account for uniqueness checks if the unique flag is set.
// In a non-sparse index, these documents will be indexed (for non-present indexed attributes, a value of null will be used) and will be taken into account for uniqueness checks if the unique flag is set.
//Note: unique indexes on non-shard keys are not supported in a cluster.
// https://macrometa.com/docs/api#/operations/createIndex:skiplist
func (c IndexManager) CreateSkipListIndex(fabric, collectionName, field string, deduplicate, sparse, unique bool) (response *index_req.IndexEntry, err error) {
	if benchmark {
		defer utils.TimeTrack(time.Now(), "CreateSkipListIndex")
	}

	req := index_req.NewRequestForCreateSkipListIndex(fabric, collectionName, field, deduplicate, sparse, unique)
	response = index_req.NewResponseForCreateSkipListIndex()
	if err = c.client.Request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

// CreateTTLIndex
// Creates a TTL index for the collection collection-name if it does not already exist. The call expects an object containing the index details.
// fields (string): An array with exactly one attribute path.
// type: Must be equal to "ttl".
// expireAfter: The time (in seconds) after a document's creation after which the documents count as "expired".
// https://macrometa.com/docs/api#/operations/createIndex:ttl
func (c IndexManager) CreateTTLIndex(fabric, collectionName, field string, expireAfter int) (response *index_req.IndexEntry, err error) {
	if benchmark {
		defer utils.TimeTrack(time.Now(), "CreateGeoIndex")
	}

	req := index_req.NewRequestForCreateTTLIndex(fabric, collectionName, field, expireAfter)
	response = index_req.NewResponseForCreateTTLIndex()
	if err = c.client.Request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

// DeleteIndex
// Remove an index.
// https://macrometa.com/docs/api#/operations/dropIndex
func (c IndexManager) DeleteIndex(fabric, collectionName, indexName string) (response *index_req.ResponseForDeleteIndex, err error) {
	if benchmark {
		defer utils.TimeTrack(time.Now(), "DeleteIndex")
	}

	req := index_req.NewRequestForDeleteIndex(fabric, collectionName, indexName)
	response = index_req.NewResponseForDeleteIndex()
	if err = c.client.Request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}
