package index

import (
	"github.com/marvin-hansen/goC8"
	"github.com/stretchr/testify/assert"
	"testing"
)

const verbose = true

func TestGetIndexes(t *testing.T) {
	c := goC8.NewClient(nil)
	fabric := "_system"
	collectionID := "teachers"

	res, err := c.GetIndexes(fabric, collectionID)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	printRes(res, verbose)
}
