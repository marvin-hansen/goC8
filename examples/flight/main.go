package main

import (
	"github.com/marvin-hansen/goC8"
	"github.com/marvin-hansen/goC8/examples/sample_data"
	"github.com/marvin-hansen/goC8/requests/collection_req"
	"log"
)

const (
	verbose          = true
	fabric           = "SouthEastAsia"
	graph            = "airline"
	collectionID     = "cities"
	edgeCollectionID = "flights"
)

func main() {
	c := goC8.NewClient(nil)

	setupCities(c)

	//teardown(c)

	//setupFlights(c)
	// query graph here

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

		// We have to create a geo index before importing geoJson
		field := "location"
		geoJson := true
		_, err = c.CreateGeoIndex(fabric, collectionID, field, geoJson)
		checkError(err, "Error CreateNewDocument")

		// import city data
		silent := false
		jsonDocument := sample_data.GetCityData()
		_, err = c.CreateNewDocument(fabric, collectionID, silent, jsonDocument, nil)
		checkError(err, "Error CreateNewDocument")
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

		// import flight data
		silent := false
		jsonDocument := sample_data.GetFlightData()
		_, err = c.CreateNewDocument(fabric, edgeCollectionID, silent, jsonDocument, nil)
		checkError(err, "Error CreateNewCollection")
	}

	// test if graph exists
	exists, err = c.CheckGraphExists(fabric, graph)
	checkError(err, "Error CheckGraphExists")
	if !exists {
		// if not, create one
		jsonGraph := sample_data.GetAirlineGraph()
		_, createErr := c.CreateGraph(fabric, jsonGraph)
		checkError(createErr, "Error CreateGraph")
	}
}

func teardown(c *goC8.Client) {
	// test if graph exists
	exists, err := c.CheckGraphExists(fabric, graph)
	checkError(err, "Error CheckGraphExists: ")
	if exists {
		// if so delete  graph
		_, delErr := c.DeleteGraph(fabric, graph, true)
		checkError(delErr, "Error DeleteGraph: ")
	}

	// test if city collection exists
	exists, err = c.CheckCollectionExists(fabric, collectionID)
	checkError(err, "Error CheckCollectionExists: ")
	if exists {
		// if so, delete
		err = c.DeleteCollection(fabric, collectionID, false)
		checkError(err, "Error DeleteCollection")

	}

	// test if flight collection exists
	_, err = c.CheckCollectionExists(fabric, edgeCollectionID)
	checkError(err, "Error CheckCollectionExists")
	if exists {
		// if so, delete
		err = c.DeleteCollection(fabric, edgeCollectionID, false)
		checkError(err, "Error DeleteCollection")
	}
}

func checkError(err error, msg string) {
	if err != nil {
		log.Println("error: " + err.Error())
		log.Fatalf(msg)
	}
}
