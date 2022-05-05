package goC8

import (
	"github.com/marvin-hansen/goC8/requests/graph_req"
	"github.com/marvin-hansen/goC8/requests/graph_req/vertex_req"
	"github.com/marvin-hansen/goC8/utils"
	"time"
)

// GetAllVertices
// Lists all vertex collections within this graph.
// https://macrometa.com/do    cs/api#/operations/ListVertexCollections
func (c GraphManager) GetAllVertices(fabric, graphName string) (response *vertex_req.ResponseForGetAllVertices, err error) {
	if benchmark {
		defer utils.TimeTrack(time.Now(), "GetAllVertices")
	}

	req := vertex_req.NewRequestForGetAllVertices(fabric, graphName)
	response = vertex_req.NewResponseForGetAllVertices()
	if err = c.client.Request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetVertex
// Gets a vertex from the given collection.
// https://macrometa.com/docs/api#/operations/GetAVertex
func (c GraphManager) GetVertex(fabric, graphName, collectionName, vertexKey string) (response *graph_req.ResponseForVertex, err error) {
	if benchmark {
		defer utils.TimeTrack(time.Now(), "GetVertex")
	}

	req := vertex_req.NewRequestForGetVertex(fabric, graphName, collectionName, vertexKey)
	response = graph_req.NewResponseForVertex()
	if err = c.client.Request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

// CreateVertex
// Adds a vertex to the given collection.
// returnNew - boolean  Define if the response should contain the complete new version of the document.Show all...
// https://macrometa.com/docs/api#/operations/CreateAVertex
func (c GraphManager) CreateVertex(fabric, graphName, collectionName string, jsonDef []byte, returnNew bool) (response *graph_req.ResponseForVertex, err error) {
	if benchmark {
		defer utils.TimeTrack(time.Now(), "CreateVertex")
	}

	req := vertex_req.NewRequestForCreateVertex(fabric, graphName, collectionName, jsonDef, returnNew)
	response = graph_req.NewResponseForVertex()
	if err = c.client.Request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

// AddVertexCollection
// Adds a vertex collection to the set of orphan collections of the graph. If the collection does not exist, it will be created.
func (c GraphManager) AddVertexCollection(fabric, graphName, vertexCollectionName string) (response *graph_req.ResponseForGraph, err error) {
	if benchmark {
		defer utils.TimeTrack(time.Now(), "AddVertexCollection")
	}

	req := vertex_req.NewRequestForAddVertexCollection(fabric, graphName, vertexCollectionName)
	response = graph_req.NewResponseForGraph()
	if err = c.client.Request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

// DeleteVertexCollection
// Removes a vertex collection from the graph and optionally removes the collection, if it is not used in any other graph.
// It only removes vertex collections that are no longer part of edge definitions.
// dropCollection - Remove the collection as well. Collection is only removed if it is not used in other graphs.
// https://macrometa.com/docs/api#/operations/RemoveVertexCollection
func (c GraphManager) DeleteVertexCollection(fabric, graphName, collectionName string, dropCollections bool) (response *graph_req.ResponseForGraph, err error) {
	if benchmark {
		defer utils.TimeTrack(time.Now(), "DeleteVertexCollection")
	}

	req := vertex_req.NewRequestForDeleteVertexCollection(fabric, graphName, collectionName, dropCollections)
	response = graph_req.NewResponseForGraph()
	err = c.client.Request(req, response)
	return response, CheckReturnError(err)
}

// DeleteVertex
// Removes a vertex from the collection.
// returnOld boolean - Define if a presentation of the deleted document should be returned
// https://macrometa.com/docs/api#/operations/RemoveAVertex
func (c GraphManager) DeleteVertex(fabric, graphName, collectionName, vertexKey string, returnOld bool) (response *graph_req.ResponseForVertex, err error) {
	if benchmark {
		defer utils.TimeTrack(time.Now(), "DeleteVertex")
	}

	req := vertex_req.NewRequestForDeleteVertex(fabric, graphName, collectionName, vertexKey, returnOld)
	response = graph_req.NewResponseForVertex()
	err = c.client.Request(req, response)
	return response, CheckReturnError(err)
}
