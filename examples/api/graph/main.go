package main

import (
	"github.com/marvin-hansen/goC8"
	"github.com/marvin-hansen/goC8/src/requests/collection_req"
	"github.com/marvin-hansen/goC8/src/utils"
)

const (
	// client config
	apiKey   = "email.root.secretkey"
	endpoint = "https://YOUR-ID-us-west.paas.macrometa.io"
	fabric   = "SouthEastAsia"
	timeout  = 5 // http connection timeout in seconds
	// collection & graph config
	verbose          = true
	graph            = "airline"
	collectionID     = "cities"
	edgeCollectionID = "flights"
)

func main() {
	println("Create new config ")
	config := goC8.NewConfig(apiKey, endpoint, fabric, timeout)

	println("Create new client with config ")
	c := goC8.NewClient(config)

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

		// Create a geo index
		field := "location"
		geoJson := true
		_, err = c.Index.CreateGeoIndex(fabric, collectionID, field, geoJson)
		utils.CheckError(err, "Error CreateNewDocument")
		utils.DbgPrint("Create GeoIndex on: "+field, verbose)
	}
	
}
