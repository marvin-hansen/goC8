package main

import (
	"github.com/marvin-hansen/goC8"
	"github.com/marvin-hansen/goC8/src"
	"github.com/marvin-hansen/goC8/src/requests/collection_req"
	"github.com/marvin-hansen/goC8/src/utils"
)

const (
	apiKey   = "email.root.your-secret-key"
	endpoint = "https://YOUR-ID-us-west.paas.macrometa.io"
	fabric   = "yourFabric"
	timeout  = 5 // http connection timeout in seconds
	collName = "TestCollection"
	verbose  = true
)

func main() {
	// Chose between document collection for storing JSON and edge collections that are used for graphs.
	collType := collection_req.DocumentCollectionType

	println("Create new config ")
	config := goC8.NewConfig(apiKey, endpoint, fabric, timeout)

	println("Create new client with config ")
	c := goC8.NewClient(config)

	println("Create new collection: " + collName)
	createCollErr := c.CreateNewCollection(fabric, collName, false, collType)
	utils.CheckError(createCollErr, "Failed to create a new collection. "+collName)

	println("Create new document! ")
	silent := false // When true, an empty reply will be retruned. If false, the document ID will be returned
	jsonDocument := getTestInsertData()

	res, createDocErr := c.CreateNewDocument(fabric, collName, silent, jsonDocument, nil)
	utils.CheckError(createDocErr, "Failed to create a new document. "+collName)

	if verbose {
		if res != nil {
			for _, v := range *res {
				println(v.String())
			}
		}
	}

	println("Get a document! ")
	key := "4"
	getRes, getDocErr := c.GetDocument(fabric, collName, key)
	utils.CheckError(getDocErr, "Failed to get document: "+key)
	printJsonRes(getRes)

}

func printJsonRes(res src.JsonResponder) {
	if verbose {
		println(res.String())
	}
}
