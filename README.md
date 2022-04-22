# goC8

Welcome to the GitHub page for goC8, a Golang driver for the macrometa global data platform (GDN).

## About Macrometa

[Macrometa](https://www.macrometa.com/) is a secure, global data platform with integrated pub/sub, stream processing,
search, functions, and 4 databases all through one single API.

Supported databases:

* Collections
* Documents

## Install

```Bash
go get github.com/marvin-hansen/goC8/client
```

## Authentication

Currently, only API key authentication is supported.

## Configuration

If you have not worked with macrometa's plattform before, you have to do a one-time setup. See this [guide](setup.md)
for details about creating a data fabric for your project.

The client config requires the following settings:

* Api Key
* Endpoint
* Fabric
* Timeout

Api Key refers to the generated api access key. Endpoint refers to the POP provided by the GDN. Fabric refers to the GDN
Geo Fabric. Timeout refers to the http connection timeout in seconds.

If you do not have these value at hand, please read the [setup guide](setup.md) for details.

### Examples

* [Collections](examples/collections)
* [Documents](examples/documentstore)
* [Flight example](examples/flight)
* [Tests](tests)

### Known issues and solutions

##### Illegal key while creating .... 

This error means that the provided key, name, or identifier in question does not adhere to the established convention.
Names for collections, documents, graphs, and edges follow the same convention as keys, 
there is a high chance that the identifier contains a blank or non-allowed character triggering this error. 
When writing insert functions with custom primary keys, it is paramount to stick to the [naming convention](README_KEYS.md)
to prevent runtime errors.

For details, please look at the 

* [key naming convention](README_KEYS.md) 
* [Official Arango Docs](https://www.arangodb.com/docs/stable/data-modeling-naming-conventions-document-keys.html)
* [Stack Overflow](https://stackoverflow.com/questions/68118693/arangodb-illegal-document-key-error-while-creating-graph)

Notice, the problem in the Stack Overflow question was the blank in the graph name. 
A fix would be eliminating the blank by renaming the graph from "Friends visit" to "FriendsVisit" to adhere to the naming convention. 


### Collections

```Go

import (
   "github.com/marvin-hansen/goC8"
   "github.com/marvin-hansen/goC8/requests/collection_req"
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
   createCollErr := c.CreateNewCollection(fabric, collName, false, collType)
   checkError(createCollErr, "Failed to create a new collection. "+collName)
   
   println("Get collection Info: " + collName)
   res, getCollErr := c.GetCollectionInfo(fabric, collName)
   checkError(getCollErr, "Failed to get a new collection. ")
   
   println(res.String())
   
   println("Get all collections in the fabric: " + fabric)
   resGetAll, getAllErr := c.GetAllCollections(fabric)
   checkError(getAllErr, "Failed to get all collections for fabric: "+fabric)
   println(resGetAll.String())

    println("Delete collection Info: " + collName)
    delErr := c.DeleteCollection(fabric, collName, false)
    checkError(delErr, "Failed to delete collection: "+collName)
}

func checkError(err error, msg string) {
    if err != nil {
        log.Println("error: " + err.Error())
        log.Fatalf(msg)
}
}
```

### Documents

* [Full document store example](examples/documentstore)

```Go

package main

import (
	"github.com/marvin-hansen/goC8"
	"github.com/marvin-hansen/goC8/requests/collection_req"
	"log"
)

const (
	apiKey   = "email.root.secretkey"
	endpoint = "https://YOUR-ID-us-west.paas.macrometa.io"
	fabric   = "uswest"
	timeout  = 5 // http connection timeout in seconds
)

const verbose = true

func main() {
	collName := "TestCollection"
	config := goC8.NewConfig(apiKey, endpoint, fabric, timeout)
	c := goC8.NewClient(config)

	println("Create new document! ")
	silent := false // When true, an empty reply will be retruned. If false, the document ID will be returned
	jsonDocument := getTestInsertData()

	res, createDocErr := c.CreateNewDocument(fabric, collName, silent, jsonDocument, nil)
	checkError(createDocErr, "Failed to create a new document. "+collName)

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
	checkError(getDocErr, "Failed to get document: "+key)
	printJsonRes(getRes)

}

func checkError(err error, msg string) {
	if err != nil {
		log.Println("error: " + err.Error())
		log.Fatalf(msg)
	}
}

func printJsonRes(res goC8.JsonResponder) {
	if verbose {
		println(res.String())
	}
}
```

## Author

* Marvin Hansen
* GPG key ID: 210D39BC
* Github key ID: 369D5A0B210D39BC
* GPG Fingerprint: 4B18 F7B2 04B9 7A72 967E 663E 369D 5A0B 210D 39BC
* Public key: [key](pubkey.txt)

## Licence

* [MIT Licence](LICENSE)
* Software is "as is" without any warranty. 