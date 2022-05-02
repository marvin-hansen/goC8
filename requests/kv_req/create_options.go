package kv_req

func GetDefaultCreateKVOptions() *CreateKVOptions {
	return &CreateKVOptions{
		Stream:       false,
		EnableShards: false,
		WaitForSync:  false,
		ShardKeys:    []string{""},
	}

}

type CreateKVOptions struct {
	Stream       bool     `json:"stream"`
	EnableShards bool     `json:"enableShards"`
	WaitForSync  bool     `json:"waitForSync"`
	ShardKeys    []string `json:"shardKeys"`
}
