package graph

import (
	"github.com/marvin-hansen/goC8"
	"github.com/marvin-hansen/goC8/examples/sample_data"
	config "github.com/marvin-hansen/goC8/tests/conf"
	"github.com/marvin-hansen/goC8/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	verbose             = true
	silent              = true
	fabric              = "SouthEastAsia"
	graphName           = "lectureteacher"
	collectionTeachers  = "teachers"
	collectionLectures  = "lectures"
	edgeCollectionTeach = "teach"
)

func TestSetup(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	goC8.CreateCollection(c, fabric, collectionTeachers, types.DocumentCollectionType, false)
	goC8.ImportCollectionData(c, fabric, collectionTeachers, sample_data.GetTeacherData(), silent)
	goC8.CreateCollection(c, fabric, collectionLectures, types.DocumentCollectionType, false)
	goC8.ImportCollectionData(c, fabric, collectionLectures, sample_data.GetLecturesData(), silent)
	goC8.CreateCollection(c, fabric, edgeCollectionTeach, types.EdgeCollectionType, false)
	goC8.ImportCollectionData(c, fabric, edgeCollectionTeach, sample_data.GetTeachEdgeData(), silent)
}

func TestCreateGraph(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	graphDef := getGraphDefinition()
	res, err := c.Graph.CreateGraph(fabric, graphDef)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	goC8.PrintRes(res, verbose)
}

func TestGetAllGraphs(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	res, err := c.Graph.GetAllGraphs(fabric)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	goC8.PrintRes(res, verbose)
}

func TestGetGraph(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	res, err := c.Graph.GetGraph(fabric, graphName)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	goC8.PrintRes(res, verbose)
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
	goC8.PrintRes(res, verbose)
}

func TestGetEdge(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	collectionID := "teach"
	edgeID := "Jean-CSC101"
	res, err := c.Graph.GetEdge(fabric, graphName, collectionID, edgeID)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	goC8.PrintRes(res, verbose)
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
	goC8.CheckError(err, "Error CheckEdgeExists")
	if !exists {
		// if not, add a new edge to the edge collection
		from := "teachers/Bruce"
		to := "lectures/CSC105"
		_, createErr := c.Graph.CreateEdge(fabric, graphName, collectionID, from, to, returnNew)
		goC8.CheckError(createErr, "Error CreateEdge")
	}
}

func TestReplaceEdge(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	collectionID := "teach"
	edgeID := "Bruce-CSC105"
	dropCollections := false

	// check if edge exits
	exists, err := c.Graph.CheckEdgeExists(fabric, graphName, collectionID, edgeID)
	goC8.CheckError(err, "Error CheckEdgeExists")
	if exists {
		// if exists, replace edge with a new edge to the edge collection
		from := "teachers/Bruce"
		to := "lectures/CSC105"
		_, createErr := c.Graph.ReplaceEdge(fabric, graphName, collectionID, from, to, dropCollections)
		goC8.CheckError(createErr, "Error CreateEdge")
	}
}

func TestGetAllVertices(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	res, err := c.Graph.GetAllVertices(fabric, graphName)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	goC8.PrintRes(res, verbose)
}

func TestGetVertex(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	collectionID := "teachers"
	edgeID := "Jean"
	res, err := c.Graph.GetVertex(fabric, graphName, collectionID, edgeID)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	goC8.PrintRes(res, verbose)
}

func TestDeleteGraph(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	dropCollections := true
	res, err := c.Graph.DeleteGraph(fabric, graphName, dropCollections)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	goC8.PrintRes(res, verbose)
}

func TestTeardown(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	goC8.TeardownGraph(c, fabric, graphName, true)
	goC8.TeardownCollection(c, fabric, collectionTeachers)
	goC8.TeardownCollection(c, fabric, collectionLectures)
	goC8.TeardownCollection(c, fabric, edgeCollectionTeach)
}
