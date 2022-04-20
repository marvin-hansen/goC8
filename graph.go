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
