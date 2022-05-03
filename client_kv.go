package goC8

import (
	"github.com/marvin-hansen/goC8/src/requests/kv_req"
	"time"
)

type KVManager struct {
	client *Client
}

func NewKVManager(client *Client) *KVManager {
	return &KVManager{client: client}
}

// GetAllKVCollections
// Lists all collections.
// https://macrometa.com/docs/api#/paths/_fabric-fabric--_api-kv/get
func (c KVManager) GetAllKVCollections(fabric string) (response *kv_req.KVResult, err error) {
	if benchmark {
		defer TimeTrack(time.Now(), "GetAllKVCollections")
	}

	req := kv_req.NewRequestForGetAllKVCollections(fabric)
	response = kv_req.NewKVResponse()
	err = c.client.Request(req, response)
	return response, CheckError(err)
}

// CreateNewKVCollection
// Create key-value collection.
// https://macrometa.com/docs/api#/operations/CreateNamespace
func (c KVManager) CreateNewKVCollection(fabric, collectionName string, expiration bool, options *kv_req.CreateKVOptions) (response *kv_req.KVResult, err error) {
	if benchmark {
		defer TimeTrack(time.Now(), "CreateNewKVCollection")
	}

	if options == nil {
		options = kv_req.GetDefaultCreateKVOptions()
	}

	req := kv_req.NewRequestForCreateKVCollection(fabric, collectionName, expiration, options)
	response = kv_req.NewKVResponse()
	err = c.client.Request(req, response)
	return response, CheckError(err)
}

// CountKVCollection
// Get number of key-value pairs in collection.
// https://macrometa.com/docs/api#/paths/_fabric-fabric--_api-kv--collection--count/get
func (c KVManager) CountKVCollection(fabric, collectionName string) (response *kv_req.KVResult, err error) {
	if benchmark {
		defer TimeTrack(time.Now(), "CountKVCollection")
	}

	req := kv_req.NewRequestForCount(fabric, collectionName)
	response = kv_req.NewKVResponse()
	err = c.client.Request(req, response)
	return response, CheckError(err)
}

// TruncateKVCollection
// Remove all key-value pairs in a collection.
// https://macrometa.com/docs/api#/paths/_fabric-fabric--_api-kv--collection--truncate/put
func (c KVManager) TruncateKVCollection(fabric, collectionName string) (response *kv_req.KVResult, err error) {
	if benchmark {
		defer TimeTrack(time.Now(), "TruncateKVCollection")
	}

	req := kv_req.NewRequestForTruncateKVCollection(fabric, collectionName)
	response = kv_req.NewKVResponse()
	err = c.client.Request(req, response)
	return response, CheckError(err)
}

// DeleteKVCollection
// Delete collection.
// https://macrometa.com/docs/api#/paths/_fabric-fabric--_api-kv--collection/delete
func (c KVManager) DeleteKVCollection(fabric, collectionName string) (response *kv_req.KVResult, err error) {
	if benchmark {
		defer TimeTrack(time.Now(), "DeleteKVCollection")
	}

	req := kv_req.NewRequestForDeleteKVCollection(fabric, collectionName)
	response = kv_req.NewKVResponse()
	err = c.client.Request(req, response)
	return response, CheckError(err)
}

// SetKeyValuePairs
// Set one or more key-value pairs in key-value collection.
// If the input is an array of objects then key-value pairs are created in batch.
// If the key does not exist the key-value pairs are created. Otherwise the entry for the key is updated.
// Specify expiration in UTC timestamp.
// Max limit is 100 key-value pairs per request.
// https://macrometa.com/docs/api#/paths/_fabric-fabric--_api-kv--collection--value/put
func (c KVManager) SetKeyValuePairs(fabric, collectionName string, kvPairs kv_req.KVPairCollection) (response *kv_req.KVPairCollection, err error) {
	if benchmark {
		defer TimeTrack(time.Now(), "SetKeyValuePairs")
	}

	req := kv_req.NewRequestForSetKeyValue(fabric, collectionName, kvPairs)
	response = kv_req.NewKVPairCollection()
	err = c.client.Request(req, response)
	return response, CheckError(err)
}

// GetAllKeys
// Get keys from key-value collection.
// https://macrometa.com/docs/api#/paths/_fabric-fabric--_api-kv--collection--keys/get
func (c KVManager) GetAllKeys(fabric, collectionName string, offset, limit int, order kv_req.Order) (response *kv_req.ResponseForGetAllKeys, err error) {
	if benchmark {
		defer TimeTrack(time.Now(), "GetAllKeys")
	}

	req := kv_req.NewRequestForGetAllKeys(fabric, collectionName, offset, limit, order)
	response = kv_req.NewResponseForGetAllKeys()
	err = c.client.Request(req, response)
	return response, CheckError(err)
}

// GetAllValues
// Get key-value pairs from collection.
//Optional list of keys. Max limit is 100 keys per request.
// https://macrometa.com/docs/api#/operations/GetValues
func (c KVManager) GetAllValues(fabric, collectionName string, offset, limit int, keys kv_req.KeyCollection) (response *kv_req.ResponseForGetAllValues, err error) {
	if benchmark {
		defer TimeTrack(time.Now(), "GetAllValues")
	}

	req := kv_req.NewRequestForGetAllValues(fabric, collectionName, keys, offset, limit)
	response = kv_req.NewResponseForGetAllValues()
	err = c.client.Request(req, response)
	return response, CheckError(err)
}

// DeleteKeyValuePairs
// Remove key-value collection.
// https://macrometa.com/docs/api#/paths/_fabric-fabric--_api-kv--collection--values/delete
func (c KVManager) DeleteKeyValuePairs(fabric, collectionName string, keys kv_req.KeyCollection) (response *kv_req.KVPairCollection, err error) {
	if benchmark {
		defer TimeTrack(time.Now(), "DeleteKeyValuePairs")
	}

	req := kv_req.NewRequestForDeleteKeyValue(fabric, collectionName, keys)
	response = kv_req.NewKVPairCollection()
	err = c.client.Request(req, response)
	return response, CheckError(err)
}

// GetValue
// Get value from key-value collection.
// https://macrometa.com/docs/api#/paths/_fabric-fabric--_api-kv--collection--value--key/get
func (c KVManager) GetValue(fabric, collectionName, key string) (response *kv_req.KVPair, err error) {
	if benchmark {
		defer TimeTrack(time.Now(), "GetValue")
	}

	req := kv_req.NewRequestForGetValue(fabric, collectionName, key)
	response = kv_req.NewEmptyKVPair()
	err = c.client.Request(req, response)
	return response, CheckError(err)
}

// DeleteValue
// Remove key-value pair.
// https://macrometa.com/docs/api#/paths/_fabric-fabric--_api-kv--collection--value--key/delete
func (c KVManager) DeleteValue(fabric, collectionName, key string) (response *kv_req.KVPair, err error) {
	if benchmark {
		defer TimeTrack(time.Now(), "DeleteValue")
	}

	req := kv_req.NewRequestForDeleteValue(fabric, collectionName, key)
	response = kv_req.NewEmptyKVPair()
	err = c.client.Request(req, response)
	return response, CheckError(err)
}
