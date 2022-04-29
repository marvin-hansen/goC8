package goC8

import (
	"github.com/marvin-hansen/goC8/requests/graph_req"
	"time"
)

// GetAllVertices
// Lists all vertex collections within this graph.
// https://macrometa.com/docs/api#/operations/ListVertexCollections
func (c Client) GetAllVertices(fabric, graphName string) (response *graph_req.ResponseForGetAllVertices, err error) {
	if benchmark {
		defer TimeTrack(time.Now(), "GetAllVertices")
	}

	req := graph_req.NewRequestForGetAllVertices(fabric, graphName)
	response = graph_req.NewResponseForGetAllVertices()
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetVertex
// Gets a vertex from the given collection.
// https://macrometa.com/docs/api#/operations/GetAVertex
func (c Client) GetVertex(fabric, graphName, collectionName, vertexKey string) (response *graph_req.ResponseForGetVertex, err error) {
	if benchmark {
		defer TimeTrack(time.Now(), "GetVertex")
	}

	req := graph_req.NewRequestForGetVertex(fabric, graphName, collectionName, vertexKey)
	response = graph_req.NewResponseForGetVertex()
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}
