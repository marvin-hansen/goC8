package main

import (
	"github.com/marvin-hansen/goC8"
	"github.com/marvin-hansen/goC8/types"
	"github.com/marvin-hansen/goC8/utils"
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
	utils.CheckError(createCollErr, "Failed to create a new collection. "+collName)

	println("Test if collection exists: " + collName)
	exists, checkErr := c.Collection.CheckCollectionExists(fabric, collName)
	utils.CheckError(checkErr, "Failed to check collection exists. ")
	println("Collection exists: " + strconv.FormatBool(exists))

	println("Get collection Info: " + collName)
	res, getCollErr := c.Collection.GetCollectionInfo(fabric, collName)
	utils.CheckError(getCollErr, "Failed to get a new collection. ")
	println(res.String())

	println("Get all collections in the fabric: " + fabric)
	resGetAll, errGetAll := c.Collection.GetAllCollections(fabric)
	utils.CheckError(errGetAll, "Failed to get all collections. ")
	println(resGetAll.String())

	println("Delete collection: " + collName)
	delErr := c.Collection.DeleteCollection(fabric, collName, false)
	utils.CheckError(delErr, "Failed to delete collection: "+collName)

}
