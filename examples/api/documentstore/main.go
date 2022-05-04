package main

import (
	"github.com/marvin-hansen/goC8"
	"github.com/marvin-hansen/goC8/types"
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
	collType := types.DocumentCollectionType

	println("Create new config ")
	config := goC8.NewConfig(apiKey, endpoint, fabric, timeout)

	println("Create new client with config ")
	c := goC8.NewClient(config)

	println("Create new collection: " + collName)
	createCollErr := c.Collection.CreateNewCollection(fabric, collName, false, collType)
	goC8.CheckError(createCollErr, "Failed to create a new collection. "+collName)

	println("Create new document! ")
	silent := false // When true, an empty reply will be retruned. If false, the document ID will be returned
	jsonDocument := getTestInsertData()

	res, createDocErr := c.Document.CreateNewDocument(fabric, collName, silent, jsonDocument, nil)
	goC8.CheckError(createDocErr, "Failed to create a new document. "+collName)

	if verbose {
		if res != nil {
			for _, v := range *res {
				println(v.String())
			}
		}
	}

	println("Get a document! ")
	key := "4"
	getRes, getDocErr := c.Document.GetDocument(fabric, collName, key)
	goC8.CheckError(getDocErr, "Failed to get document: "+key)
	printJsonRes(getRes)

}

func printJsonRes(res types.JsonResponder) {
	if verbose {
		println(res.String())
	}
}
