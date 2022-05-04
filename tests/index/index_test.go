package index

import (
	"github.com/marvin-hansen/goC8"
	"github.com/marvin-hansen/goC8/requests/collection_req"
	config "github.com/marvin-hansen/goC8/tests/conf"
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

// We need to store the index name when creating it to delete it later
var sharedIndexName string = ""

func TestSetup(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())

	// test if city collection exists
	exists, err := c.Collection.CheckCollectionExists(fabric, citiesCollection)
	assert.NoError(t, err)
	if !exists {
		// if not create collection
		err = c.Collection.CreateNewCollection(fabric, citiesCollection, false, collType)
		assert.NoError(t, err)
	}

	// test if text collection exists
	exists, err = c.Collection.CheckCollectionExists(fabric, textCollection)
	assert.NoError(t, err)
	if !exists {
		// if not create collection
		err = c.Collection.CreateNewCollection(fabric, textCollection, false, collType)
		assert.NoError(t, err)
	}
}

func TestGetIndexes(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	res, err := c.Index.GetIndexes(fabric, citiesCollection)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	utils.PrintRes(res, verbose)
}

func TestCreateFulltextIndex(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	field := "Summary"
	minLength := 3

	res, err := c.Index.CreateFulltextIndex(fabric, textCollection, field, minLength)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	utils.PrintRes(res, verbose)

	sharedIndexName = res.Name
}

func TestCreateGeoIndex(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	field := "location"
	geoJson := true

	res, err := c.Index.CreateGeoIndex(fabric, citiesCollection, field, geoJson)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	utils.PrintRes(res, verbose)
}

func TestCreateHashIndex(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	field := "Text"
	deduplicate := true
	sparse := true
	unique := true

	res, err := c.Index.CreateHashIndex(fabric, textCollection, field, deduplicate, sparse, unique)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	utils.PrintRes(res, verbose)
}

func TestCreatePersistentIndex(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	field := "Keywords"
	deduplicate := true
	sparse := true
	unique := true

	res, err := c.Index.CreatePersistentIndex(fabric, textCollection, field, deduplicate, sparse, unique)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	utils.PrintRes(res, verbose)
}

func TestCreateSkipListIndex(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	field := "SkipText"
	deduplicate := true
	sparse := false
	unique := true

	res, err := c.Index.CreateSkipListIndex(fabric, textCollection, field, deduplicate, sparse, unique)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	utils.PrintRes(res, verbose)
}

func TestCreateTTLIndex(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	field := "requests"
	expiration := 10

	res, err := c.Index.CreateTTLIndex(fabric, textCollection, field, expiration)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	utils.PrintRes(res, verbose)
}

func TestDeleteIndex(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	indexName := sharedIndexName

	res, err := c.Index.DeleteIndex(fabric, textCollection, indexName)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	utils.PrintRes(res, verbose)
}

func TestTeardown(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())

	// test if graph exists
	exists, err := c.Graph.CheckGraphExists(fabric, graph)
	assert.NoError(t, err)
	if exists {
		// if so delete  graph
		_, delErr := c.Graph.DeleteGraph(fabric, graph, true)
		assert.NoError(t, delErr)
	}

	existsCol, errCol := c.Collection.CheckCollectionExists(fabric, textCollection)
	assert.NoError(t, errCol)

	if existsCol {
		errDel := c.Collection.DeleteCollection(fabric, textCollection, false)
		assert.NoError(t, errDel)
	}

	// It's possible that the graph has been deleted without dropCollections,
	// but the underlying collections still exists. Let's check one by one.

	// test if city collection exists
	exists, err = c.Collection.CheckCollectionExists(fabric, citiesCollection)
	assert.NoError(t, err)
	if exists {
		// if so, delete
		err = c.Collection.DeleteCollection(fabric, citiesCollection, false)
		assert.NoError(t, err)
	}

}
