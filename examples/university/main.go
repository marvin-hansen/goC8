package main

import (
	"github.com/marvin-hansen/goC8"
	"github.com/marvin-hansen/goC8/examples/sample_data"
	"github.com/marvin-hansen/goC8/types"
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
	update(c)

	if delete {
		println("Run teardown")
		teardown(c)
	}
}

func setup(c *goC8.Client) {
	goC8.CreateCollection(c, fabric, collectionTeachers, types.DocumentCollectionType, false)
	goC8.ImportData(c, fabric, collectionTeachers, sample_data.GetTeacherData(), silent)
	goC8.CreateCollection(c, fabric, collectionLectures, types.DocumentCollectionType, false)
	goC8.ImportData(c, fabric, collectionLectures, sample_data.GetLecturesData(), silent)

	goC8.CreateCollection(c, fabric, edgeCollectionTeach, types.EdgeCollectionType, false)
	goC8.ImportData(c, fabric, edgeCollectionTeach, sample_data.GetTeachEdgeData(), silent)
	goC8.CreateGraph(c, fabric, graph, sample_data.GetUniversityGraphDefinition())
}

func teardown(c *goC8.Client) {
	goC8.TeardownGraph(c, fabric, graph, true)
	goC8.TeardownCollection(c, fabric, collectionTeachers)
	goC8.TeardownCollection(c, fabric, collectionLectures)
	goC8.TeardownCollection(c, fabric, edgeCollectionTeach)
}
