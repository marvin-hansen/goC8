## Implementation status

| Endpoint      | Implemented | Tested |
|---------------|-------------|--------|
| API Keys      | No          | No     |
| Auth          | No          | No     |
| Key Value     | Yes         | Yes    |
| Dymamo        | No          | No     |
| Documents     | Yes         | Yes    |
| Collections   | Yes         | Yes    |
| Query         | Yes         | Yes    |
| Query Workers | No          | No     |
| Indexes       | Yes         | Yes    |
| Graphs        | Yes         | No     |
| Streams       | No          | No     |
| Stream IO     | No          | No     |
| Search        | No          | No     |
| Geo Fabric    | No          | No     |
| Environments  | No          | No     |
| Users         | No          | No     |
| Billing       | No          | No     |
| Import/Export | No          | No     |
| Compute       | No          | No     |
| Data Centers  | No          | No     |
| Admin         | No          | No     |
| Support       | No          | No     |


## Collection API

| Http   | Endpoint                                 | Function                                           | Implemented | Tested | Done |
|--------|------------------------------------------|----------------------------------------------------|-------------|--------|------|
| Get    | /collection                              | Fetches the list of all collections.               | Yes         | Yes    | Yes  |
| Post   | /collection                              | Create collection.                                 | Yes         | Yes    | Yes  |
| Get    | /collection/{collection-name}            | Fetches the information about collection.          | Yes         | Yes    | Yes  |
| Delete | /collection/{collection-name}            | Remove the collection.                             | Yes         | Yes    | Yes  |
| Get    | /collection/{collection-name}/count      | Returns the number of documents in the collection. | Yes         | Yes    | Yes  |
| Put    | /collection/{collection-name}/truncate   | Truncate collection                                | Yes         | Yes    | Yes  |
| Put    | /collection/{collection-name}/properties | Changes the properties of a collection.            | Yes         | Yes    | Yes  |


## Document API

| Http   | Endpoint                     | Function                                     | Implemented | Tested | Done  |
|--------|------------------------------|----------------------------------------------|-------------|--------|-------|
| Delete | /document/{collection}       | Removes multiple documents                   | Yes         | Yes    | Yes   |
| Patch  | /document/{collection}       | Updates multiple documents                   | Yes         | Yes    | Yes   |
| Put    | /document/{collection}       | Replaces multiple documents                  | Yes         | Yes    | Yes   |
| Post   | /document/{collection}       | Create document                              | Yes         | Yes    | Yes   |
| Delete | /document/{collection}/{key} | Deletes one document                         | Yes         | Yes    | Yes   |
| Get    | /document/{collection}/{key} | Gets one document                            | Yes         | Yes    | Yes   |
| Head   | /document/{collection}/{key} | Like GET, but only returns the header fields | No          | No     | No[1] |
| Patch  | /document/{collection}/{key} | Updated one document                         | Yes         | Yes    | Yes   |
| Put    | /document/{collection}/{key} | Replace one document                         | Yes         | Yes    | Yes   |

[1] Custom HEAD method not supported in Golang http client

