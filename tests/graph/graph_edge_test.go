package graph

import (
	"github.com/marvin-hansen/goC8"
	"github.com/marvin-hansen/goC8/utils"
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

func TestCheckEdgeExists(t *testing.T) {
	c := goC8.NewClient(nil)
	collectionID := "teach"
	edgeID := "Jean-CSC101"

	res, err := c.CheckEdgeExists(fabric, graphName, collectionID, edgeID)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	expected := true
	actual := res
	assert.Equal(t, expected, actual)
}

func TestAddEdge(t *testing.T) {
	c := goC8.NewClient(nil)
	collectionID := "teach"
	edgeID := "Bruce-CSC105"

	// check if edge already exits
	exists, err := c.CheckEdgeExists(fabric, graphName, collectionID, edgeID)
	utils.CheckError(err, "Error CheckEdgeExists")
	if !exists {
		// if not, add a new edge to the edge collection
		from := "teachers/Bruce"
		to := "lectures/CSC105"
		_, createErr := c.CreateEdge(fabric, graphName, collectionID, from, to, false)
		utils.CheckError(createErr, "Error CreateEdge")
	}
}
