package goC8

import (
	"github.com/marvin-hansen/goC8/requests/kv_req"
	"time"
)

// SetKeyValuePairs
// Set one or more key-value pairs in key-value collection.
// If the input is an array of objects then key-value pairs are created in batch.
// If the key does not exist the key-value pairs are created. Otherwise the entry for the key is updated.
// Specify expiration in UTC timestamp.
// Max limit is 100 key-value pairs per request.
// https://macrometa.com/docs/api#/paths/_fabric-fabric--_api-kv--collection--value/put
func (c Client) SetKeyValuePairs(fabric, collectionName string, kvPairs kv_req.KVPairCollection) (response *kv_req.KVPairCollection, err error) {
	if benchmark {
		defer TimeTrack(time.Now(), "SetKeyValuePairs")
	}

	req := kv_req.NewRequestForSetKeyValue(fabric, collectionName, kvPairs)
	response = kv_req.NewKVPairCollection()
	err = c.request(req, response)
	return response, CheckError(err)
}

func (c Client) GetAllKeys(fabric, collectionName string, offset, limit int, order kv_req.Order) (response *kv_req.ResponseForGetAllKeys, err error) {
	if benchmark {
		defer TimeTrack(time.Now(), "GetAllKeys")
	}

	req := kv_req.NewRequestForGetAllKeys(fabric, collectionName, offset, limit, order)
	response = kv_req.NewResponseForGetAllKeys()
	err = c.request(req, response)
	return response, CheckError(err)
}

func (c Client) GetAllValues(fabric, collectionName string, offset, limit int, keys kv_req.KeyCollection) (response *kv_req.ResponseForGetAllValues, err error) {
	if benchmark {
		defer TimeTrack(time.Now(), "GetAllKeys")
	}

	req := kv_req.NewRequestForGetAllValues(fabric, collectionName, keys, offset, limit)
	response = kv_req.NewResponseForGetAllValues()
	err = c.request(req, response)
	return response, CheckError(err)
}

func (c Client) DeleteKeyValuePairs(fabric, collectionName string, keys kv_req.KeyCollection) (response *kv_req.KVPairCollection, err error) {
	if benchmark {
		defer TimeTrack(time.Now(), "SetKeyValuePairs")
	}

	req := kv_req.NewRequestForDeleteKeyValue(fabric, collectionName, keys)
	response = kv_req.NewKVPairCollection()
	err = c.request(req, response)
	return response, CheckError(err)
}

func (c Client) GetValue(fabric, collectionName, key string) (response *kv_req.KVPair, err error) {
	if benchmark {
		defer TimeTrack(time.Now(), "GetValue")
	}

	req := kv_req.NewRequestForGetValue(fabric, collectionName, key)
	response = kv_req.NewEmptyKVPair()
	err = c.request(req, response)
	return response, CheckError(err)
}

func (c Client) DeleteValue(fabric, collectionName, key string) (response *kv_req.KVPair, err error) {
	if benchmark {
		defer TimeTrack(time.Now(), "DeleteValue")
	}

	req := kv_req.NewRequestForDeleteValue(fabric, collectionName, key)
	response = kv_req.NewEmptyKVPair()
	err = c.request(req, response)
	return response, CheckError(err)
}
