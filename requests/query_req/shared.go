package query_req

import (
	"fmt"
	"strings"
)

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

func (r *Cursor) IsResponse() {}

func (c Cursor) Update(r *Cursor) {
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

type Plan struct {
	Nodes       []Node        `json:"nodes,omitempty"`
	Rules       []interface{} `json:"rules,omitempty"`
	Collections []struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"collections,omitempty"`
	Variables []struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"variables,omitempty"`
	EstimatedCost       int  `json:"estimatedCost,omitempty"`
	EstimatedNrItems    int  `json:"estimatedNrItems,omitempty"`
	Initialize          bool `json:"initialize,omitempty"`
	IsModificationQuery bool `json:"isModificationQuery,omitempty"`
}

func (s Plan) String() string {
	return fmt.Sprintf("\n Nodes: %v\n Rules: %v\n Collections: %v\n Variables: %v\n EstimatedCost: %v\n EstimatedNrItems: %v\n Initialize: %v\n IsModificationQuery: %v",
		s.Nodes,
		s.Rules,
		s.Collections,
		s.Variables,
		s.EstimatedCost,
		s.EstimatedNrItems,
		s.Initialize,
		s.IsModificationQuery,
	)
}

type Node struct {
	Type             string        `json:"type"`
	Dependencies     []int         `json:"dependencies"`
	Id               int           `json:"id"`
	EstimatedCost    int           `json:"estimatedCost"`
	EstimatedNrItems int           `json:"estimatedNrItems"`
	Random           bool          `json:"random,omitempty"`
	IndexHint        IndexHint     `json:"indexHint,omitempty"`
	OutVariable      Variable      `json:"outVariable,omitempty"`
	Projections      []interface{} `json:"projections,omitempty"`
	ProducesResult   bool          `json:"producesResult,omitempty"`
	Database         string        `json:"database,omitempty"`
	Collection       string        `json:"collection,omitempty"`
	Satellite        bool          `json:"satellite,omitempty"`
	InVariable       Variable      `json:"inVariable,omitempty"`
	Count            bool          `json:"count,omitempty"`
}

func (s Node) String() string {
	return fmt.Sprintf(" \n Node: \n Type: %v\n Dependencies: %v\n Id: %v\n EstimatedCost: %v\n EstimatedNrItems: %v\n Random: %v\n IndexHint: %v\n OutVariable: %v\n Projections: %v\n ProducesResult: %v\n Database: %v\n Collection: %v\n Satellite: %v\n InVariable: %v\n Count: %v\n",
		s.Type,
		s.Dependencies,
		s.Id,
		s.EstimatedCost,
		s.EstimatedNrItems,
		s.Random,
		s.IndexHint,
		s.OutVariable,
		s.Projections,
		s.ProducesResult,
		s.Database,
		s.Collection,
		s.Satellite,
		s.InVariable,
		s.Count,
	)
}

type ExplainStats struct {
	RulesExecuted int `json:"rulesExecuted"`
	RulesSkipped  int `json:"rulesSkipped"`
	PlansCreated  int `json:"plansCreated"`
}

func (s ExplainStats) String() string {
	return fmt.Sprintf("\n RulesExecuted: %v\n RulesSkipped: %v\n PlansCreated: %v\n",
		s.RulesExecuted,
		s.RulesSkipped,
		s.PlansCreated,
	)
}

type IndexHint struct {
	Forced bool   `json:"forced"`
	Type   string `json:"type"`
}

func (s IndexHint) String() string {
	return fmt.Sprintf("Forced: %v Type: %v",
		s.Forced,
		s.Type,
	)
}

type Variable struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (s Variable) String() string {
	return fmt.Sprintf("ID: %v Name: %v",
		s.Id,
		s.Name,
	)
}
