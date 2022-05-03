package main

import (
	"github.com/marvin-hansen/goC8"
	"github.com/marvin-hansen/goC8/src/requests/collection_req"
	"log"
)

const (
	apiKey   = "email.root.secretkey"
	endpoint = "https://YOUR-ID-us-west.paas.macrometa.io"
	fabric   = "uswest"
	timeout  = 5 // http connection timeout in seconds
)

func main() {
	// Chose between document collection for storing JSON and edge collections that are used for graphs.
	collType := collection_req.DocumentCollectionType
	collName := "TestCollection"

	println("Create new config ")
	config := goC8.NewConfig(apiKey, endpoint, fabric, timeout)

	println("Create new client with config ")
	c := goC8.NewClient(config)

	println("Create new collection: " + collName)
	createCollErr := c.Collection.CreateNewCollection(fabric, collName, false, collType)
	checkError(createCollErr, "Failed to create a new collection. "+collName)

	println("Get collection Info: " + collName)
	res, getCollErr := c.Collection.GetCollectionInfo(fabric, collName)
	checkError(getCollErr, "Failed to get a new collection. ")
	println(res.String())

	println("Delete collection Info: " + collName)
	delErr := c.Collection.DeleteCollection(fabric, collName, false)
	checkError(delErr, "Failed to delete collection: "+collName)

}

func checkError(err error, msg string) {
	if err != nil {
		log.Println("error: " + err.Error())
		log.Fatalf(msg)
	}
}
