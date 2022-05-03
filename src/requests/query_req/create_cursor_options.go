package query_req

func NewDefaultCursorOptions() *CursorOptions {
	return &CursorOptions{
		FailOnWarning:               true,
		FullCount:                   true,
		IntermediateCommitCount:     0,
		IntermediateCommitSize:      0,
		MaxTransactionSize:          0,
		MaxWarningCount:             10,
		Profile:                     0,
		SkipInaccessibleCollections: true,
		Stream:                      false,
		OptimizerRules:              []string{"string"},
	}
}

func NewEmptyCursorOptions() *CursorOptions {
	return new(CursorOptions)
}

type CursorOptions struct {
	FailOnWarning               bool     `json:"failOnWarning,omitempty"`
	FullCount                   bool     `json:"fullCount,omitempty"`
	IntermediateCommitCount     int      `json:"intermediateCommitCount,omitempty"`
	IntermediateCommitSize      int      `json:"intermediateCommitSize,omitempty"`
	MaxTransactionSize          int      `json:"maxTransactionSize,omitempty"`
	MaxWarningCount             int      `json:"maxWarningCount,omitempty"`
	Profile                     int      `json:"profile,omitempty"`
	SkipInaccessibleCollections bool     `json:"skipInaccessibleCollections,omitempty"`
	Stream                      bool     `json:"stream,omitempty"`
	OptimizerRules              []string `json:"optimizer.rules,omitempty"`
}
