package query_req

import "fmt"

type CursorResponse struct {
	Code    int    `json:"code"`
	Error   bool   `json:"error"`
	Count   int    `json:"count"`
	Extra   Extra  `json:"extra"`
	HasMore bool   `json:"hasMore"`
	Id      string `json:"id"`
	Cached  bool   `json:"cached"`
	Result  Result `json:"result"`
}

type Extra struct {
	Stats    `json:"stats"`
	Warnings `json:"warnings"`
}

func (r Extra) String() string {
	return fmt.Sprintf("Stats: %v\n Warnings: %v\n",
		r.Stats,
		r.Warnings,
	)
}

type Result map[string]interface{}

func (r Result) String() string {
	return fmt.Sprintf("Result: %v", r.String())
}

type Stats struct {
	WritesExecuted  int     `json:"writesExecuted"`
	WritesIgnored   int     `json:"writesIgnored"`
	ScannedFull     int     `json:"scannedFull"`
	ScannedIndex    int     `json:"scannedIndex"`
	Filtered        int     `json:"filtered"`
	HttpRequests    int     `json:"httpRequests"`
	ExecutionTime   float64 `json:"executionTime"`
	PeakMemoryUsage int     `json:"peakMemoryUsage"`
}

func (r Stats) String() string {
	return fmt.Sprintf(" WritesExecuted: %v\n WritesIgnored: %v\n ScannedFull: %v\n ScannedIndex: %v\n Filtered: %v\n HttpRequests: %v\n ExecutionTime: %v\n PeakMemoryUsage: %v\n",
		r.WritesExecuted,
		r.WritesIgnored,
		r.ScannedFull,
		r.ScannedIndex,
		r.Filtered,
		r.HttpRequests,
		r.ExecutionTime,
		r.PeakMemoryUsage,
	)
}

type Warnings []interface{}

func (r Warnings) String() string {
	return fmt.Sprintf("Warnings: %v", r.String())
}
