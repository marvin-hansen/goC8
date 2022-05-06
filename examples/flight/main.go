package main

import (
	"github.com/marvin-hansen/goC8"
	"github.com/marvin-hansen/goC8/examples/sample_data"
	"github.com/marvin-hansen/goC8/types"
)

const (
	// client config
	apiKey   = "email.root.secretkey"
	endpoint = "https://YOUR-ID-us-west.paas.macrometa.io"
	fabric   = "SouthEastAsia"
	timeout  = 5 // http connection timeout in seconds
	// collection & graph config
	delete           = false
	verbose          = true
	silent           = false
	graph            = "airline"
	collectionID     = "cities"
	edgeCollectionID = "flights"
)

func main() {
	println("Create new config ")
	config := goC8.NewConfig(apiKey, endpoint, fabric, timeout)

	println("Create new client with config ")
	c := goC8.NewClient(config)

	println("Run setup")
	setup(c)

	println("Query: Document & Graph")
	query(c)

	if delete {
		println("Teardown: Delete Graph & Data")
		teardown(c)
	}
}

func query(c *goC8.Client) {
	var q = ""
	var msg = ""

	q = sample_data.GetAllCitiesQuery()
	msg = "Get all cities."
	runQuery(c, q, msg)

	q = sample_data.GetBreadthFirstQuery()
	msg = "Get all cities with a direct flight to New York."
	runQuery(c, q, msg)

	q = sample_data.GetShortestPathQuery()
	msg = "Get the shortest path from San Francisco to Paris."
	runQuery(c, q, msg)

	q = sample_data.GetShortestDistanceQuery()
	msg = "Get the distance on the shortest path from San Francisco to Paris."
	runQuery(c, q, msg)

	q = sample_data.GetNearestCities()
	msg = "Get the 2 nearest cities to a specified latitude and longitude."
	runQuery(c, q, msg)

	q = sample_data.GetCitiesMaxDistance()
	msg = "Get the cities that are no more than 2500km away from houston."
	runQuery(c, q, msg)
}

func runQuery(c *goC8.Client, q, msg string) {
	println(msg)
	res, err := c.Query.Query(fabric, q, nil, nil)
	goC8.CheckError(err, "Error Query: "+q)
	goC8.PrintQuery(res, verbose)
}

// setup & teardown use the built-in utilities to create collection, index, graph & import data.
func setup(c *goC8.Client) {
	goC8.CreateCollection(c, fabric, collectionID, types.DocumentCollectionType, false)
	field := "location" // We have to create a geo index with geoJson enabled on field location before importing data
	goC8.CreateIndex(c, fabric, collectionID, field, types.GeoIndex, true, true, true, true)
	goC8.ImportCollectionData(c, fabric, collectionID, sample_data.GetCityData(), silent)
	goC8.CreateCollection(c, fabric, edgeCollectionID, types.EdgeCollectionType, false)
	goC8.ImportCollectionData(c, fabric, edgeCollectionID, sample_data.GetFlightData(), silent)
	goC8.CreateGraph(c, fabric, graph, sample_data.GetAirlineGraph())
}

func teardown(c *goC8.Client) {
	goC8.TeardownGraph(c, fabric, graph, true)
	goC8.TeardownCollection(c, fabric, collectionID)
	goC8.TeardownCollection(c, fabric, edgeCollectionID)
}
