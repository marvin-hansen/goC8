package main

import (
	"github.com/marvin-hansen/goC8"
	"github.com/marvin-hansen/goC8/src/utils"
)

func teardown(c *goC8.Client) {
	// test if graph exists
	exists, err := c.Graph.CheckGraphExists(fabric, graph)
	utils.CheckError(err, "Error CheckGraphExists: ")
	if exists {
		// if so delete  graph
		_, delErr := c.Graph.DeleteGraph(fabric, graph, true)
		utils.CheckError(delErr, "Error DeleteGraph: ")
		utils.DbgPrint("Deleted graph and all of its collections: "+graph, verbose)
	}

	// It's possible that the graph has been deleted without dropCollections,
	// but the underlying collections still exists. Let's check one by one.

	// test if city collection exists
	exists, err = c.Collection.CheckCollectionExists(fabric, collectionID)
	utils.CheckError(err, "Error CheckCollectionExists: ")
	if exists {
		// if so, delete
		err = c.Collection.DeleteCollection(fabric, collectionID, false)
		utils.CheckError(err, "Error DeleteCollection")
		utils.DbgPrint("Deleted collection: "+collectionID, verbose)
	}

	// test if flight collection exists
	_, err = c.Collection.CheckCollectionExists(fabric, edgeCollectionID)
	utils.CheckError(err, "Error CheckCollectionExists")
	if exists {
		// if so, delete
		err = c.Collection.DeleteCollection(fabric, edgeCollectionID, false)
		utils.CheckError(err, "Error DeleteCollection")
		utils.DbgPrint("Deleted collection: "+edgeCollectionID, verbose)
	}
}
