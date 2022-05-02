package kv_req

import (
	"bytes"
	"encoding/json"
)

type KeyCollection []string

func (r *KeyCollection) IsResponse() {}

func (r KeyCollection) String() string {
	var s bytes.Buffer
	for _, v := range r {
		s.WriteString(v)
		s.WriteString(" \n ")
	}
	return s.String()
}

func (r KeyCollection) Json() []byte {
	b, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	return b
}
