package goC8

func TeardownGraph(c *Client, fabric, graphName string, dropCollections bool) {
	exists, err := c.Graph.CheckGraphExists(fabric, graphName)
	CheckError(err, "Error CheckGraphExists: ")
	if exists {
		// if so delete graph
		_, createGraphErr := c.Graph.DeleteGraph(fabric, graphName, dropCollections)
		CheckError(createGraphErr, "Error CreateGraph")
	}
}

func TeardownCollection(c *Client, fabric, collectionName string) {
	exists, err := c.Collection.CheckCollectionExists(fabric, collectionName)
	CheckError(err, "Error CheckCollectionExists: ")
	if exists {
		// if so, delete collection
		err = c.Collection.DeleteCollection(fabric, collectionName, false)
		CheckError(err, "Error DeleteCollection")
	}
}
