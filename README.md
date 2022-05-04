# goC8

Welcome to the GitHub page for goC8, a Golang client for the macrometa global data platform (GDN).

## About

[Macrometa](https://www.macrometa.com/) is a secure, global data platform with integrated pub/sub, stream processing,
search, functions, and 4 databases all through one single API. GDN clients for [Python](https://github.com/Macrometacorp/pyC8), [Java](https://github.com/Macrometacorp/c84j), and [Javascript](https://github.com/Macrometacorp/jsC8) are available from [Macrometa](https://github.com/Macrometacorp).
goC8 is an open source client that implements the Macrometa API for Golang. 

## Use cases

* Legal compliance with [data residency regulations](https://incountry.com/blog/data-residency-laws-by-country-overview/) & [EU GDPR regulations](https://www.omnitas.se/data-residency-affected-by-gdpr-and-schrems/)
* Global data to cluster co-location for low latency data access
* [Web Assembly (WASM) storage](https://itnext.io/webassemply-with-golang-by-scratch-e05ec5230558)
* Unified single storage API 
* Microservice storage
* Kubernetes storage
* Mobile app storage 
* Single, globally distributed & fail-safe, storage backend shared between web, mobile, & cluster applications
* Event storage & streaming for real-time data analytics wth real-time client updates

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
Geo Fabric. Timeout refers to the http connection timeout in seconds. If you do not have these value at hand, please read the [setup guide](setup.md) for details.

## Code examples

* Collections [code example](examples/api/collections/main.go) and [tests](tests/collection)
* Documents [code example](examples/api/documentstore/main.go) and [tests](tests/document)
* Graph [code example](examples/api/graph/main.go) and [tests](tests/graph)
* Index [code example](examples/api/index/main.go) and [tests](tests/index)
* KeyValue [code example](examples/api/kv/main.go) and [tests](tests/kv)
* [Flight example](examples/flight)
* [University example](examples/university)

Golang demo apps from Macrometa
* [Salesforce PII](https://github.com/Macrometacorp/demo-salesforce-pii)
* [Data privacy and residency](https://github.com/Macrometacorp/demo-pii)


## Usage: Graph, collection, & document API

Full code in [Flight example](examples/flight)

```Go
package main

import (
	"log"
	"github.com/marvin-hansen/goC8"
	"github.com/marvin-hansen/goC8/examples/sample_data"
	"github.com/marvin-hansen/goC8/tests/utils"
)

const (
	// client config
	apiKey   = "email.root.secretkey.xxxxxxxxxxxxxxxxx"
	endpoint = "https://YOUR-ID-us-west.paas.macrometa.io"
	fabric   = "MyFabric"
	timeout  = 5 // http connection timeout in seconds 
	// collection, document & graph config
	delete           = false
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
	
	println("Setup: Create Graph, collections & import data")
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
	res, err := c.Query(fabric, q, nil, nil)
	utils.CheckError(err, "Error Query: "+q)
	utils.PrintQuery(res, verbose)
}

// setup uses the built-in low-code utilities to create a collection, index, graph & import data.
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
```

## Make reference

```bash 
Setup: 
    make check                  Checks all requirements.
 
Test: 
    make test-all               Runs all API tests.
    make test-collection        Tests collection API. 
    make test-document          Tests document API. 
    make test-index             Tests index API. 
    make test-kv                Tests key-value API. 
```


## Known issues and solutions

### Illegal key while creating ....

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

### VPackError error: Expecting digit

Error number 400 / 600 

This error means that the provided JSON doesn't conform to the standard. 

Use a JSON [validator](https://jsonformatter.curiousconcept.com/) and check for common issues:

* Missing comma between array values
* Comma after last attribute
* Singe quote ' instead of double " quote

Notice, the JSON unpacker is rather fragile so whatever "small fixes" the validator suggests, 
apply them otherwise the error persists. 

## Author

* Marvin Hansen
* GPG key ID: 210D39BC
* Github key ID: 369D5A0B210D39BC
* GPG Fingerprint: 4B18 F7B2 04B9 7A72 967E 663E 369D 5A0B 210D 39BC
* Public key: [key](pubkey.txt)

## Licence

* [MIT Licence](LICENSE)
* Software is "as is" without any warranty. 
