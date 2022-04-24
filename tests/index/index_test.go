package index

import (
	"github.com/marvin-hansen/goC8"
	"github.com/marvin-hansen/goC8/requests/collection_req"
	"github.com/marvin-hansen/goC8/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	verbose          = true
	fabric           = "SouthEastAsia"
	graph            = "airline"
	citiesCollection = "cities"
	textCollection   = "textCollection"
	collType         = collection_req.DocumentCollectionType
)

func TestSetup(t *testing.T) {
	c := goC8.NewClient(nil)

	// test if city collection exists
	exists, err := c.CheckCollectionExists(fabric, citiesCollection)
	assert.NoError(t, err)
	if !exists {
		// if not create collection
		err = c.CreateNewCollection(fabric, citiesCollection, false, collType)
		assert.NoError(t, err)
	}

	// test if text collection exists
	exists, err = c.CheckCollectionExists(fabric, textCollection)
	assert.NoError(t, err)
	if !exists {
		// if not create collection
		err = c.CreateNewCollection(fabric, textCollection, false, collType)
		assert.NoError(t, err)
	}
}

func TestGetIndexes(t *testing.T) {
	c := goC8.NewClient(nil)
	res, err := c.GetIndexes(fabric, citiesCollection)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	utils.PrintRes(res, verbose)
}

func TestCreateFulltextIndex(t *testing.T) {
	c := goC8.NewClient(nil)
	field := "Summary"
	minLength := 3

	res, err := c.CreateFulltextIndex(fabric, textCollection, field, minLength)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	utils.PrintRes(res, verbose)
}

func TestCreateGeoIndex(t *testing.T) {
	c := goC8.NewClient(nil)
	field := "location"
	geoJson := true

	res, err := c.CreateGeoIndex(fabric, citiesCollection, field, geoJson)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	utils.PrintRes(res, verbose)
}

func TestCreateHashIndex(t *testing.T) {
	//c := goC8.NewClient(nil)
	//field := "Text"
	//
	//res, err := c.CreateHashIndex()
	//assert.NoError(t, err)
	//assert.NotNil(t, res)
	//utils.PrintRes(res, verbose)
}

func TestDeleteIndex(t *testing.T) {
	c := goC8.NewClient(nil)
	indexName := ""

	res, err := c.DeleteIndex(fabric, textCollection, indexName)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	utils.PrintRes(res, verbose)
}

func TestTeardown(t *testing.T) {
	c := goC8.NewClient(nil)

	// test if graph exists
	exists, err := c.CheckGraphExists(fabric, graph)
	assert.NoError(t, err)
	if !exists {
		// if so delete  graph
		_, delErr := c.DeleteGraph(fabric, graph, true)
		assert.NoError(t, delErr)
	}
}
