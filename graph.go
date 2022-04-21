package goC8

import (
	r "github.com/marvin-hansen/goC8/requests/graph_req"
	"strings"
)

// GetAllGraphs
// Lists all graphs stored in this GeoFabric.
func (c Client) GetAllGraphs(fabric string) (response *r.ResponseForGetAllGraphs, err error) {
	req := r.NewRequestForGetAllGraphs(fabric)
	response = r.NewResponseForGetAllGraphs()
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetGraph
// Retrieve information for a graph. Returns the edge definitions and orphan collections.
func (c Client) GetGraph(fabric, graphName string) (response *r.ResponseForGetGraph, err error) {
	req := r.NewRequestForGetGraph(fabric, graphName)
	response = r.NewResponseForGetGraph()
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c Client) CheckGraphExists(fabric, graphName string) (exists bool, err error) {
	req := r.NewRequestForGetGraph(fabric, graphName)
	response := r.NewResponseForGetGraph()
	if err = c.request(req, response); err != nil {
		if strings.Contains(err.Error(), "1924") { //  Number=1924,  Error Message=graph 'graphName' not found
			return false, nil
		} else {
			return false, err // Any other error
		}
	}

	return true, nil
}

// GetAllEdges
// Lists all edge collections within this graph.
func (c Client) GetAllEdges(fabric, graphName string) (response *r.ResponseForGetAllEdges, err error) {
	req := r.NewRequestForGetAllEdges(fabric, graphName)
	response = r.NewResponseForGetAllEdges()
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetEdge
// Gets an edge from the given collection.
func (c Client) GetEdge(fabric, graphName, collectionName, edgeKey string) (response *r.ResponseForGetEdge, err error) {
	req := r.NewRequestForGetEdge(fabric, graphName, collectionName, edgeKey)
	response = r.NewResponseForGetEdge()
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetAllVertices
// Lists all vertex collections within this graph.
func (c Client) GetAllVertices(fabric, graphName string) (response *r.ResponseForGetAllVertices, err error) {
	req := r.NewRequestForGetAllVertices(fabric, graphName)
	response = r.NewResponseForGetAllVertices()
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetVertex
// Gets a vertex from the given collection.
func (c Client) GetVertex(fabric, graphName, collectionName, vertexKey string) (response *r.ResponseForGetVertex, err error) {
	req := r.NewRequestForGetVertex(fabric, graphName, collectionName, vertexKey)
	response = r.NewResponseForGetVertex()
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}
