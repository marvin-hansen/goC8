package query

import (
	"github.com/marvin-hansen/goC8"
	"github.com/marvin-hansen/goC8/examples/sample_data"
	"github.com/marvin-hansen/goC8/tests/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	verbose = true
	fabric  = "SouthEastAsia"
)

func TestGetIndexes(t *testing.T) {
	c := goC8.NewClient(nil)
	query := sample_data.GetBreadthFirstQuery()
	res, err := c.Query(fabric, query, nil, nil)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	utils.PrintQuery(res, verbose)
}
