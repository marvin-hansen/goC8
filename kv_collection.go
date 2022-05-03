package goC8

import (
	"github.com/marvin-hansen/goC8/src/kv/kv_req"
	"time"
)

// GetAllKVCollections
// Lists all collections.
// https://macrometa.com/docs/api#/paths/_fabric-fabric--_api-kv/get
func (c Client) GetAllKVCollections(fabric string) (response *kv_req.KVResult, err error) {
	if benchmark {
		defer TimeTrack(time.Now(), "GetAllKVCollections")
	}

	req := kv_req.NewRequestForGetAllKVCollections(fabric)
	response = kv_req.NewKVResponse()
	err = c.request(req, response)
	return response, CheckError(err)
}

// CreateNewKVCollection
// Create key-value collection.
// https://macrometa.com/docs/api#/operations/CreateNamespace
func (c Client) CreateNewKVCollection(fabric, collectionName string, expiration bool, options *kv_req.CreateKVOptions) (response *kv_req.KVResult, err error) {
	if benchmark {
		defer TimeTrack(time.Now(), "CreateNewKVCollection")
	}

	if options == nil {
		options = kv_req.GetDefaultCreateKVOptions()
	}

	req := kv_req.NewRequestForCreateKVCollection(fabric, collectionName, expiration, options)
	response = kv_req.NewKVResponse()
	err = c.request(req, response)
	return response, CheckError(err)
}

// CountKVCollection
// Get number of key-value pairs in collection.
// https://macrometa.com/docs/api#/paths/_fabric-fabric--_api-kv--collection--count/get
func (c Client) CountKVCollection(fabric, collectionName string) (response *kv_req.KVResult, err error) {
	if benchmark {
		defer TimeTrack(time.Now(), "CountKVCollection")
	}

	req := kv_req.NewRequestForCount(fabric, collectionName)
	response = kv_req.NewKVResponse()
	err = c.request(req, response)
	return response, CheckError(err)
}

// TruncateKVCollection
// Remove all key-value pairs in a collection.
// https://macrometa.com/docs/api#/paths/_fabric-fabric--_api-kv--collection--truncate/put
func (c Client) TruncateKVCollection(fabric, collectionName string) (response *kv_req.KVResult, err error) {
	if benchmark {
		defer TimeTrack(time.Now(), "TruncateKVCollection")
	}

	req := kv_req.NewRequestForTruncateKVCollection(fabric, collectionName)
	response = kv_req.NewKVResponse()
	err = c.request(req, response)
	return response, CheckError(err)
}

// DeleteKVCollection
// Delete collection.
// https://macrometa.com/docs/api#/paths/_fabric-fabric--_api-kv--collection/delete
func (c Client) DeleteKVCollection(fabric, collectionName string) (response *kv_req.KVResult, err error) {
	if benchmark {
		defer TimeTrack(time.Now(), "DeleteKVCollection")
	}

	req := kv_req.NewRequestForDeleteKVCollection(fabric, collectionName)
	response = kv_req.NewKVResponse()
	err = c.request(req, response)
	return response, CheckError(err)
}
