package main

import (
	"github.com/marvin-hansen/goC8"
	"github.com/marvin-hansen/goC8/examples/sample_data"
	"github.com/marvin-hansen/goC8/types"
	"github.com/marvin-hansen/goC8/utils"
)

// client config
const (
	apiKey   = "email.root.secretkey"
	endpoint = "https://YOUR-ID-us-west.paas.macrometa.io"
	fabric   = "MyFabric"
	timeout  = 5 // http connection timeout in seconds
)

const (
	delete               = false
	silent               = false
	collectionTeachers   = "teachers"
	collectionLectures   = "lectures"
	collectionTutorials  = "tutorials"
	edgeCollectionTeach  = "teach"
	edgeCollectionTutors = "tutors"
	graph                = "lectureteacher"
)

func main() {
	println("Create new config ")
	config := goC8.NewConfig(apiKey, endpoint, fabric, timeout)

	println("Create new client")
	c := goC8.NewClient(config)

	println("Run setup")
	setup(c)

	println("Update graph")
	addTutorials(c)
	addTutorEdge(c)
	addEdgeCollection(c)
	addSingleEdge(c)

	if delete {
		println("Run teardown")
		teardown(c)
	}
}

// setup uses the built-in / low-code utilities to create a collection & import data.
// These utils were written to accelerate developer velocity on standard API operations
func setup(c *goC8.Client) {
	goC8.CreateCollection(c, fabric, collectionTeachers, types.DocumentCollectionType, false)
	goC8.ImportCollectionData(c, fabric, collectionTeachers, sample_data.GetTeacherData(), silent)
	goC8.CreateCollection(c, fabric, collectionLectures, types.DocumentCollectionType, false)
	goC8.ImportCollectionData(c, fabric, collectionLectures, sample_data.GetLecturesData(), silent)
	goC8.CreateCollection(c, fabric, edgeCollectionTeach, types.EdgeCollectionType, false)
	goC8.ImportCollectionData(c, fabric, edgeCollectionTeach, sample_data.GetTeachEdgeData(), silent)
	goC8.CreateGraph(c, fabric, graph, sample_data.GetUniversityGraphDefinition())
}

// addTutorials and subsequent update methods use the standard API to create a collection & import data
// Note the update function can be re-implemented with the utils used in setup.
// However, in non-standard cases direction API usage may be preferred for customization
func addTutorials(c *goC8.Client) {
	exists, err := c.Collection.CheckCollectionExists(fabric, collectionTutorials)
	utils.CheckError(err, "Error CheckCollectionExists: ")
	if !exists {
		// if not create collection
		collType := types.DocumentCollectionType
		allowUserKeys := false
		err = c.Collection.CreateNewCollection(fabric, collectionTutorials, allowUserKeys, collType)
		utils.CheckError(err, "Error CreateNewCollection")

		// import data
		jsonDocument := sample_data.GetTutorialsData()
		_, err = c.Document.CreateNewDocument(fabric, collectionTutorials, silent, jsonDocument, nil)
		utils.CheckError(err, "Error CreateNewDocument")
	}
}

func addTutorEdge(c *goC8.Client) {
	// test if flight collection exists
	exists, err := c.Collection.CheckCollectionExists(fabric, edgeCollectionTutors)
	utils.CheckError(err, "Error CheckCollectionExists")

	if !exists {
		// if not create edge collection
		collType := types.EdgeCollectionType
		err = c.Collection.CreateNewCollection(fabric, edgeCollectionTutors, false, collType)
		utils.CheckError(err, "Error CreateNewCollection")

		// import data
		jsonDocument := sample_data.GetTutorsEdgeData()
		_, err = c.Document.CreateNewDocument(fabric, edgeCollectionTutors, silent, jsonDocument, nil)
		utils.CheckError(err, "Error CreateNewCollection")
	}
}

func addEdgeCollection(c *goC8.Client) {
	// check if collection exists. It should, just in case
	exists, err := c.Collection.CheckCollectionExists(fabric, edgeCollectionTutors)
	utils.CheckError(err, "Error CheckCollectionExists")
	if exists {
		// if exists, update graph with new edge collection
		collectionName := edgeCollectionTutors
		sourceVertex := "teachers"
		destinationVertex := "tutorials"

		_, addErr := c.Graph.AddEdgeCollection(fabric, graph, collectionName, sourceVertex, destinationVertex)
		if addErr != nil {
			return
		}
	} else {
		println("Can't update graph. Add edge collection: " + edgeCollectionTutors)
	}
}

func addSingleEdge(c *goC8.Client) {
	collectionID := "teach"
	edgeID := "Bruce-CSC105"
	// check if edge exits
	exists, err := c.Graph.CheckEdgeExists(fabric, graph, collectionID, edgeID)
	utils.CheckError(err, "Error CheckEdgeExists")
	if !exists {
		// if not, add a new edge to the edge collection
		from := "teachers/Bruce"
		to := "lectures/CSC105"
		_, createErr := c.Graph.CreateEdge(fabric, graph, collectionID, from, to, false)
		utils.CheckError(createErr, "Error CreateEdge")
	}
}

func teardown(c *goC8.Client) {
	goC8.TeardownGraph(c, fabric, graph, true)
	goC8.TeardownCollection(c, fabric, collectionTeachers)
	goC8.TeardownCollection(c, fabric, collectionLectures)
	goC8.TeardownCollection(c, fabric, edgeCollectionTeach)
}
