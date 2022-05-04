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

| http   | Endpoint                                 | Function                                           | Implemented | Tested | Done |
|--------|------------------------------------------|----------------------------------------------------|-------------|--------|------|
| Get    | /collection                              | Fetches the list of all collections.               | Yes         | Yes    | Yes  |
| Post   | /collection                              | Create collection.                                 | Yes         | Yes    | Yes  |
| Get    | /collection/{collection-name}            | Fetches the information about collection.          | Yes         | Yes    | Yes  |
| Delete | /collection/{collection-name}            | Remove the collection.                             | Yes         | Yes    | Yes  |
| Get    | /collection/{collection-name}/count      | Returns the number of documents in the collection. | Yes         | Yes    | Yes  |
| Put    | /collection/{collection-name}/truncate   | Truncate collection                                | Yes         | Yes    | Yes  |
| Put    | /collection/{collection-name}/properties | Changes the properties of a collection.            | Yes         | Yes    | Yes  |


## Graph API
| http   | Endpoint                                    | Function                                         | Implemented | Tested | Done |
|--------|---------------------------------------------|--------------------------------------------------|-------------|--------|------|
| Get    | /graph                                      | Lists all graphs stored in this database.        | Yes         | Yes    | Yes  |
| Post   | /graph                                      | Create a graph                                   | Yes         | Yes    | Yes  |
| Delete | /graph/{graph}                              | Deletes an existing graph object by name.        | Yes         | Yes    | Yes  |
| Get    | graph/{graph}                               | Get a graph                                      | Yes         | Yes    | Yes  |
| Get    | graph/{graph}/edge                          | Lists all edge collections within this graph.    | Yes         | Yes    | Yes  |
| Post   | /graph/{graph}/edge                         | Adds an additional edge definition to the graph. | Yes         | Yes    | Yes  |
| Post   | /graph/{graph}/edge/{edgeCollection}        | Creates a new edge in the collection.            | Yes         | Yes    | Yes  |
| Put    | /graph/{graph}/edge/{edgeCollection}        | Replace an edge definition                       | Yes         | Yes    | Yes  |
| Delete | /graph/{graph}/edge/{edgeCollection}        | Remove an edge definition                        | Yes         | Yes    | Yes  |
| Delete | /graph/{graph}/edge/{collection}/{edge}     | Remove an edge                                   | Yes         | Yes    | Yes  |
| Get    | /graph/{graph}/edge/{collection}/{edge}     | Get an Edge                                      | Yes         | Yes    | Yes  |
| Patch  | /graph/{graph}/edge/{collection}/{edge}     | Modify an edge                                   |             |        |      |
| Put    | /graph/{graph}/edge/{collection}/{edge}     | Replace an edge                                  | Yes         | Yes    | Yes  |
| Get    | /graph/{graph}/vertex                       | List vertex collections                          | Yes         | Yes    | Yes  |
| Post   | /graph/{graph}/vertex                       | Add vertex collection                            |             |        |      |
| Delete | /graph/{graph}/vertex/{collection}          | Remove vertex collection                         |             |        |      |
| Post   | /graph/{graph}/vertex/{collection}          | Create a vertex                                  |             |        |      |
| Delete | /graph/{graph}/vertex/{collection}/{vertex} | Remove a vertex                                  |             |        |      |
| Get    | /graph/{graph}/vertex/{collection}/{vertex} | Get a vertex                                     | Yes         | Yes    | Yes  |
| patch  | /graph/{graph}/vertex/{collection}/{vertex} | Update a vertex                                  |             |        |      |
| Put    | /graph/{graph}/vertex/{collection}/{vertex} | Replace a vertex                                 |             |        |      |
| Get    | /edges/{collection-id}                      | Read in- or outbound edges                       |             |        |      |
| Post   | /graphs/traversal                           | executes a traversal                             |             |        |      |

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



Tables generated with [tablesgenerator.com](https://www.tablesgenerator.com/markdown_tables#)