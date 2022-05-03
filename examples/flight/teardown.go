package main

import "github.com/marvin-hansen/goC8"

func teardown(c *goC8.Client) {
	// test if graph exists
	exists, err := c.CheckGraphExists(fabric, graph)
	checkError(err, "Error CheckGraphExists: ")
	if exists {
		// if so delete  graph
		_, delErr := c.DeleteGraph(fabric, graph, true)
		checkError(delErr, "Error DeleteGraph: ")
		dbgPrint("Deleted graph and all of its collections: " + graph)
	}

	// It's possible that the graph has been deleted without dropCollections,
	// but the underlying collections still exists. Let's check one by one.

	// test if city collection exists
	exists, err = c.Collection.CheckCollectionExists(fabric, collectionID)
	checkError(err, "Error CheckCollectionExists: ")
	if exists {
		// if so, delete
		err = c.Collection.DeleteCollection(fabric, collectionID, false)
		checkError(err, "Error DeleteCollection")
		dbgPrint("Deleted collection: " + collectionID)
	}

	// test if flight collection exists
	_, err = c.Collection.CheckCollectionExists(fabric, edgeCollectionID)
	checkError(err, "Error CheckCollectionExists")
	if exists {
		// if so, delete
		err = c.Collection.DeleteCollection(fabric, edgeCollectionID, false)
		checkError(err, "Error DeleteCollection")
		dbgPrint("Deleted collection: " + edgeCollectionID)
	}
}
