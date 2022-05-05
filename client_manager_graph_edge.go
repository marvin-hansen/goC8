package goC8

import (
	"github.com/marvin-hansen/goC8/requests/collection_req"
	"github.com/marvin-hansen/goC8/requests/graph_req"
	"github.com/marvin-hansen/goC8/requests/graph_req/edge_req"
	"github.com/marvin-hansen/goC8/utils"
	"strings"
	"time"
)

// GetAllEdges
// Lists all edge collections within this graph.
// https://macrometa.com/docs/api#/operations/ListEdgedefinitions
func (c GraphManager) GetAllEdges(fabric, graphName string) (response *edge_req.ResponseForGetAllEdges, err error) {
	if benchmark {
		defer utils.TimeTrack(time.Now(), "GetAllEdges")
	}

	req := edge_req.NewRequestForGetAllEdges(fabric, graphName)
	response = edge_req.NewResponseForGetAllEdges()
	if err = c.client.Request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

// CheckEdgeCollectionExists
// Returns true if the collection of the name exists. False otherwise.
func (c GraphManager) CheckEdgeCollectionExists(fabric, collectionName string) (exists bool, err error) {
	req := collection_req.NewRequestForGetCollectionInfo(fabric, collectionName)
	response := collection_req.NewResponseForGetCollectionInfo()
	if err = c.client.Request(req, response); err != nil {
		if strings.Contains(err.Error(), "1203") { // Number=1203,  Error Message=collection or view not found
			return false, nil
		} else {
			return false, err // Any other error
		}
	}
	return true, nil
}

// AddEdgeCollection
// Adds an additional edge definition to the graph. This edge definition has to contain a collection and an array of each from and to vertex collections.
// An edge definition can only be added if this definition is either not used in any other graph, or it is used with exactly the same definition.
// *  edgeCollectionName: The name of the edge collection to be used.
// *  sourceVertex (string): One or many vertex collections that can contain source vertices.
// *  destinationVertex (string): One or many vertex collections that can contain target vertices.
// https://macrometa.com/docs/api#/operations/AddEdgedefinition
func (c GraphManager) AddEdgeCollection(fabric, graphName, edgeCollectionName, sourceVertex, destinationVertex string) (response *edge_req.ResponseForAddEdgeCollection, err error) {
	if benchmark {
		defer utils.TimeTrack(time.Now(), "AddEdgeCollection")
	}

	req := edge_req.NewRequestForAddEdgeCollection(fabric, graphName, edgeCollectionName, sourceVertex, destinationVertex)
	response = edge_req.NewResponseForAddEdgeCollection()
	if err = c.client.Request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

// CreateEdge
// Creates a new edge in the collection
// sourceVertex: The source vertex of this edge. Has to be valid within the used edge definition.
// destinationVertex: The target vertex of this edge. Has to be valid within the used edge definition.
func (c GraphManager) CreateEdge(fabric, graphName, edgeCollectionName, sourceVertex, destinationVertex string, returnNew bool) (response *edge_req.ResponseForCreateEdge, err error) {
	if benchmark {
		defer utils.TimeTrack(time.Now(), "CreateEdge")
	}

	req := edge_req.NewRequestForCreateEdge(fabric, graphName, edgeCollectionName, sourceVertex, destinationVertex, returnNew)
	response = edge_req.NewResponseForCreateEdge()
	if err = c.client.Request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetEdge
// Gets an edge from the given collection.
// https://macrometa.com/docs/api#/operations/GetAnEdge
func (c GraphManager) GetEdge(fabric, graphName, collectionName, edgeKey string) (response *graph_req.ResponseForEdge, err error) {
	if benchmark {
		defer utils.TimeTrack(time.Now(), "GetEdge")
	}

	req := edge_req.NewRequestForGetEdge(fabric, graphName, collectionName, edgeKey)
	response = graph_req.NewResponseForEdge()
	if err = c.client.Request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

// CheckEdgeExists
// returns true if the edge exists in the given collection
func (c GraphManager) CheckEdgeExists(fabric, graphName, collectionName, edgeKey string) (exists bool, err error) {
	if benchmark {
		defer utils.TimeTrack(time.Now(), "CheckEdgeExists")
	}

	req := edge_req.NewRequestForGetEdge(fabric, graphName, collectionName, edgeKey)
	response := graph_req.NewResponseForEdge()
	if err = c.client.Request(req, response); err != nil {
		if strings.Contains(err.Error(), "1202") { //  Code=404, Number=1202,  Error Message=document not found
			return false, nil
		} else {
			return false, err // Any other error
		}
	}
	return true, nil
}

// ReplaceEdge
// Replaces the data of an edge in the collection.
// https://macrometa.com/docs/api#/operations/ReplaceAnEdge
func (c GraphManager) ReplaceEdge(fabric, graphName, edgeCollectionName, sourceVertex, destinationVertex string, dropCollections bool) (response *edge_req.ResponseForReplaceEdge, err error) {
	if benchmark {
		defer utils.TimeTrack(time.Now(), "ReplaceEdge")
	}

	req := edge_req.NewRequestForReplaceEdge(fabric, graphName, edgeCollectionName, sourceVertex, destinationVertex, dropCollections)
	response = edge_req.NewResponseForReplaceEdge()
	if err = c.client.Request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c GraphManager) UpdateEdge(fabric, graphName, edgeCollectionName, edgeKey string, jsonUpdate []byte, keepNull, returnOld, returnNew bool) (response *graph_req.ResponseForEdge, err error) {
	if benchmark {
		defer utils.TimeTrack(time.Now(), "UpdateEdge")
	}

	req := edge_req.NewRequestForUpdateEdge(fabric, graphName, edgeCollectionName, edgeKey, jsonUpdate, keepNull, returnOld, returnNew)
	response = graph_req.NewResponseForEdge()
	if err = c.client.Request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

// DeleteEdge
// Removes an edge from the collection.
// https://macrometa.com/docs/api#/operations/RemoveAnEdge
func (c GraphManager) DeleteEdge(fabric, graphName, collectionName, edgeKey string, returnOld bool) (response *graph_req.ResponseForEdge, err error) {
	if benchmark {
		defer utils.TimeTrack(time.Now(), "DeleteEdge")
	}

	req := edge_req.NewRequestForDeleteEdge(fabric, graphName, collectionName, edgeKey, returnOld)
	response = graph_req.NewResponseForEdge()
	err = c.client.Request(req, response)
	return response, CheckReturnError(err)
}

// DeleteEdgeCollection
// Remove one edge definition from the graph.
// This only removes the edge collection, the vertex collections remain untouched and can still be used in your queries.
// https://macrometa.com/docs/api#/operations/RemoveAnEdgedefinitionFromTheGraph
func (c GraphManager) DeleteEdgeCollection(fabric, graphName, collectionName string, dropCollections bool) (response *graph_req.ResponseForGraph, err error) {
	if benchmark {
		defer utils.TimeTrack(time.Now(), "DeleteEdgeCollection")
	}

	req := edge_req.NewRequestForDeleteEdgeCollection(fabric, graphName, collectionName, dropCollections)
	response = graph_req.NewResponseForGraph()
	err = c.client.Request(req, response)
	return response, CheckReturnError(err)
}
