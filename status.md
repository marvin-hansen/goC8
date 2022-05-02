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