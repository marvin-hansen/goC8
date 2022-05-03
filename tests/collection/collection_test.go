package collection

import (
	"github.com/marvin-hansen/goC8"
	collection_req2 "github.com/marvin-hansen/goC8/src/requests/collection_req"
	"github.com/marvin-hansen/goC8/src/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	verbose  = true
	fabric   = "SouthEastAsia"
	collType = collection_req2.DocumentCollectionType
	collName = "TestCollection"
)

func TestGetAllCollections(t *testing.T) {
	c := goC8.NewClient(nil)
	res, err := c.GetAllCollections(fabric)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	utils.PrintRes(res, verbose)

}

func TestCreateCollection(t *testing.T) {
	c := goC8.NewClient(nil)
	err := c.CreateNewCollection(fabric, collName, false, collType)
	assert.NoError(t, err)
}

func TestCheckCollectionExists(t *testing.T) {
	c := goC8.NewClient(nil)
	exists, err := c.CheckCollectionExists(fabric, collName)
	assert.NoError(t, err)
	expected := true
	actual := exists
	assert.Equal(t, expected, actual, "Should be equal")

	noCollName := "DoesNotExistsCollection"
	exists, err = c.CheckCollectionExists(fabric, noCollName)
	assert.NoError(t, err)
	expected = false
	actual = exists
	assert.Equal(t, expected, actual, "Should be equal")

}

func TestGetCollectionInfo(t *testing.T) {
	c := goC8.NewClient(nil)
	res, err := c.GetCollectionInfo(fabric, collName)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	utils.PrintRes(res, verbose)
}

func TestUpdateCollection(t *testing.T) {
	c := goC8.NewClient(nil)
	properties := &collection_req2.UpdateOptions{
		// Note: except for waitForSync and hasStream, collection properties cannot be changed once a collection is created.
		HasStream:   true,
		WaitForSync: true,
	}

	res, err := c.UpdateCollectionProperties(fabric, collName, properties)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	utils.PrintRes(res, verbose)
}

func TestTruncateCollection(t *testing.T) {
	c := goC8.NewClient(nil)
	res, err := c.TruncateCollection(fabric, collName)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	utils.PrintRes(res, verbose)
}

func TestDeleteCollection(t *testing.T) {
	c := goC8.NewClient(nil)
	err := c.DeleteCollection(fabric, collName, false)
	assert.NoError(t, err)
}
