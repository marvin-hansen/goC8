package graph

import (
	"github.com/marvin-hansen/goC8"
	"github.com/marvin-hansen/goC8/src/utils"
	config "github.com/marvin-hansen/goC8/tests/conf"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAllEdges(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())

	res, err := c.GetAllEdges(fabric, graphName)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	utils.PrintRes(res, verbose)
}

func TestGetEdge(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	collectionID := "teach"
	edgeID := "Jean-CSC101"

	res, err := c.GetEdge(fabric, graphName, collectionID, edgeID)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	utils.PrintRes(res, verbose)
}

func TestCheckEdgeExists(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
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
	c := goC8.NewClient(config.GetDefaultConfig())
	collectionID := "teach"
	edgeID := "Bruce-CSC105"
	returnNew := false

	// check if edge already exits
	exists, err := c.CheckEdgeExists(fabric, graphName, collectionID, edgeID)
	utils.CheckError(err, "Error CheckEdgeExists")
	if !exists {
		// if not, add a new edge to the edge collection
		from := "teachers/Bruce"
		to := "lectures/CSC105"
		_, createErr := c.CreateEdge(fabric, graphName, collectionID, from, to, returnNew)
		utils.CheckError(createErr, "Error CreateEdge")
	}
}

func TestReplaceEdge(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	collectionID := "teach"
	edgeID := "Bruce-CSC105"
	dropCollections := false

	// check if edge exits
	exists, err := c.CheckEdgeExists(fabric, graphName, collectionID, edgeID)
	utils.CheckError(err, "Error CheckEdgeExists")
	if exists {
		// if exists, replace edge with a new edge to the edge collection
		from := "teachers/Bruce"
		to := "lectures/CSC105"
		_, createErr := c.ReplaceEdge(fabric, graphName, collectionID, from, to, dropCollections)
		utils.CheckError(createErr, "Error CreateEdge")
	}
}
