package goC8

import (
	graph_req2 "github.com/marvin-hansen/goC8/src/requests/graph_req"
	"github.com/marvin-hansen/goC8/src/utils"
	"time"
)

// GetAllVertices
// Lists all vertex collections within this graph.
// https://macrometa.com/docs/api#/operations/ListVertexCollections
func (c Client) GetAllVertices(fabric, graphName string) (response *graph_req2.ResponseForGetAllVertices, err error) {
	if benchmark {
		defer utils.TimeTrack(time.Now(), "GetAllVertices")
	}

	req := graph_req2.NewRequestForGetAllVertices(fabric, graphName)
	response = graph_req2.NewResponseForGetAllVertices()
	if err = c.Request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetVertex
// Gets a vertex from the given collection.
// https://macrometa.com/docs/api#/operations/GetAVertex
func (c Client) GetVertex(fabric, graphName, collectionName, vertexKey string) (response *graph_req2.ResponseForGetVertex, err error) {
	if benchmark {
		defer utils.TimeTrack(time.Now(), "GetVertex")
	}

	req := graph_req2.NewRequestForGetVertex(fabric, graphName, collectionName, vertexKey)
	response = graph_req2.NewResponseForGetVertex()
	if err = c.Request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}
