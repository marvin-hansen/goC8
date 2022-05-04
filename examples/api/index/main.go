package main

import (
	"github.com/marvin-hansen/goC8"
	"github.com/marvin-hansen/goC8/types"
	utils2 "github.com/marvin-hansen/goC8/utils"
)

const (
	apiKey   = "email.root.secretkey"
	endpoint = "https://YOUR-ID-us-west.paas.macrometa.io"
	fabric   = "uswest"
	timeout  = 5 // http connection timeout in seconds
	collName = "TestCollection"
	// Chose between document collection for storing JSON
	// and edge collections that are used for storing graphs edges.
	collType = types.DocumentCollectionType
	verbose  = true // set to false to disable console printout
)

func setup(c *goC8.Client) {
	exists, err := c.Collection.CheckCollectionExists(fabric, collName)
	utils2.CheckError(err, "Error CheckCollectionExists: ")
	if !exists {
		println("Create new collection: " + collName)
		createCollErr := c.Collection.CreateNewCollection(fabric, collName, false, collType)
		utils2.CheckError(createCollErr, "Failed to create a new collection. "+collName)
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
	res, getIndexErr := c.Index.GetIndexes(fabric, collName)
	utils2.CheckError(getIndexErr, "Failed to fetch all indices for "+collName)
	utils2.PrintRes(res, verbose)

	println("Create Fulltext index for: " + collName)
	fulltextIndexField := "Summary"
	fulltextMinLength := 3
	ftRes, createFTErr := c.Index.CreateFulltextIndex(fabric, collName, fulltextIndexField, fulltextMinLength)
	utils2.CheckError(createFTErr, "Failed to create fulltext index for "+collName)
	utils2.PrintRes(ftRes, verbose)

	println("Create Geo index for: " + collName)
	geoField := "location"
	geoJson := true
	geoRes, geoErr := c.Index.CreateGeoIndex(fabric, collName, geoField, geoJson)
	utils2.CheckError(geoErr, "Failed to create geo index for "+collName)
	utils2.PrintRes(geoRes, verbose)

	println("Create Hash index for: " + collName)
	hashField := "Keywords"
	deduplicate := true
	sparse := true
	unique := true
	hashRes, hashErr := c.Index.CreateHashIndex(fabric, collName, hashField, deduplicate, sparse, unique)
	utils2.CheckError(hashErr, "Failed to create hash index for "+collName)
	utils2.PrintRes(hashRes, verbose)

	println("Create persistent index for: " + collName)
	persistentField := "Content"
	perRes, perErr := c.Index.CreatePersistentIndex(fabric, collName, persistentField, deduplicate, sparse, unique)
	utils2.CheckError(perErr, "Failed to create persistent index for "+collName)
	utils2.PrintRes(perRes, verbose)

	println("Create SkipList index for: " + collName)
	skipField := "SkipText"
	skipRes, skipErr := c.Index.CreateSkipListIndex(fabric, collName, skipField, deduplicate, sparse, unique)
	utils2.CheckError(skipErr, "Failed to create SkipList index for "+collName)
	utils2.PrintRes(skipRes, verbose)

	println("Create TTL index for: " + collName)
	ttlField := "requests"
	ttl := 10
	ttlRes, ttlErr := c.Index.CreateTTLIndex(fabric, collName, ttlField, ttl)
	utils2.CheckError(ttlErr, "Failed to create TTL index for "+collName)
	utils2.PrintRes(ttlRes, verbose)

}
