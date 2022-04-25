package graph

import (
	"github.com/marvin-hansen/goC8"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAllVertices(t *testing.T) {
	c := goC8.NewClient(nil)

	res, err := c.GetAllVertices(fabric, graphName)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	goC8.PrintRes(res, verbose)
}

func TestGetVertex(t *testing.T) {
	c := goC8.NewClient(nil)
	collectionID := "teachers"
	edgeID := "Jean"

	res, err := c.GetVertex(fabric, graphName, collectionID, edgeID)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	goC8.PrintRes(res, verbose)
}