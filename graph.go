package goC8

import (
	graph_req2 "github.com/marvin-hansen/goC8/src/requests/graph_req"
	"strings"
	"time"
)

// GetAllGraphs
// Lists all graphs stored in this GeoFabric.
// https://macrometa.com/docs/api#/operations/ListAllGraphs
func (c Client) GetAllGraphs(fabric string) (response *graph_req2.ResponseForGetAllGraphs, err error) {
	if benchmark {
		defer TimeTrack(time.Now(), "GetAllGraphs")
	}

	req := graph_req2.NewRequestForGetAllGraphs(fabric)
	response = graph_req2.NewResponseForGetAllGraphs()
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
func (c Client) CreateGraph(fabric string, jsonGraph []byte) (response *graph_req2.ResponseForCreateGraph, err error) {
	if benchmark {
		defer TimeTrack(time.Now(), "CreateGraph")
	}

	req := graph_req2.NewRequestForCreateGraph(fabric, jsonGraph)
	response = graph_req2.NewResponseForCreateGraph()
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetGraph
// Retrieve information for a graph. Returns the edge definitions and orphan collections.
// https://macrometa.com/docs/api#/operations/GetAGraph
func (c Client) GetGraph(fabric, graphName string) (response *graph_req2.ResponseForGetGraph, err error) {
	if benchmark {
		defer TimeTrack(time.Now(), "GetGraph")
	}

	req := graph_req2.NewRequestForGetGraph(fabric, graphName)
	response = graph_req2.NewResponseForGetGraph()
	if err = c.request(req, response); err != nil {
		return nil, err
	}

	return response, nil
}

// CheckGraphExists
// returns true if a graph for the given name exists
func (c Client) CheckGraphExists(fabric, graphName string) (exists bool, err error) {
	if benchmark {
		defer TimeTrack(time.Now(), "CheckGraphExists")
	}

	req := graph_req2.NewRequestForGetGraph(fabric, graphName)
	response := graph_req2.NewResponseForGetGraph()
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
func (c Client) DeleteGraph(fabric, graphName string, dropCollections bool) (response *graph_req2.ResponseForDeleteGraph, err error) {
	if benchmark {
		defer TimeTrack(time.Now(), "DeleteGraph")
	}

	req := graph_req2.NewRequestForDeleteGraph(fabric, graphName, dropCollections)
	response = graph_req2.NewResponseForDeleteGraph()
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}
