package goC8

import (
	r "github.com/marvin-hansen/goC8/requests/graph_req"
	"strings"
)

// GetAllGraphs
// Lists all graphs stored in this GeoFabric.
// https://macrometa.com/docs/api#/operations/ListAllGraphs
func (c Client) GetAllGraphs(fabric string) (response *r.ResponseForGetAllGraphs, err error) {
	req := r.NewRequestForGetAllGraphs(fabric)
	response = r.NewResponseForGetAllGraphs()
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

// CreateGraph
// The creation of a graph requires the name of the graph and a definition of its edges.
// Sample edge definition:
// {
//  "edgeDefinitions": [
//    {
//      "collection": "edgeName",
//      "from": [
//        "sourceNode"
//      ],
//      "to": [
//        "destinationNode"
//      ]
//    }
//  ],
//  "name": "exampleGraph",
//  "options": {}
//}
// https://macrometa.com/docs/api#/operations/CreateAGraph
func (c Client) CreateGraph(fabric string, jsonGraph []byte) (response *r.ResponseForCreateGraph, err error) {
	req := r.NewRequestForCreateGraph(fabric, jsonGraph)
	response = r.NewResponseForCreateGraph()
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetGraph
// Retrieve information for a graph. Returns the edge definitions and orphan collections.
// https://macrometa.com/docs/api#/operations/GetAGraph
func (c Client) GetGraph(fabric, graphName string) (response *r.ResponseForGetGraph, err error) {
	req := r.NewRequestForGetGraph(fabric, graphName)
	response = r.NewResponseForGetGraph()
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

// CheckGraphExists
// returns true if a graph for the given name exists
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

// DeleteGraph
// Remove an existing graph object by name. Optionally all collections not used by other graphs can be removed as well.
// https://macrometa.com/docs/api#/operations/DropAGraph
func (c Client) DeleteGraph(fabric, graphName string, dropCollections bool) (response *r.ResponseForDeleteGraph, err error) {
	req := r.NewRequestForDeleteGraph(fabric, graphName, dropCollections)
	response = r.NewResponseForDeleteGraph()
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}
