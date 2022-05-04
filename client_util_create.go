package goC8

import (
	"github.com/marvin-hansen/goC8/types"
	"github.com/marvin-hansen/goC8/utils"
)

func CreateCollection(c *Client, fabric, collectionName string, collectionType types.CollectionType, allowUserKeys bool) {
	exists, err := c.Collection.CheckCollectionExists(fabric, collectionName)
	utils.CheckError(err, "Error CheckCollectionExists: ")
	if !exists {
		err = c.Collection.CreateNewCollection(fabric, collectionName, allowUserKeys, collectionType)
		utils.CheckError(err, "Error CreateNewCollection")
	}
}

func ImportCollectionData(c *Client, fabric, collectionName string, jsonDocument []byte, silent bool) {
	exists, err := c.Collection.CheckCollectionExists(fabric, collectionName)
	utils.CheckError(err, "Error CheckCollectionExists: ")
	if exists {
		_, err = c.Document.CreateNewDocument(fabric, collectionName, silent, jsonDocument, nil)
		utils.CheckError(err, "Error CreateNewDocument")
	}
}

func CreateGraph(c *Client, fabric, graphName string, jsonGraph []byte) {
	exists, err := c.Graph.CheckGraphExists(fabric, graphName)
	utils.CheckError(err, "Error CheckGraphExists: ")
	if !exists {
		// if not, create  graph
		_, createGraphErr := c.Graph.CreateGraph(fabric, jsonGraph)
		utils.CheckError(createGraphErr, "Error CreateGraph")
	}
}
