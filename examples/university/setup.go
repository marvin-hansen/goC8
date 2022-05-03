package main

import (
	"github.com/marvin-hansen/goC8"
	"github.com/marvin-hansen/goC8/src/requests/collection_req"
	"github.com/marvin-hansen/goC8/src/utils"
)

func setup(c *goC8.Client) {
	setupTeachers(c)
	setupCourses(c)
	setupEdges(c)
	setupGraph(c)
}

func setupTeachers(c *goC8.Client) {
	exists, err := c.Collection.CheckCollectionExists(fabric, collectionTeachers)
	utils.CheckError(err, "Error CheckCollectionExists: ")
	if !exists {
		// if not create collection
		collType := collection_req.DocumentCollectionType
		allowUserKeys := true
		err = c.Collection.CreateNewCollection(fabric, collectionTeachers, allowUserKeys, collType)
		utils.CheckError(err, "Error CreateNewCollection")

		// import city data
		silent := false
		jsonDocument := GetTeacherData()
		_, err = c.Document.CreateNewDocument(fabric, collectionTeachers, silent, jsonDocument, nil)
		utils.CheckError(err, "Error CreateNewDocument")
	}
}

func setupCourses(c *goC8.Client) {
	exists, err := c.Collection.CheckCollectionExists(fabric, collectionLectures)
	utils.CheckError(err, "Error CheckCollectionExists: ")

	if !exists {
		// if not create collection
		collType := collection_req.DocumentCollectionType
		allowUserKeys := false
		err = c.Collection.CreateNewCollection(fabric, collectionLectures, allowUserKeys, collType)
		utils.CheckError(err, "Error CreateNewCollection")

		// import  data
		silent := false
		jsonDocument := GetLecturesData()
		_, err = c.Document.CreateNewDocument(fabric, collectionLectures, silent, jsonDocument, nil)
		utils.CheckError(err, "Error CreateNewDocument")
	}
}

func setupEdges(c *goC8.Client) {
	exists, err := c.Collection.CheckCollectionExists(fabric, edgeCollectionTeach)
	utils.CheckError(err, "Error CheckCollectionExists: ")

	if !exists {
		// if not create collection
		collType := collection_req.EdgeCollectionType
		allowUserKeys := false
		err = c.Collection.CreateNewCollection(fabric, edgeCollectionTeach, allowUserKeys, collType)
		utils.CheckError(err, "Error CreateNewCollection")

		// import  data
		silent := false
		jsonDocument := GetTeachEdgeData()
		_, err = c.Document.CreateNewDocument(fabric, edgeCollectionTeach, silent, jsonDocument, nil)
		utils.CheckError(err, "Error CreateNewDocument")
	}
}

func setupGraph(c *goC8.Client) {
	// test if graph exists
	exists, err := c.Graph.CheckGraphExists(fabric, graph)
	utils.CheckError(err, "Error CheckGraphExists: ")
	if !exists {
		// if so create graph
		jsonGraph := GetUniversityGraphDefinition()
		_, createGraphErr := c.Graph.CreateGraph(fabric, jsonGraph)
		utils.CheckError(createGraphErr, "Error CreateGraph")
	}
}
