package main

import (
	"github.com/marvin-hansen/goC8"
	"github.com/marvin-hansen/goC8/types"
	"strconv"
)

const (
	apiKey   = "email.root.secretkey"
	endpoint = "https://YOUR-ID-us-west.paas.macrometa.io"
	fabric   = "uswest"
	timeout  = 5 // http connection timeout in seconds
)

func main() {
	// Chose between document collection for storing JSON and edge collections that are used for graphs.
	collType := types.DocumentCollectionType
	collName := "TestCollection"

	println("Create new config ")
	config := goC8.NewConfig(apiKey, endpoint, fabric, timeout)

	println("Create new client with config ")
	c := goC8.NewClient(config)

	println("Create new collection: " + collName)
	createCollErr := c.Collection.CreateNewCollection(fabric, collName, false, collType)
	goC8.CheckError(createCollErr, "Failed to create a new collection. "+collName)

	println("Test if collection exists: " + collName)
	exists, checkErr := c.Collection.CheckCollectionExists(fabric, collName)
	goC8.CheckError(checkErr, "Failed to check collection exists. ")
	println("Collection exists: " + strconv.FormatBool(exists))

	println("Get collection Info: " + collName)
	res, getCollErr := c.Collection.GetCollectionInfo(fabric, collName)
	goC8.CheckError(getCollErr, "Failed to get a new collection. ")
	println(res.String())

	println("Get all collections in the fabric: " + fabric)
	resGetAll, errGetAll := c.Collection.GetAllCollections(fabric)
	goC8.CheckError(errGetAll, "Failed to get all collections. ")
	println(resGetAll.String())

	println("Delete collection: " + collName)
	delErr := c.Collection.DeleteCollection(fabric, collName, false)
	goC8.CheckError(delErr, "Failed to delete collection: "+collName)

}
