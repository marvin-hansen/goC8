package goC8

import (
	"github.com/marvin-hansen/goC8/requests/graph_req/vertex_req"
	"time"
)

// GetAllVertices
// Lists all vertex collections within this graph.
// https://macrometa.com/docs/api#/operations/ListVertexCollections
func (c Client) GetAllVertices(fabric, graphName string) (response *vertex_req.ResponseForGetAllVertices, err error) {
	if benchmark {
		defer TimeTrack(time.Now(), "GetAllVertices")
	}

	req := vertex_req.NewRequestForGetAllVertices(fabric, graphName)
	response = vertex_req.NewResponseForGetAllVertices()
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetVertex
// Gets a vertex from the given collection.
// https://macrometa.com/docs/api#/operations/GetAVertex
func (c Client) GetVertex(fabric, graphName, collectionName, vertexKey string) (response *vertex_req.ResponseForGetVertex, err error) {
	if benchmark {
		defer TimeTrack(time.Now(), "GetVertex")
	}

	req := vertex_req.NewRequestForGetVertex(fabric, graphName, collectionName, vertexKey)
	response = vertex_req.NewResponseForGetVertex()
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}
