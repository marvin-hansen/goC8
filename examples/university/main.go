package main

import (
	"github.com/marvin-hansen/goC8"
	"github.com/marvin-hansen/goC8/requests/collection_req"
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
	goC8.CreateCollection(c, fabric, collectionTeachers, collection_req.DocumentCollectionType, false)
	goC8.ImportData(c, fabric, collectionTeachers, GetTeacherData(), silent)
	goC8.CreateCollection(c, fabric, collectionLectures, collection_req.DocumentCollectionType, false)
	goC8.ImportData(c, fabric, collectionLectures, GetLecturesData(), silent)
	goC8.CreateCollection(c, fabric, edgeCollectionTeach, collection_req.EdgeCollectionType, false)
	goC8.ImportData(c, fabric, edgeCollectionTeach, GetLecturesData(), silent)
	goC8.CreateGraph(c, fabric, graph, GetUniversityGraphDefinition())
}

func teardown(c *goC8.Client) {
	goC8.TeardownGraph(c, fabric, graph, true)
	goC8.TeardownCollection(c, fabric, collectionTeachers)
	goC8.TeardownCollection(c, fabric, collectionLectures)
	goC8.TeardownCollection(c, fabric, edgeCollectionTeach)
}
