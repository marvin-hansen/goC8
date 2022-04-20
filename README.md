# goC8


## Install

```Bash
go get github.com/marvin-hansen/goC8/client

```

## ApiKey 

1. Create a free [macrometa account (https://www.macrometa.com/)
2. Create new API Key:
   1. Log in
   2. Go to Account
   3. Click on API keys 
   4. Click  on New Api key (Top right corner)
   5. Give it a key ID i.e. TestKey & click create
   6. Copy the actual KEY
   7. Put it into a file that is .gitignored to prevent committing the key
   
3. Bind the API key to the config


## Config

The client config requires the following settings:

* Api Key
* Endpoint
* Fabric 
* Timeout 

Api Key refers to the generated key as a string. 

Endpoint refers to the POP provided by the GDN. To get your endpoint URI:

1. Log in
2. Go to Geo Fabrics
3. Select _system
4. Select the region closest to your location i.e. Fremont
5. Copy the Endpoint URL and add it to the config 

Fabric refers to the GDN geofabric. To get your custom fabric:  
1. In the web console, go to Geo Fabrics
2. Click "New Geo Fabric" (Top right corner)
3. Give it a name
4. Select some initial locations. You can add more later. 
5. Copy the exact fabric name (case-sensitive) and add it to your config

### Working with collections. 

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