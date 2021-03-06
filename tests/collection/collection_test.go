package collection

import (
	"fmt"
	"github.com/marvin-hansen/goC8"
	collection_req2 "github.com/marvin-hansen/goC8/requests/collection_req"
	config "github.com/marvin-hansen/goC8/tests/conf"
	"github.com/marvin-hansen/goC8/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	verbose  = true
	fabric   = "SouthEastAsia"
	collType = types.DocumentCollectionType
	collName = "TestCollection"
)

func TestGetAllCollections(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	res, err := c.Collection.GetAllCollections(fabric)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	goC8.PrintRes(res, verbose)
}

func TestCreateCollection(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	err := c.Collection.CreateNewCollection(fabric, collName, false, collType)
	assert.NoError(t, err)
}

func TestCheckCollectionExists(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	exists, err := c.Collection.CheckCollectionExists(fabric, collName)
	assert.NoError(t, err)
	expected := true
	actual := exists
	assert.Equal(t, expected, actual, "Should be equal")

	noCollName := "DoesNotExistsCollection"
	exists, err = c.Collection.CheckCollectionExists(fabric, noCollName)
	assert.NoError(t, err)
	expected = false
	actual = exists
	assert.Equal(t, expected, actual, "Should be equal")
}

func TestCountCollection(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	res, err := c.Collection.CountCollection(fabric, collName)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	//utils.PrintRes(res, verbose)
	println("Count: " + fmt.Sprint(res.Count))
}

func TestGetCollectionInfo(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	res, err := c.Collection.GetCollectionInfo(fabric, collName)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	goC8.PrintRes(res, verbose)
}

func TestUpdateCollection(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	properties := &collection_req2.UpdateOptions{
		// Note: except for waitForSync and hasStream, collection properties cannot be changed once a collection is created.
		HasStream:   true,
		WaitForSync: true,
	}

	res, err := c.Collection.UpdateCollectionProperties(fabric, collName, properties)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	goC8.PrintRes(res, verbose)
}

func TestTruncateCollection(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	res, err := c.Collection.TruncateCollection(fabric, collName)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	goC8.PrintRes(res, verbose)
}

func TestDeleteCollection(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	err := c.Collection.DeleteCollection(fabric, collName, false)
	assert.NoError(t, err)
}
