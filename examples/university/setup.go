package main

import (
	"github.com/marvin-hansen/goC8"
	"github.com/marvin-hansen/goC8/requests/collection_req"
)

func setup(c *goC8.Client) {
	setupTeachers(c)

	//setupCourses(c)
	//setupGraph(c)
}

func setupTeachers(c *goC8.Client) {
	exists, err := c.CheckCollectionExists(fabric, collectionTeachers)
	CheckError(err, "Error CheckCollectionExists: ")
	if !exists {
		// if not create collection
		collType := collection_req.DocumentCollectionType
		allowUserKeys := true
		err = c.CreateNewCollection(fabric, collectionTeachers, allowUserKeys, collType)
		CheckError(err, "Error CreateNewCollection")

		//// import city data
		//silent := false
		//jsonDocument := sample_data.GetTeacherData()
		//_, err = c.CreateNewDocument(fabric, collectionTeachers, silent, jsonDocument, nil)
		//CheckError(err, "Error CreateNewDocument")
	}
}

func setupCourses(c *goC8.Client) {
	exists, err := c.CheckCollectionExists(fabric, collectionLectures)
	CheckError(err, "Error CheckCollectionExists: ")

	if !exists {
		// if not create collection
		collType := collection_req.DocumentCollectionType
		allowUserKeys := false
		err = c.CreateNewCollection(fabric, collectionLectures, allowUserKeys, collType)
		CheckError(err, "Error CreateNewCollection")

		// import  data
		silent := false
		jsonDocument := GetLecturesData()
		_, err = c.CreateNewDocument(fabric, collectionLectures, silent, jsonDocument, nil)
		CheckError(err, "Error CreateNewDocument")
	}

}

func setupGraph(c *goC8.Client) {

}
