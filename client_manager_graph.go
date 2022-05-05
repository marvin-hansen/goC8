package goC8

import (
	"github.com/marvin-hansen/goC8/requests/graph_req"
	"github.com/marvin-hansen/goC8/utils"
	"strings"
	"time"
)

type GraphManager struct {
	client *Client
}

func NewGraphManager(client *Client) *GraphManager {
	return &GraphManager{client: client}
}

// GetAllGraphs
// Lists all graphs stored in this GeoFabric.
// https://macrometa.com/docs/api#/operations/ListAllGraphs
func (c GraphManager) GetAllGraphs(fabric string) (response *graph_req.ResponseForGetAllGraphs, err error) {
	if benchmark {
		defer utils.TimeTrack(time.Now(), "GetAllGraphs")
	}

	req := graph_req.NewRequestForGetAllGraphs(fabric)
	response = graph_req.NewResponseForGetAllGraphs()
	if err = c.client.Request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

// CreateGraph
// The creation of a graph requires the name of the graph and a definition of its edges.
// Note: Requires Administrator permissions to access the GeoFabric and Read Only access on every collection used within this graph.
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
func (c GraphManager) CreateGraph(fabric string, jsonGraph []byte) (response *graph_req.ResponseForCreateGraph, err error) {
	if benchmark {
		defer utils.TimeTrack(time.Now(), "CreateGraph")
	}

	req := graph_req.NewRequestForCreateGraph(fabric, jsonGraph)
	response = graph_req.NewResponseForCreateGraph()
	if err = c.client.Request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetGraph
// Retrieve information for a graph. Returns the edge definitions and orphan collections.
// https://macrometa.com/docs/api#/operations/GetAGraph
func (c GraphManager) GetGraph(fabric, graphName string) (response *graph_req.ResponseForGraph, err error) {
	if benchmark {
		defer utils.TimeTrack(time.Now(), "GetGraph")
	}

	req := graph_req.NewRequestForGetGraph(fabric, graphName)
	response = graph_req.NewResponseForGraph()
	if err = c.client.Request(req, response); err != nil {
		return nil, err
	}

	return response, nil
}

// CheckGraphExists
// returns true if a graph for the given name exists
func (c GraphManager) CheckGraphExists(fabric, graphName string) (exists bool, err error) {
	if benchmark {
		defer utils.TimeTrack(time.Now(), "CheckGraphExists")
	}

	req := graph_req.NewRequestForGetGraph(fabric, graphName)
	response := graph_req.NewResponseForGraph()
	if err = c.client.Request(req, response); err != nil {
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
func (c GraphManager) DeleteGraph(fabric, graphName string, dropCollections bool) (response *graph_req.ResponseForDeleteGraph, err error) {
	if benchmark {
		defer utils.TimeTrack(time.Now(), "DeleteGraph")
	}

	req := graph_req.NewRequestForDeleteGraph(fabric, graphName, dropCollections)
	response = graph_req.NewResponseForDeleteGraph()
	if err = c.client.Request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}
