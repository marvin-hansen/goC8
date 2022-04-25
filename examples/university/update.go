package main

import (
	"github.com/marvin-hansen/goC8"
	"github.com/marvin-hansen/goC8/requests/collection_req"
	"github.com/marvin-hansen/goC8/utils"
)

func update(c *goC8.Client) {
	addTutorials(c)
	addTutorEdge(c)
	updateAddEdge(c)
}

func addTutorials(c *goC8.Client) {
	exists, err := c.CheckCollectionExists(fabric, collectionTutorials)
	utils.CheckError(err, "Error CheckCollectionExists: ")
	if !exists {
		// if not create collection
		collType := collection_req.DocumentCollectionType
		allowUserKeys := false
		err = c.CreateNewCollection(fabric, collectionTutorials, allowUserKeys, collType)
		utils.CheckError(err, "Error CreateNewCollection")

		// import data
		silent := false
		jsonDocument := GetTutorialsData()
		_, err = c.CreateNewDocument(fabric, collectionTutorials, silent, jsonDocument, nil)
		utils.CheckError(err, "Error CreateNewDocument")
	}
}

func addTutorEdge(c *goC8.Client) {
	// test if flight collection exists
	exists, err := c.CheckCollectionExists(fabric, edgeCollectionTutors)
	utils.CheckError(err, "Error CheckCollectionExists")

	if !exists {
		// if not create edge collection
		collType := collection_req.EdgeCollectionType
		err = c.CreateNewCollection(fabric, edgeCollectionTutors, false, collType)
		utils.CheckError(err, "Error CreateNewCollection")

		// import data
		silent := false
		jsonDocument := GetTutorsEdgeData()
		_, err = c.CreateNewDocument(fabric, edgeCollectionTutors, silent, jsonDocument, nil)
		utils.CheckError(err, "Error CreateNewCollection")
	}
}

func updateAddEdge(c *goC8.Client) {
	// check if collection exists. It should, just in case
	exists, err := c.CheckCollectionExists(fabric, edgeCollectionTutors)
	utils.CheckError(err, "Error CheckCollectionExists")
	if exists {
		// if exists, update graph with new edge collection
		collectionName := edgeCollectionTutors
		sourceVertex := "teachers"
		destinationVertex := "tutorials"

		_, addErr := c.AddEdgeCollection(fabric, graph, collectionName, sourceVertex, destinationVertex)
		if addErr != nil {
			return
		}
	} else {
		println("Can't update graph. Add edge collection: " + edgeCollectionTutors)
	}
}
