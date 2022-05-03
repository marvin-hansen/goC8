package main

import (
	"github.com/marvin-hansen/goC8"
	"github.com/marvin-hansen/goC8/src/requests/collection_req"
)

func setup(c *goC8.Client) {
	setupCities(c)
	setupFlights(c)
	setupGraph(c)
}

func setupCities(c *goC8.Client) {
	// test if city collection exists
	exists, err := c.CheckCollectionExists(fabric, collectionID)
	checkError(err, "Error CheckCollectionExists: ")
	if !exists {
		// if not create collection
		collType := collection_req.DocumentCollectionType
		allowUserKeys := false
		err = c.CreateNewCollection(fabric, collectionID, allowUserKeys, collType)
		checkError(err, "Error CreateNewCollection")
		dbgPrint("Create collection: " + collectionID)

		// We have to create a geo index before importing geoJson
		field := "location"
		geoJson := true
		_, err = c.CreateGeoIndex(fabric, collectionID, field, geoJson)
		checkError(err, "Error CreateNewDocument")
		dbgPrint("Create GeoIndex on: " + field)

		// import city data
		silent := false
		jsonDocument := GetCityData()
		_, err = c.CreateNewDocument(fabric, collectionID, silent, jsonDocument, nil)
		checkError(err, "Error CreateNewDocument")
		dbgPrint("Imported data into: " + collectionID)
	}

}

func setupFlights(c *goC8.Client) {
	// test if flight collection exists
	exists, err := c.CheckCollectionExists(fabric, edgeCollectionID)
	checkError(err, "Error CheckCollectionExists")
	if !exists {
		// if not create edge collection
		collType := collection_req.EdgeCollectionType
		err = c.CreateNewCollection(fabric, edgeCollectionID, false, collType)
		checkError(err, "Error CreateNewCollection")
		dbgPrint("Create collection: " + edgeCollectionID)

		// import flight data
		silent := false
		jsonDocument := GetFlightData()
		_, err = c.CreateNewDocument(fabric, edgeCollectionID, silent, jsonDocument, nil)
		checkError(err, "Error CreateNewCollection")
		dbgPrint("Imported data into: " + edgeCollectionID)
	}
}

func setupGraph(c *goC8.Client) {
	// test if graph exists
	exists, err := c.CheckGraphExists(fabric, graph)
	checkError(err, "Error CheckGraphExists: ")
	if !exists {
		// if so create graph
		jsonGraph := GetAirlineGraph()
		_, createGraphErr := c.CreateGraph(fabric, jsonGraph)
		checkError(createGraphErr, "Error CreateGraph")
		dbgPrint("Created Graph: " + graph)
	}
}
