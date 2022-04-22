package query_req

import (
	"fmt"
	"strings"
)

func NewCursorFromCreateCursor(r *ResponseForCreateCursor) *Cursor {
	return &Cursor{
		Code:    r.Code,
		Error:   r.Error,
		Count:   r.Count,
		Extra:   r.Extra,
		HasMore: r.HasMore,
		Id:      r.Id,
		Cached:  r.Cached,
		Result:  r.Result,
	}
}

type Cursor struct {
	Code    int    `json:"code"`
	Error   bool   `json:"error"`
	Count   int    `json:"count"`
	Extra   Extra  `json:"extra"`
	HasMore bool   `json:"hasMore"`
	Id      string `json:"id"`
	Cached  bool   `json:"cached"`
	Result  Result `json:"result"`
}

func (c Cursor) Update(r *ResponseForReadNextCursor) {
	c.Code = r.Code
	c.Error = r.Error
	c.Count = r.Count
	c.Extra = r.Extra
	c.HasMore = r.HasMore
	c.Id = r.Id
	c.Cached = r.Cached
	// append results
	c.Result = append(c.Result, r.Result)
}

func (c Cursor) String() string {
	return fmt.Sprintf("Code: %v\n Error: %v\n Count: %v\n Extra: %v\n HasMore: %v\n Id: %v\n Cached: %v\n Result: %v\n",
		c.Code,
		c.Error,
		c.Count,
		c.Extra,
		c.HasMore,
		c.Id,
		c.Cached,
		c.Result,
	)
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

type Result []interface{}

func (r Result) String() string {
	// https://stackoverflow.com/questions/62055988/golang-a-map-interface-how-to-print-key-and-value
	// https://pkg.go.dev/strings#Builder
	var s strings.Builder
	for k, v := range r {
		s.WriteString(fmt.Sprintf("%v: %v \n", k, v))
	}
	return s.String()
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
	return fmt.Sprintf(" WritesExecuted: %v\n WritesIgnored: %v\n ScannedFull: %v\n ScannedIndex: %v\n Filtered: %v\n HttpRequests: %v\n ExecutionTime: %v\n PeakMemoryUsage: %v",
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

	if len(r) == 0 {
		return "[]"
	} else {
		var s strings.Builder
		for k, v := range r {
			s.WriteString(fmt.Sprintf("%v: %v \n", k, v))
		}
		return s.String()
	}
}