## Graph API
| Http   | Endpoint                                    | Function                                         | Implemented | Tested | Done  |
|--------|---------------------------------------------|--------------------------------------------------|-------------|--------|-------|
| Get    | /graph                                      | Lists all graphs stored in this database.        | Yes         | Yes    | Yes   |
| Post   | /graph                                      | Create a graph                                   | Yes         | Yes    | Yes   |
| Delete | /graph/{graph}                              | Deletes an existing graph object by name.        | Yes         | Yes    | Yes   |
| Get    | graph/{graph}                               | Get a graph                                      | Yes         | Yes    | Yes   |
| Get    | graph/{graph}/edge                          | Lists all edge collections within this graph.    | Yes         | Yes    | Yes   |
| Post   | /graph/{graph}/edge                         | Adds an additional edge definition to the graph. | Yes         | Yes    | Yes   |
| Post   | /graph/{graph}/edge/{edgeCollection}        | Creates a new edge in the collection.            | Yes         | Yes    | Yes   |
| Put    | /graph/{graph}/edge/{edgeCollection}        | Replace an edge definition                       | Yes         | Yes    | Yes   |
| Delete | /graph/{graph}/edge/{edgeCollection}        | Remove an edge definition                        | Yes         | Yes    | Yes   |
| Delete | /graph/{graph}/edge/{collection}/{edge}     | Remove an edge                                   | Yes         | Yes    | Yes   |
| Get    | /graph/{graph}/edge/{collection}/{edge}     | Get an Edge                                      | Yes         | Yes    | Yes   |
| Patch  | /graph/{graph}/edge/{collection}/{edge}     | Modify an edge                                   | Yes         | Yes    | Yes   |
| Put    | /graph/{graph}/edge/{collection}/{edge}     | Replace an edge                                  | Yes         | Yes    | Yes   |
| Get    | /graph/{graph}/vertex                       | List vertex collections                          | Yes         | Yes    | Yes   |
| Post   | /graph/{graph}/vertex                       | Add vertex collection                            | Yes         | Yes    | Yes   |
| Delete | /graph/{graph}/vertex/{collection}          | Remove vertex collection                         | Yes         | Yes    | Yes   |
| Post   | /graph/{graph}/vertex/{collection}          | Create a vertex                                  | Yes         | Yes    | Yes   |
| Delete | /graph/{graph}/vertex/{collection}/{vertex} | Remove a vertex                                  | Yes         | Yes    | Yes   |
| Get    | /graph/{graph}/vertex/{collection}/{vertex} | Get a vertex                                     | Yes         | Yes    | Yes   |
| patch  | /graph/{graph}/vertex/{collection}/{vertex} | Update a vertex                                  | Yes         | Yes    | Yes   |
| Put    | /graph/{graph}/vertex/{collection}/{vertex} | Replace a vertex                                 | Yes         | Yes    | Yes   |
| Get    | /edges/{collection-id}                      | Read in- or outbound edges                       | Yes         | Yes    | Yes   |
| Post   | /graphs/traversal                           | executes a traversal                             | No          | No     | No[1] |

[1] use query instead 


## Index API

| Http   | Endpoint                        | Function                          | Implemented | Tested | Done  |
|--------|---------------------------------|-----------------------------------|-------------|--------|-------|
| Get    | /index                          | Read all indexes of a collection  | Yes         | Yes    | Yes   |
| Get    | /index/{collection}/{indexName} | Read one index of a collection    | Yes         | Yes    | Yes   |
| Post   | /index/fulltext                 | Creates fulltext index            | Yes         | Yes    | Yes   |
| Post   | /index/general                  | Creates a new index in collection | No          | No     | No[1] |
| Post   | /index/geo                      | Creates geo index                 | Yes         | Yes    | Yes   |
| Post   | /index/hash                     | Creates hash index                | Yes         | Yes    | Yes   |
| Post   | /index/persistent               | Creates a persistent index        | Yes         | Yes    | Yes   |
| Post   | /index/skiplist                 | Creates skip list index           | Yes         | Yes    | Yes   |
| Post   | /index/ttl                      | Creates TTl index                 | Yes         | Yes    | Yes   |
| Delete | /index/{collection}/{indexName} | Deletes index                     | Yes         | Yes    | Yes   |

[1] Non fixable error. Possible legacy endpoint that wasn't removed. 

## KV API 

| Http   | Endpoint                | Function                                                 | Implemented | Tested | Done |
|--------|-------------------------|----------------------------------------------------------|-------------|--------|------|
| Get    | Count                   | Get number of key-value pairs in collection.             | Yes         | Yes    | Yes  |
| Get    | /kv/collection/keys     | Get all keys from key-value collection.                  | Yes         | Yes    | Yes  |
| Get    | /kv                     | Lists all key-value collections.                         | Yes         | Yes    | Yes  |
| Post   | /kv/collection          | Create key-value collection.                             | Yes         | Yes    | Yes  |
| Delete | /kv/collection          | Remove key-value collection.                             | Yes         | Yes    | Yes  |
| Get    | /kv/collection/value    | Get value from key-value collection.                     | Yes         | Yes    | Yes  |
| Delete | /kv/collection/value    | Remove key-value pair                                    | Yes         | Yes    | Yes  |
| Post   | /kv/collection/values   | Get some or all key-value pairs from collection.         | Yes         | Yes    | Yes  |
| Delete | /kv/collection/values   | Remove key-value pairs.                                  | Yes         | Yes    | Yes  |
| Put    | /kv/collection/value    | Set one or more key-value pairs in key-value collection. | Yes         | Yes    | Yes  |
| Put    | /kv/collection/truncate | Remove all key-value pairs in a collection.              | Yes         | Yes    | Yes  |


## Query API

Query API was aggregated in one singele "Query" method that combined all the bits and pieces of the API.

Note, query is stateful and as such not thread safe. 

Tables generated with [tablesgenerator.com](https://www.tablesgenerator.com/markdown_tables#)