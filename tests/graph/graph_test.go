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
	verbose              = true
	silent               = true
	fabric               = "SouthEastAsia"
	graphName            = "lectureteacher"
	collectionTeachers   = "teachers"
	collectionLectures   = "lectures"
	collectionTutorials  = "tutorials"
	edgeCollectionTeach  = "teach"
	edgeCollectionTutors = "tutors"
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
		jsonPayload := []byte(`{
            "_key": "Bruce-CSC105",
            "_from": "teachers/Bruce",
            "_to": "lectures/CSC105",
            "online": false
        }`)
		_, createErr := c.Graph.CreateEdge(fabric, graphName, collectionID, "", "", jsonPayload, returnNew)
		assert.NoError(t, createErr)
	}
}

func TestAddEdgeCollection(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())

	// check if collection already exists
	exists, err := c.Collection.CheckCollectionExists(fabric, collectionTutorials)
	goC8.CheckError(err, "Error CheckCollectionExists: ")
	if !exists {
		// 1. if not create collection
		collType := types.DocumentCollectionType
		allowUserKeys := false
		err = c.Collection.CreateNewCollection(fabric, collectionTutorials, allowUserKeys, collType)
		assert.NoError(t, err)

		// 2. import data
		jsonDocument := sample_data.GetTutorialsData()
		_, err = c.Document.CreateNewDocument(fabric, collectionTutorials, silent, jsonDocument, nil)
		assert.NoError(t, err)

		// 3. Add vertex collection to graph
		_, err = c.Graph.AddVertexCollection(fabric, graphName, collectionTutorials)
		assert.NoError(t, err)
	}

	// check if collection already exists
	exists, err = c.Collection.CheckCollectionExists(fabric, edgeCollectionTutors)
	goC8.CheckError(err, "Error CheckCollectionExists")
	if !exists {
		// 1. if not create edge collection
		collType := types.EdgeCollectionType
		err = c.Collection.CreateNewCollection(fabric, edgeCollectionTutors, false, collType)
		assert.NoError(t, err)

		// 2. import data
		jsonDocument := sample_data.GetTutorsEdgeData()
		_, err = c.Document.CreateNewDocument(fabric, edgeCollectionTutors, silent, jsonDocument, nil)
		assert.NoError(t, err)

		// 3. add edge collection to graph
		collectionName := edgeCollectionTutors
		sourceVertex := "teachers"
		destinationVertex := "tutorials"

		res, createErr := c.Graph.AddEdgeCollection(fabric, graphName, collectionName, sourceVertex, destinationVertex)
		assert.NoError(t, createErr)
		assert.NotNil(t, res)
		goC8.PrintRes(res, verbose)
	}
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
	edgeID := "Jean-CSC102"
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

func TestUpdateEdge(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	collectionID := "teach"
	edgeID := "Jean-CSC101"
	// Note, only the fields to update must be present in the JSON message.

	updateJSON := []byte(`{
  "online": true
}`)

	// check if edge exits
	exists, err := c.Graph.CheckEdgeExists(fabric, graphName, collectionID, edgeID)
	goC8.CheckError(err, "Error CheckEdgeExists")
	if exists {

		// check if online is false, which it should
		res, err := c.Graph.GetEdge(fabric, graphName, collectionID, edgeID)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res.Edge.Online, false, "Should be false")

		// if exists, update edge by setting online to true
		res, updateErr := c.Graph.UpdateEdge(fabric, graphName, collectionID, edgeID, updateJSON, true, false, true)
		assert.NoError(t, updateErr)
		assert.NotNil(t, res)
		expected := true
		actual := res.New.Online
		assert.Equal(t, expected, actual)
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

func TestAddVertexCollection(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	collectionID := "classrooms"
	res, err := c.Graph.AddVertexCollection(fabric, graphName, collectionID)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	goC8.PrintRes(res, verbose)
}

func TestDeleteEdge(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	collectionID := "teach"
	edgeID := "Jean-CSC101"
	res, err := c.Graph.DeleteEdge(fabric, graphName, collectionID, edgeID, false)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	goC8.PrintRes(res, verbose)
}

func TestDeleteEdgeCollection(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	collectionID := "teach"

	res, err := c.Graph.DeleteEdgeCollection(fabric, graphName, collectionID, false)
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
	goC8.TeardownCollection(c, fabric, collectionTutorials)
	goC8.TeardownCollection(c, fabric, edgeCollectionTeach)
	goC8.TeardownCollection(c, fabric, edgeCollectionTutors)
}
