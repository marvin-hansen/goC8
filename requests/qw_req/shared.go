package qw_req

import "fmt"

func getQueryWorkerPayload(workerName, queryString, parameterString string) []byte {
	str := fmt.Sprintf(`{
  "query": {
    "name": "%v",
    "parameter": {%v},
    "value": "%v"
  }
}`, workerName, parameterString, queryString)
	return []byte(str)
}

type QueryWorker struct {
	Name      string   `json:"name,omitempty"`
	Key       string   `json:"_key,omitempty"`
	Rev       string   `json:"_rev,omitempty"`
	Userid    string   `json:"userid,omitempty"`
	Tenant    string   `json:"tenant,omitempty"`
	Fabric    string   `json:"fabric,omitempty"`
	Value     string   `json:"value,omitempty"`
	Parameter struct{} `json:"parameter,omitempty"`
	Type      string   `json:"type,omitempty"`
	Data      string   `json:"data,omitempty"`
}

func (r QueryWorker) String() string {
	return fmt.Sprintf("Name: %v\n Key: %v\n Rev: %v\n Userid: %v\n Tenant: %v\n Fabric: %v\n Value: %v\n Parameter: %v\n Type: %v\n Data: %v\n",
		r.Name,
		r.Key,
		r.Rev,
		r.Userid,
		r.Tenant,
		r.Fabric,
		r.Value,
		r.Parameter,
		r.Type,
		r.Data,
	)
}
