package kv_req

import (
	"encoding/json"
	"fmt"
)

type KVPairCollection []KVPair

func (r KVPairCollection) Json() []byte {
	jsonArray, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	return jsonArray
}

func NewKVPair(key, value string, expiration int) *KVPair {
	return &KVPair{
		Key:      key,
		Value:    value,
		ExpireAt: expiration,
	}
}

type KVPair struct {
	Id       string `json:"_id,omitempty"`
	Key      string `json:"_key,omitempty"`
	Value    string `json:"value,omitempty"`
	ExpireAt int    `json:"expireAt,omitempty"`
	Revision string `json:"_rev,omitempty"`
}

func (r KVPair) String() string {
	return fmt.Sprintf("Key: %v\n Value: %v\n Expiration: %v\n",
		r.Key,
		r.Value,
		r.ExpireAt,
	)
}

func (r KVPair) Json() []byte {
	bytes, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	return bytes
}
