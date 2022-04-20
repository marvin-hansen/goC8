package graph

import (
	"github.com/marvin-hansen/goC8"
	"github.com/stretchr/testify/assert"
	"testing"
)

const verbose = true

func TestGetAllGraphs(t *testing.T) {
	c := goC8.NewClient(nil)
	fabric := "_system"

	res, err := c.GetAllGraphs(fabric)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	printRes(res, verbose)
}

func TestGetGraph(t *testing.T) {
	c := goC8.NewClient(nil)
	fabric := "_system"
	graphName := "lectureteacher"

	res, err := c.GetGraph(fabric, graphName)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	printRes(res, verbose)
}

func TestGetAllEdges(t *testing.T) {
	c := goC8.NewClient(nil)
	fabric := "_system"
	graphName := "lectureteacher"

	res, err := c.GetAllEdges(fabric, graphName)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	printRes(res, verbose)
}

func TestGetEdge(t *testing.T) {
	c := goC8.NewClient(nil)
	fabric := "_system"
	graphName := "lectureteacher"
	collectionID := "teach"
	edgeID := "Jean-CSC101"

	res, err := c.GetEdge(fabric, graphName, collectionID, edgeID)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	printRes(res, verbose)
}
