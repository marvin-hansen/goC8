package kv_req

import "fmt"

func NewKVResponse() *KVResult {
	return new(KVResult)
}

type KVResult struct {
	Code   int      `json:"code,omitempty"`
	Error  bool     `json:"error,omitempty"`
	Name   string   `json:"name,omitempty"`
	Count  int      `json:"count,omitempty"`
	Result []Result `json:"result,omitempty"`
}

type Result struct {
	Name       string `json:"name,omitempty"`
	Expiration bool   `json:"expiration,omitempty"`
}

func (r Result) String() string {
	return fmt.Sprintf("Name: %v Expiration: %v",
		r.Name,
		r.Expiration,
	)
}

func (r *KVResult) IsResponse() {}

func (r KVResult) String() string {
	return fmt.Sprintf(" Code: %v\n Error: %v\n Name: %v\n Count: %v\n Result: %v\n  ",
		r.Code,
		r.Error,
		r.Name,
		r.Count,
		r.Result,
	)
}
