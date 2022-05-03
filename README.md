# goC8

Welcome to the GitHub page for goC8, a Golang driver for the macrometa global data platform (GDN).

## About

[Macrometa](https://www.macrometa.com/) is a secure, global data platform with integrated pub/sub, stream processing,
search, functions, and 4 databases all through one single API. goC8 is an open source client that implements the macrometa API.  
Currently, goC8 supports the following:

* Collections [code example](examples/api/collections) and [tests](tests/collection)
* Documents [code example](examples/api/documentstore) and [tests](tests/document)
* Graph [code example](examples/api/graph) and [tests](tests/graph)
* Index [code example](examples/api/index) and [tests](tests/index)
* KeyValue [code example](examples/api/kv) and [tests](tests/kv)

GDN clients for [Python](https://github.com/Macrometacorp/pyC8), [Java](https://github.com/Macrometacorp/c84j), and [Javascript](https://github.com/Macrometacorp/jsC8) are available from [Macrometa](https://github.com/Macrometacorp). 

## GDN use case
* Legal compliance with [data residency regulations](https://incountry.com/blog/data-residency-laws-by-country-overview/) & [EU GDPR regulations](https://www.omnitas.se/data-residency-affected-by-gdpr-and-schrems/)
* Global low latency web applications  
* Global data to cluster co-location for low latency data access

Suppose an e-commerce site wants to serve customers in the US, EU, and Asia. However, EU regulations require GDPR compliance
that do not apply to the US & Asia. Chine, however, has strict data residency requirements that do not apply to any other country, but must
be respected to enter the Chinese market. With the GDN, a dedicated EU data fabric allows the company to comply with EU regulations while
storing and processing data of EU customers only on EU territory. Likewise, a dedicated US data fabric ensures low latency especially when 
combined with Cloudflare cache while ascertain that US data & privacy regulations only apply to US customers. For Asia, 
a third data fabric may cover all of South East Asia except China. 

An internal data analytic service then pulls data from the US & Asia fabric, but may not access the EU data fabric
unless EU private data remain masked at all times. For performance reasons, the data analytics cluster resides in a POP 
shared by both the US & Asia data fabric as to stream real-time analytics back to the eCommerce store. 

Because of strict data residence regulations in China, a private cloud deployment on a Chinese cloud provider 
would then separate data from chinese users from the rest of the world while complying with Chinese regulations.
Note, in any case, seek legal advice on national and international privacy & data processing before entering
regulated markets. 

## Golang & GDN use cases
* [Web Assembly (WASM) storage](https://itnext.io/webassemply-with-golang-by-scratch-e05ec5230558)
* Unified single storage API 
* Microservice storage
* Kubernetes storage
* Mobile app storage 
* Single, globally distributed & fail-safe, storage backend shared between web, mobile, & cluster applications
* Event storage & streaming for real-time data analytics wth real-time client updates

## Golang demo apps from Macrometa
* [Salesforce PII](https://github.com/Macrometacorp/demo-salesforce-pii)
* [Data privacy and residency](https://github.com/Macrometacorp/demo-pii)


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

### Code examples

* [Api usage](examples/api) 
* [Flight example](examples/flight)
* [University example](examples/university)
* [Tests](tests)

## Usage

### Graph & document API

Full code in [Flight example](examples/flight)

```Go
package main

import (
	"log"
	"github.com/marvin-hansen/goC8"
	"github.com/marvin-hansen/goC8/examples/sample_data"
	"github.com/marvin-hansen/goC8/tests/utils"
)

// client config
const (
	apiKey   = "email.root.secretkey.xxxxxxxxxxxxxxxxx"
	endpoint = "https://YOUR-ID-us-west.paas.macrometa.io"
	fabric   = "SouthEastAsia"
	timeout  = 5 // http connection timeout in seconds
)

// collection & graph config 
const (
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
	checkError(err, "Error Query: "+q)
	utils.PrintQuery(res, verbose)
}

func checkError(err error, msg string) {
	if err != nil {
		log.Println("error: " + err.Error())
		log.Fatalf(msg)
	}
}
```


### KV API 

Full code in [KV API example](examples/api/kv)

```Go
package main

import (
	"github.com/marvin-hansen/goC8"
	"github.com/marvin-hansen/goC8/src/requests/kv_req"
)

const (
	// client config
	apiKey   = "email.root.secretkey"
	endpoint = "https://YOUR-ID-us-west.paas.macrometa.io"
	fabric   = "SouthEastAsia"
	timeout  = 5 // http connection timeout in seconds
	// collection config
	delete         = false
	verbose        = true
	collectionName = "TestKVCollection"
	expiration     = false
)

func main() {
	println("Create new config ")
	config := goC8.NewConfig(apiKey, endpoint, fabric, timeout)

	println("Create new client with config ")
	c := goC8.NewClient(config)

	println("Create KV collection: " + collectionName)
	res, err := c.CreateNewKVCollection(fabric, collectionName, expiration, nil)
	goC8.CheckErrorLog(err, "Error creatin KV collection"+collectionName)
	goC8.PrintRes(res, verbose)

	println("Get all KV collection: ")
	res, err = c.GetAllKVCollections(fabric)
	goC8.CheckErrorLog(err, "Error getting KV collection"+collectionName)
	goC8.PrintRes(res, verbose)

	println("Insert values into KV collection: ")
	kvPair1 := kv_req.NewKVPair("key1", "value1", -1)
	kvPair2 := kv_req.NewKVPair("key2", "value2", -1)
	kvPair3 := kv_req.NewKVPair("key3", "value3", -1)
	kvPair4 := kv_req.NewKVPair("key4", "value4", -1)
	kvPair5 := kv_req.NewKVPair("key5", "value5", -1)

	kvCollection := kv_req.NewKVPairCollection(*kvPair1, *kvPair2, *kvPair3, *kvPair4, *kvPair5)
	resSet, errSet := c.SetKeyValuePairs(fabric, collectionName, *kvCollection)
	goC8.CheckErrorLog(errSet, "Error setting values into KV collection"+collectionName)
	goC8.PrintRes(resSet, verbose)

	println("Count all items in KV collection: ")
	resCount, errCount := c.CountKVCollection(fabric, collectionName)
	goC8.CheckErrorLog(errCount, "Error counting values into KV collection"+collectionName)
	goC8.PrintRes(resCount, verbose)

	println("Get single value from KV collection: ")
	key := "value1"
	resOneVal, errOneVal := c.GetValue(fabric, collectionName, key)
	goC8.CheckErrorLog(errOneVal, "Error getting one value from KV collection"+collectionName)
	goC8.PrintRes(resOneVal, verbose)

	println("Delete single value from KV collection: ")
	resOneDel, errOneDel := c.DeleteValue(fabric, collectionName, key)
	goC8.CheckErrorLog(errOneDel, "Error deleting one value from KV collection"+collectionName)
	goC8.PrintRes(resOneDel, verbose)

	println("Update values into KV collection: ")
	kvPairUpdate1 := kv_req.NewKVPair("key1", "value42", -1)
	kvPairUpdate2 := kv_req.NewKVPair("key2", "value23", -1)

	kvCollectionUpdate := kv_req.NewKVPairCollection(*kvPairUpdate1, *kvPairUpdate2)
	resUpd, errUpd := c.SetKeyValuePairs(fabric, collectionName, *kvCollectionUpdate)
	goC8.CheckErrorLog(errUpd, "Error updating values into KV collection"+collectionName)
	goC8.PrintRes(resUpd, verbose)

	println("Get updated values from KV collection: ")
	resGetValues, errGetValues = c.GetAllValues(fabric, collectionName, offset, limit, nil)
	goC8.CheckErrorLog(errGetValues, "Error getting all values from KV collection"+collectionName)
	goC8.PrintRes(resGetValues, verbose)

	if delete {
		_, errTruncate := c.TruncateKVCollection(fabric, collectionName)
		goC8.CheckErrorLog(errTruncate, "Error truncating all values in KV collection"+collectionName)

		_, err = c.DeleteKVCollection(fabric, collectionName)
		goC8.CheckErrorLog(errTruncate, "Error deleting KV collection"+collectionName)
	}

}
```

## Make reference

```bash 
Setup: 
    make check                  Checks whether all project requirements are present.
     
Dev: 
    make build                  Builds the code base incrementally (fast).
    make rebuild                Rebuilds dependencies, build files, & code base (slow). Use after go mod changes.
    make stats                  Shows the latest project stats. 
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
