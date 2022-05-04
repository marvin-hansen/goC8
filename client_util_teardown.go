package goC8

import (
	"github.com/marvin-hansen/goC8/utils"
)

func TeardownGraph(c *Client, fabric, graphName string, dropCollections bool) {
	exists, err := c.Graph.CheckGraphExists(fabric, graphName)
	utils.CheckError(err, "Error CheckGraphExists: ")
	if exists {
		// if so delete graph
		_, createGraphErr := c.Graph.DeleteGraph(fabric, graphName, dropCollections)
		utils.CheckError(createGraphErr, "Error CreateGraph")
	}
}

func TeardownCollection(c *Client, fabric, collectionName string) {
	exists, err := c.Collection.CheckCollectionExists(fabric, collectionName)
	utils.CheckError(err, "Error CheckCollectionExists: ")
	if exists {
		// if so, delete collection
		err = c.Collection.DeleteCollection(fabric, collectionName, false)
		utils.CheckError(err, "Error DeleteCollection")
	}
}
