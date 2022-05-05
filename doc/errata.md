# Corrections 

## Graph API 

### Add vertex collection
https://macrometa.com/docs/api#/operations/AddVertexCollection

Missing payload specification.

Please add: 

JSON Body: 

{'collection': name}

See 
https://github.com/Macrometacorp/pyC8/blob/c18fa188e1e19168e6de3f50e2b5a5598ea1505d/c8/graph.py#L152