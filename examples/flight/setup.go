package main

import (
	"github.com/marvin-hansen/goC8"
	"github.com/marvin-hansen/goC8/src/requests/collection_req"
	"github.com/marvin-hansen/goC8/src/utils"
)

func setup(c *goC8.Client) {
	setupCities(c)
	setupFlights(c)
	setupGraph(c)
}

func setupCities(c *goC8.Client) {
	// test if city collection exists
	exists, err := c.Collection.CheckCollectionExists(fabric, collectionID)
	utils.CheckError(err, "Error CheckCollectionExists: ")
	if !exists {
		// if not create collection
		collType := collection_req.DocumentCollectionType
		allowUserKeys := false
		err = c.Collection.CreateNewCollection(fabric, collectionID, allowUserKeys, collType)
		utils.CheckError(err, "Error CreateNewCollection")
		utils.DbgPrint("Create collection: "+collectionID, verbose)

		// We have to create a geo index before importing geoJson
		field := "location"
		geoJson := true
		_, err = c.Index.CreateGeoIndex(fabric, collectionID, field, geoJson)
		utils.CheckError(err, "Error CreateNewDocument")
		utils.DbgPrint("Create GeoIndex on: "+field, verbose)

		// import city data
		silent := false
		jsonDocument := GetCityData()
		_, err = c.Document.CreateNewDocument(fabric, collectionID, silent, jsonDocument, nil)
		utils.CheckError(err, "Error CreateNewDocument")
		utils.DbgPrint("Imported data into: "+collectionID, verbose)
	}

}

func setupFlights(c *goC8.Client) {
	// test if flight collection exists
	exists, err := c.Collection.CheckCollectionExists(fabric, edgeCollectionID)
	utils.CheckError(err, "Error CheckCollectionExists")
	if !exists {
		// if not create edge collection
		collType := collection_req.EdgeCollectionType
		err = c.Collection.CreateNewCollection(fabric, edgeCollectionID, false, collType)
		utils.CheckError(err, "Error CreateNewCollection")
		utils.DbgPrint("Create collection: "+edgeCollectionID, verbose)

		// import flight data
		silent := false
		jsonDocument := GetFlightData()
		_, err = c.Document.CreateNewDocument(fabric, edgeCollectionID, silent, jsonDocument, nil)
		utils.CheckError(err, "Error CreateNewCollection")
		utils.DbgPrint("Imported data into: "+edgeCollectionID, verbose)
	}
}

func setupGraph(c *goC8.Client) {
	// test if graph exists
	exists, err := c.Graph.CheckGraphExists(fabric, graph)
	utils.CheckError(err, "Error CheckGraphExists: ")
	if !exists {
		// if so create graph
		jsonGraph := GetAirlineGraph()
		_, createGraphErr := c.Graph.CreateGraph(fabric, jsonGraph)
		utils.CheckError(createGraphErr, "Error CreateGraph")
		utils.DbgPrint("Created Graph: "+graph, verbose)
	}
}
