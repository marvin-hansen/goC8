package query

import (
	"github.com/marvin-hansen/goC8"
	"github.com/marvin-hansen/goC8/examples/sample_data"
	config "github.com/marvin-hansen/goC8/tests/conf"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	verbose = true
	fabric  = "SouthEastAsia"
)

func TestGetIndexes(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	query := sample_data.GetBreadthFirstQuery()
	res, err := c.Query(fabric, query, nil, nil)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	goC8.PrintQuery(res, verbose)
}
