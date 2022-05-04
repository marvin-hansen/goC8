package graph

import (
	"github.com/marvin-hansen/goC8"
	config "github.com/marvin-hansen/goC8/tests/conf"
	"github.com/marvin-hansen/goC8/types"
	utils2 "github.com/marvin-hansen/goC8/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	verbose             = true
	fabric              = "SouthEastAsia"
	graphName           = "lectureteacher"
	collectionTeachers  = "teachers"
	collectionLectures  = "lectures"
	edgeCollectionTeach = "teach"
)

func TestSetup(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())

	goC8.CreateCollection(c, fabric, collectionTeachers, types.DocumentCollectionType, false)
	//goC8.ImportData(c, fabric, collectionTeachers, GetTeacherData(), silent)
	goC8.CreateCollection(c, fabric, collectionLectures, types.DocumentCollectionType, false)
	//goC8.ImportData(c, fabric, collectionLectures, GetLecturesData(), silent)
	goC8.CreateCollection(c, fabric, edgeCollectionTeach, types.EdgeCollectionType, false)
	// goC8.ImportData(c, fabric, edgeCollectionTeach, GetLecturesData(), silent)
}

func TestCreateGraph(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	graphDef := getGraphDefinition()
	res, err := c.Graph.CreateGraph(fabric, graphDef)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	utils2.PrintRes(res, verbose)
}

func TestGetAllGraphs(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())

	res, err := c.Graph.GetAllGraphs(fabric)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	utils2.PrintRes(res, verbose)
}

func TestGetGraph(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())

	res, err := c.Graph.GetGraph(fabric, graphName)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	utils2.PrintRes(res, verbose)
}

func TestCheckGraphExists(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())

	exists, err := c.Graph.CheckGraphExists(fabric, graphName)
	assert.NoError(t, err)
	expected := true
	actual := exists
	assert.Equal(t, expected, actual, "Should exists")
	noneExistingGraphName := "noneExistingGraphName"
	exists, err = c.Graph.CheckGraphExists(fabric, noneExistingGraphName)
	assert.NoError(t, err)

	expected = false
	actual = exists
	assert.Equal(t, expected, actual, "Should exists")
}

func TestGetAllEdges(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())

	res, err := c.Graph.GetAllEdges(fabric, graphName)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	utils2.PrintRes(res, verbose)
}

func TestGetEdge(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	collectionID := "teach"
	edgeID := "Jean-CSC101"

	res, err := c.Graph.GetEdge(fabric, graphName, collectionID, edgeID)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	utils2.PrintRes(res, verbose)
}

func TestCheckEdgeExists(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	collectionID := "teach"
	edgeID := "Jean-CSC101"

	res, err := c.Graph.CheckEdgeExists(fabric, graphName, collectionID, edgeID)
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
	exists, err := c.Graph.CheckEdgeExists(fabric, graphName, collectionID, edgeID)
	utils2.CheckError(err, "Error CheckEdgeExists")
	if !exists {
		// if not, add a new edge to the edge collection
		from := "teachers/Bruce"
		to := "lectures/CSC105"
		_, createErr := c.Graph.CreateEdge(fabric, graphName, collectionID, from, to, returnNew)
		utils2.CheckError(createErr, "Error CreateEdge")
	}
}

func TestReplaceEdge(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	collectionID := "teach"
	edgeID := "Bruce-CSC105"
	dropCollections := false

	// check if edge exits
	exists, err := c.Graph.CheckEdgeExists(fabric, graphName, collectionID, edgeID)
	utils2.CheckError(err, "Error CheckEdgeExists")
	if exists {
		// if exists, replace edge with a new edge to the edge collection
		from := "teachers/Bruce"
		to := "lectures/CSC105"
		_, createErr := c.Graph.ReplaceEdge(fabric, graphName, collectionID, from, to, dropCollections)
		utils2.CheckError(createErr, "Error CreateEdge")
	}
}

func TestGetAllVertices(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())

	res, err := c.Graph.GetAllVertices(fabric, graphName)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	utils2.PrintRes(res, verbose)
}

func TestGetVertex(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	collectionID := "teachers"
	edgeID := "Jean"

	res, err := c.Graph.GetVertex(fabric, graphName, collectionID, edgeID)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	utils2.PrintRes(res, verbose)
}

func TestDeleteGraph(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	dropCollections := false

	res, err := c.Graph.DeleteGraph(fabric, graphName, dropCollections)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	utils2.PrintRes(res, verbose)
}

func TestTeardown(t *testing.T) {

}
