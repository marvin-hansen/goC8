package index

import (
	"github.com/marvin-hansen/goC8"
	"github.com/marvin-hansen/goC8/examples/sample_data"
	"github.com/marvin-hansen/goC8/requests/collection_req"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	verbose          = true
	fabric           = "_system"
	graph            = "airline"
	collectionID     = "cities"
	edgecollectionID = "flights"
)

func TestSetup(t *testing.T) {
	c := goC8.NewClient(nil)

	// test if city collection exists
	exists, err := c.CheckCollectionExists(fabric, collectionID)
	assert.NoError(t, err)
	if !exists {
		// if not create collection
		collType := collection_req.DocumentCollectionType
		err = c.CreateNewCollection(fabric, collectionID, false, collType)
		assert.NoError(t, err)

		// import city data
		silent := false
		jsonDocument := sample_data.GetCityData()
		_, err = c.CreateNewDocument(fabric, collectionID, silent, jsonDocument, nil)
		assert.NoError(t, err)
	}

	// test if edge collection exists
	exists, err = c.CheckCollectionExists(fabric, edgecollectionID)
	assert.NoError(t, err)
	if !exists {
		// if not create edge collection
		collType := collection_req.EdgeCollectionType
		err = c.CreateNewCollection(fabric, edgecollectionID, false, collType)
		assert.NoError(t, err)

		// import flight data
		silent := false
		jsonDocument := sample_data.GetFlightData()
		_, err = c.CreateNewDocument(fabric, edgecollectionID, silent, jsonDocument, nil)
		assert.NoError(t, err)
	}

	// test if graph exists
	exists, err = c.CheckGraphExists(fabric, graph)
	assert.NoError(t, err)
	if !exists {
		// if not create graph
		// TODO
	}

}

func TestGetIndexes(t *testing.T) {
	c := goC8.NewClient(nil)

	res, err := c.GetIndexes(fabric, collectionID)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	printRes(res, verbose)
}

func TestTeardown(t *testing.T) {

	c := goC8.NewClient(nil)

	// test if city collection exists
	exists, err := c.CheckCollectionExists(fabric, collectionID)
	assert.NoError(t, err)
	if exists {
		// if so, delete
		err = c.DeleteCollection(fabric, collectionID, false)
		assert.NoError(t, err)
	}

	// test if edge collection exists
	exists, err = c.CheckCollectionExists(fabric, edgecollectionID)
	assert.NoError(t, err)
	if exists {
		// if so, delete
		err = c.DeleteCollection(fabric, edgecollectionID, false)
		assert.NoError(t, err)
	}

	// test if graph exists
	exists, err = c.CheckGraphExists(fabric, graph)
	assert.NoError(t, err)
	if !exists {
		// if so delete  graph
		// TODO
	}

}
