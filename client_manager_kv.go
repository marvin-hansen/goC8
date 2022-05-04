package goC8

import (
	kv_req2 "github.com/marvin-hansen/goC8/requests/kv_req"
	utils2 "github.com/marvin-hansen/goC8/utils"
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
func (c KVManager) GetAllKVCollections(fabric string) (response *kv_req2.KVResult, err error) {
	if benchmark {
		defer utils2.TimeTrack(time.Now(), "GetAllKVCollections")
	}

	req := kv_req2.NewRequestForGetAllKVCollections(fabric)
	response = kv_req2.NewKVResponse()
	err = c.client.Request(req, response)
	return response, CheckReturnError(err)
}

// CreateNewKVCollection
// Create key-value collection.
// https://macrometa.com/docs/api#/operations/CreateNamespace
func (c KVManager) CreateNewKVCollection(fabric, collectionName string, expiration bool, options *kv_req2.CreateKVOptions) (response *kv_req2.KVResult, err error) {
	if benchmark {
		defer utils2.TimeTrack(time.Now(), "CreateNewKVCollection")
	}

	if options == nil {
		options = kv_req2.GetDefaultCreateKVOptions()
	}

	req := kv_req2.NewRequestForCreateKVCollection(fabric, collectionName, expiration, options)
	response = kv_req2.NewKVResponse()
	err = c.client.Request(req, response)
	return response, CheckReturnError(err)
}

// CountKVCollection
// Get number of key-value pairs in collection.
// https://macrometa.com/docs/api#/paths/_fabric-fabric--_api-kv--collection--count/get
func (c KVManager) CountKVCollection(fabric, collectionName string) (response *kv_req2.KVResult, err error) {
	if benchmark {
		defer utils2.TimeTrack(time.Now(), "CountKVCollection")
	}

	req := kv_req2.NewRequestForCount(fabric, collectionName)
	response = kv_req2.NewKVResponse()
	err = c.client.Request(req, response)
	return response, CheckReturnError(err)
}

// TruncateKVCollection
// Remove all key-value pairs in a collection.
// https://macrometa.com/docs/api#/paths/_fabric-fabric--_api-kv--collection--truncate/put
func (c KVManager) TruncateKVCollection(fabric, collectionName string) (response *kv_req2.KVResult, err error) {
	if benchmark {
		defer utils2.TimeTrack(time.Now(), "TruncateKVCollection")
	}

	req := kv_req2.NewRequestForTruncateKVCollection(fabric, collectionName)
	response = kv_req2.NewKVResponse()
	err = c.client.Request(req, response)
	return response, CheckReturnError(err)
}

// DeleteKVCollection
// Delete collection.
// https://macrometa.com/docs/api#/paths/_fabric-fabric--_api-kv--collection/delete
func (c KVManager) DeleteKVCollection(fabric, collectionName string) (response *kv_req2.KVResult, err error) {
	if benchmark {
		defer utils2.TimeTrack(time.Now(), "DeleteKVCollection")
	}

	req := kv_req2.NewRequestForDeleteKVCollection(fabric, collectionName)
	response = kv_req2.NewKVResponse()
	err = c.client.Request(req, response)
	return response, CheckReturnError(err)
}

// SetKeyValuePairs
// Set one or more key-value pairs in key-value collection.
// If the input is an array of objects then key-value pairs are created in batch.
// If the key does not exist the key-value pairs are created. Otherwise the entry for the key is updated.
// Specify expiration in UTC timestamp.
// Max limit is 100 key-value pairs per request.
// https://macrometa.com/docs/api#/paths/_fabric-fabric--_api-kv--collection--value/put
func (c KVManager) SetKeyValuePairs(fabric, collectionName string, kvPairs kv_req2.KVPairCollection) (response *kv_req2.KVPairCollection, err error) {
	if benchmark {
		defer utils2.TimeTrack(time.Now(), "SetKeyValuePairs")
	}

	req := kv_req2.NewRequestForSetKeyValue(fabric, collectionName, kvPairs)
	response = kv_req2.NewKVPairCollection()
	err = c.client.Request(req, response)
	return response, CheckReturnError(err)
}

// GetAllKeys
// Get keys from key-value collection.
// https://macrometa.com/docs/api#/paths/_fabric-fabric--_api-kv--collection--keys/get
func (c KVManager) GetAllKeys(fabric, collectionName string, offset, limit int, order kv_req2.Order) (response *kv_req2.ResponseForGetAllKeys, err error) {
	if benchmark {
		defer utils2.TimeTrack(time.Now(), "GetAllKeys")
	}

	req := kv_req2.NewRequestForGetAllKeys(fabric, collectionName, offset, limit, order)
	response = kv_req2.NewResponseForGetAllKeys()
	err = c.client.Request(req, response)
	return response, CheckReturnError(err)
}

// GetAllValues
// Get key-value pairs from collection.
//Optional list of keys. Max limit is 100 keys per request.
// https://macrometa.com/docs/api#/operations/GetValues
func (c KVManager) GetAllValues(fabric, collectionName string, offset, limit int, keys kv_req2.KeyCollection) (response *kv_req2.ResponseForGetAllValues, err error) {
	if benchmark {
		defer utils2.TimeTrack(time.Now(), "GetAllValues")
	}

	req := kv_req2.NewRequestForGetAllValues(fabric, collectionName, keys, offset, limit)
	response = kv_req2.NewResponseForGetAllValues()
	err = c.client.Request(req, response)
	return response, CheckReturnError(err)
}

// DeleteKeyValuePairs
// Remove key-value collection.
// https://macrometa.com/docs/api#/paths/_fabric-fabric--_api-kv--collection--values/delete
func (c KVManager) DeleteKeyValuePairs(fabric, collectionName string, keys kv_req2.KeyCollection) (response *kv_req2.KVPairCollection, err error) {
	if benchmark {
		defer utils2.TimeTrack(time.Now(), "DeleteKeyValuePairs")
	}

	req := kv_req2.NewRequestForDeleteKeyValue(fabric, collectionName, keys)
	response = kv_req2.NewKVPairCollection()
	err = c.client.Request(req, response)
	return response, CheckReturnError(err)
}

// GetValue
// Get value from key-value collection.
// https://macrometa.com/docs/api#/paths/_fabric-fabric--_api-kv--collection--value--key/get
func (c KVManager) GetValue(fabric, collectionName, key string) (response *kv_req2.KVPair, err error) {
	if benchmark {
		defer utils2.TimeTrack(time.Now(), "GetValue")
	}

	req := kv_req2.NewRequestForGetValue(fabric, collectionName, key)
	response = kv_req2.NewEmptyKVPair()
	err = c.client.Request(req, response)
	return response, CheckReturnError(err)
}

// DeleteValue
// Remove key-value pair.
// https://macrometa.com/docs/api#/paths/_fabric-fabric--_api-kv--collection--value--key/delete
func (c KVManager) DeleteValue(fabric, collectionName, key string) (response *kv_req2.KVPair, err error) {
	if benchmark {
		defer utils2.TimeTrack(time.Now(), "DeleteValue")
	}

	req := kv_req2.NewRequestForDeleteValue(fabric, collectionName, key)
	response = kv_req2.NewEmptyKVPair()
	err = c.client.Request(req, response)
	return response, CheckReturnError(err)
}
