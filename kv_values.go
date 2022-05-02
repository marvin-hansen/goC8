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
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}
