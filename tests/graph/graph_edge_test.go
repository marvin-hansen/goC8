package graph

import (
	"github.com/marvin-hansen/goC8"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAllEdges(t *testing.T) {
	c := goC8.NewClient(nil)

	res, err := c.GetAllEdges(fabric, graphName)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	goC8.PrintRes(res, verbose)
}

func TestGetEdge(t *testing.T) {
	c := goC8.NewClient(nil)
	collectionID := "teach"
	edgeID := "Jean-CSC101"

	res, err := c.GetEdge(fabric, graphName, collectionID, edgeID)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	goC8.PrintRes(res, verbose)
}
