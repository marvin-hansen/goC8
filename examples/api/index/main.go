package main

import (
	"github.com/marvin-hansen/goC8"
	"github.com/marvin-hansen/goC8/src/requests/collection_req"
	"github.com/marvin-hansen/goC8/src/utils"
)

const (
	apiKey   = "email.root.secretkey"
	endpoint = "https://YOUR-ID-us-west.paas.macrometa.io"
	fabric   = "uswest"
	timeout  = 5 // http connection timeout in seconds
	collName = "TestCollection"
	// Chose between document collection for storing JSON
	// and edge collections that are used for storing graphs edges.
	collType = collection_req.DocumentCollectionType
	verbose  = true // set to false to disable console printout
)

func setup(c *goC8.Client) {
	exists, err := c.CheckCollectionExists(fabric, collName)
	utils.CheckError(err, "Error CheckCollectionExists: ")
	if !exists {
		println("Create new collection: " + collName)
		createCollErr := c.CreateNewCollection(fabric, collName, false, collType)
		utils.CheckError(createCollErr, "Failed to create a new collection. "+collName)
	}

}

func main() {
	println("Create new config ")
	config := goC8.NewConfig(apiKey, endpoint, fabric, timeout)

	println("Create new client with config ")
	c := goC8.NewClient(config)

	println("Run setup")
	setup(c)

	println("Get all indices for collection: " + collName)
	res, getIndexErr := c.GetIndexes(fabric, collName)
	utils.CheckError(getIndexErr, "Failed to fetch all indices for "+collName)
	utils.PrintRes(res, verbose)

	println("Create Fulltext index for: " + collName)
	fulltextIndexField := "Summary"
	fulltextMinLength := 3
	ftRes, createFTErr := c.CreateFulltextIndex(fabric, collName, fulltextIndexField, fulltextMinLength)
	utils.CheckError(createFTErr, "Failed to create fulltext index for "+collName)
	utils.PrintRes(ftRes, verbose)

	println("Create Geo index for: " + collName)
	geoField := "location"
	geoJson := true
	geoRes, geoErr := c.CreateGeoIndex(fabric, collName, geoField, geoJson)
	utils.CheckError(geoErr, "Failed to create geo index for "+collName)
	utils.PrintRes(geoRes, verbose)

	println("Create Hash index for: " + collName)
	hashField := "Keywords"
	deduplicate := true
	sparse := true
	unique := true
	hashRes, hashErr := c.CreateHashIndex(fabric, collName, hashField, deduplicate, sparse, unique)
	utils.CheckError(hashErr, "Failed to create hash index for "+collName)
	utils.PrintRes(hashRes, verbose)

	println("Create persistent index for: " + collName)
	persistentField := "Content"
	perRes, perErr := c.CreatePersistentIndex(fabric, collName, persistentField, deduplicate, sparse, unique)
	utils.CheckError(perErr, "Failed to create persistent index for "+collName)
	utils.PrintRes(perRes, verbose)

	println("Create SkipList index for: " + collName)
	skipField := "SkipText"
	skipRes, skipErr := c.CreateSkipListIndex(fabric, collName, skipField, deduplicate, sparse, unique)
	utils.CheckError(skipErr, "Failed to create SkipList index for "+collName)
	utils.PrintRes(skipRes, verbose)

	println("Create TTL index for: " + collName)
	ttlField := "requests"
	ttl := 10
	ttlRes, ttlErr := c.CreateTTLIndex(fabric, collName, ttlField, ttl)
	utils.CheckError(ttlErr, "Failed to create TTL index for "+collName)
	utils.PrintRes(ttlRes, verbose)

}
