package goC8

import (
	"github.com/marvin-hansen/goC8/requests/kv_req"
	"time"
)

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
