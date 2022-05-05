package goC8

import (
	"github.com/marvin-hansen/goC8/requests/graph_req/vertex_req"
	"github.com/marvin-hansen/goC8/utils"
	"time"
)

// GetAllVertices
// Lists all vertex collections within this graph.
// https://macrometa.com/docs/api#/operations/ListVertexCollections
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
func (c GraphManager) GetVertex(fabric, graphName, collectionName, vertexKey string) (response *vertex_req.ResponseForGetVertex, err error) {
	if benchmark {
		defer utils.TimeTrack(time.Now(), "GetVertex")
	}

	req := vertex_req.NewRequestForGetVertex(fabric, graphName, collectionName, vertexKey)
	response = vertex_req.NewResponseForGetVertex()
	if err = c.client.Request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c GraphManager) AddVertex(fabric, graphName, vertexCollectionName string) (response *vertex_req.ResponseForGetVertex, err error) {
	if benchmark {
		defer utils.TimeTrack(time.Now(), "GetVertex")
	}

	req := vertex_req.NewRequestForAddVertexCollection(fabric, graphName, vertexCollectionName)
	response = vertex_req.NewResponseForGetVertex()
	if err = c.client.Request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}
