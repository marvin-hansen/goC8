package goC8

import r "github.com/marvin-hansen/goC8/requests/graph_req"

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
func (c Client) GetEdge(fabric, graphName, collectionName, edgeName string) (response *r.ResponseForGetEdge, err error) {
	req := r.NewRequestForGetEdge(fabric, graphName, collectionName, edgeName)
	response = r.NewResponseForGetEdge()
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c Client) GetAllVertices(fabric, graphName string) (response *r.ResponseForGetAllVertices, err error) {
	req := r.NewRequestForGetAllVertices(fabric, graphName)
	response = r.NewResponseForGetAllVertices()
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}
