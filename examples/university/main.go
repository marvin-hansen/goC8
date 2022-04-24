package main

import (
	"github.com/marvin-hansen/goC8"
)

// client config
const (
	apiKey   = "email.root.secretkey"
	endpoint = "https://YOUR-ID-us-west.paas.macrometa.io"
	fabric   = "SouthEastAsia"
	timeout  = 5 // http connection timeout in seconds
)

const (
	delete               = true
	verbose              = true
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

	println("Create new client with config ")
	c := goC8.NewClient(config)

	println("Run setup")
	setup(c)

	if delete {
		println("Run teardown")
		teardown(c)
	}
}
