package goC8

import (
	"errors"
	"github.com/marvin-hansen/goC8/requests/index_req"
	"github.com/marvin-hansen/goC8/types"
	"github.com/marvin-hansen/goC8/utils"
)

func CreateCollection(c *Client, fabric, collectionName string, collectionType types.CollectionType, allowUserKeys bool) {
	exists, err := c.Collection.CheckCollectionExists(fabric, collectionName)
	utils.CheckError(err, "Error CheckCollectionExists: ")
	if !exists {
		// if not, create
		err = c.Collection.CreateNewCollection(fabric, collectionName, allowUserKeys, collectionType)
		utils.CheckError(err, "Error CreateNewCollection")
	}
}

func ImportCollectionData(c *Client, fabric, collectionName string, jsonDocument []byte, silent bool) {
	exists, err := c.Collection.CheckCollectionExists(fabric, collectionName)
	utils.CheckError(err, "Error CheckCollectionExists: ")
	if exists {
		// if not, create
		_, err = c.Document.CreateNewDocument(fabric, collectionName, silent, jsonDocument, nil)
		utils.CheckError(err, "Error CreateNewDocument")
	}
}

func CreateGraph(c *Client, fabric, graphName string, jsonGraph []byte) {
	exists, err := c.Graph.CheckGraphExists(fabric, graphName)
	utils.CheckError(err, "Error CheckGraphExists: ")
	if !exists {
		// if not, create
		_, createGraphErr := c.Graph.CreateGraph(fabric, jsonGraph)
		utils.CheckError(createGraphErr, "Error CreateGraph")
	}
}

func CreateIndex(c *Client, fabric, collectionName, field string, indexType types.IndexType, deduplicate, sparse, unique, geoJson bool) (response *index_req.IndexEntry, err error) {
	switch indexType {
	case types.FulltextIndex:
		return c.Index.CreateFulltextIndex(fabric, collectionName, field, 3)
	case types.HashIndex:
		return c.Index.CreateHashIndex(fabric, collectionName, field, deduplicate, sparse, unique)
	case types.GeoIndex:
		return c.Index.CreateGeoIndex(fabric, collectionName, field, geoJson)
	case types.PersistentIndex:
		return c.Index.CreatePersistentIndex(fabric, collectionName, field, deduplicate, sparse, unique)
	case types.SkipListIndex:
		return c.Index.CreateSkipListIndex(fabric, collectionName, field, deduplicate, sparse, unique)
	case types.TTLIndex:
		return c.Index.CreateTTLIndex(fabric, collectionName, field, 5)
	default:
		return nil, errors.New("unknown index type")
	}
}
