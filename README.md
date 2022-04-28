# goC8

Welcome to the GitHub page for goC8, a Golang driver for the macrometa global data platform (GDN).

## About

[Macrometa](https://www.macrometa.com/) is a secure, global data platform with integrated pub/sub, stream processing,
search, functions, and 4 databases all through one single API. 

goC8 is an open source client that implements the macrometa API. 
Currently, goC8 supports the following data api:

* Collections [Code Example](examples/api/collections)
* Documents [Code Example](examples/api/documentstore)
* Graph [Code Example](examples/api/graph)
* Index [Code Example](examples/api/index)

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

### Usage

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