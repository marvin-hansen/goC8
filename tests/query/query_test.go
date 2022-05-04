package query

import (
	"github.com/marvin-hansen/goC8"
	"github.com/marvin-hansen/goC8/examples/flight"
	config "github.com/marvin-hansen/goC8/tests/conf"
	"github.com/marvin-hansen/goC8/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	verbose = true
	fabric  = "SouthEastAsia"
)

func TestGetIndexes(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	query := main.GetBreadthFirstQuery()
	res, err := c.Query(fabric, query, nil, nil)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	utils.PrintQuery(res, verbose)
}
