package main

import (
	"github.com/marvin-hansen/goC8"
	"github.com/marvin-hansen/goC8/examples/sample_data"
	"github.com/marvin-hansen/goC8/requests/collection_req"
	"log"
)

const (
	verbose              = true
	fabric               = "SouthEastAsia"
	collectionTeachers   = "teachers"
	collectionLectures   = "lectures"
	collectionTutorials  = "tutorials"
	edgeCollectionTeach  = "teach"
	edgeCollectionTutors = "tutors"
	graph                = "lectureteacher"
)

func main() {
	c := goC8.NewClient(nil)
	setup(c)

}

func setup(c *goC8.Client) {
	setupCourses(c)
	setupTeachers(c)
	setupGraph(c)
}

func setupTeachers(c *goC8.Client) {
	exists, err := c.CheckCollectionExists(fabric, collectionTeachers)
	checkError(err, "Error CheckCollectionExists: ")
	if !exists {
		// if not create collection
		collType := collection_req.DocumentCollectionType
		allowUserKeys := false
		err = c.CreateNewCollection(fabric, collectionTeachers, allowUserKeys, collType)
		checkError(err, "Error CreateNewCollection")

		// import city data
		silent := false
		jsonDocument := sample_data.GetTeacherData()
		_, err = c.CreateNewDocument(fabric, collectionTeachers, silent, jsonDocument, nil)
		checkError(err, "Error CreateNewDocument")
	}
}

func setupCourses(c *goC8.Client) {
	exists, err := c.CheckCollectionExists(fabric, collectionLectures)
	checkError(err, "Error CheckCollectionExists: ")

	if !exists {
		// if not create collection
		collType := collection_req.DocumentCollectionType
		allowUserKeys := false
		err = c.CreateNewCollection(fabric, collectionLectures, allowUserKeys, collType)
		checkError(err, "Error CreateNewCollection")

		// import  data
		silent := false
		jsonDocument := sample_data.GetLecturesData()
		_, err = c.CreateNewDocument(fabric, collectionLectures, silent, jsonDocument, nil)
		checkError(err, "Error CreateNewDocument")
	}

}

func setupGraph(c *goC8.Client) {

}

func teardown(c *goC8.Client) {

}

func checkError(err error, msg string) {
	if err != nil {
		log.Println("error: " + err.Error())
		log.Fatalf(msg)
	}
}
