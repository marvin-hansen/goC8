package kv_req

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func NewKVPairCollection(pairs ...KVPair) *KVPairCollection {
	kvCol := KVPairCollection{}
	for _, v := range pairs {
		kvCol = append(kvCol, v)
	}
	return &kvCol
}

func NewEmptyKVPairCollection() *KVPairCollection {
	return new(KVPairCollection)
}

type KVPairCollection []KVPair

func (r *KVPairCollection) IsResponse() {}

func (r KVPairCollection) String() string {
	var s bytes.Buffer
	for _, v := range r {
		s.WriteString(v.String())
		s.WriteString(" \n ")
	}
	return s.String()
}

func (r KVPairCollection) Json() []byte {
	jsonArray, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	return jsonArray
}

func NewEmptyKVPair() *KVPair {
	return new(KVPair)
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
	Revision string `json:"_rev,omitempty"`
	ExpireAt int    `json:"expireAt,omitempty"`
}

func (r *KVPair) IsResponse() {}

func (r KVPair) String() string {
	return fmt.Sprintf(" ID: %v\n Key: %v\n Value: %v\n Revision: %v\n Expiration: %v\n",
		r.Id,
		r.Key,
		r.Value,
		r.Revision,
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
