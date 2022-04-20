package collection

import (
	"github.com/marvin-hansen/goC8"
	"github.com/marvin-hansen/goC8/requests/collection_req"
	"github.com/stretchr/testify/assert"
	"testing"
)

const verbose = false

func TestGetAllCollections(t *testing.T) {
	c := goC8.NewClient(nil)
	fabric := "SouthEastAsia"

	res, err := c.GetAllCollections(fabric)
	assert.NoError(t, err)
	assert.NotNil(t, res)

	if verbose {
		println(res.String())
	}
}

func TestCreateCollection(t *testing.T) {
	c := goC8.NewClient(nil)
	fabric := "SouthEastAsia"
	collType := collection_req.DocumentCollectionType
	collName := "TestCollection"

	err := c.CreateNewCollection(fabric, collName, false, collType)
	assert.NoError(t, err)
}

func TestGetCollectionInfo(t *testing.T) {
	c := goC8.NewClient(nil)
	fabric := "SouthEastAsia"
	collName := "TestCollection"

	res, err := c.GetCollectionInfo(fabric, collName)
	assert.NoError(t, err)
	assert.NotNil(t, res)

	if verbose {
		println(res.String())
	}
}

func TestUpdateCollection(t *testing.T) {
	c := goC8.NewClient(nil)
	fabric := "SouthEastAsia"
	collName := "TestCollection"
	properties := &collection_req.UpdateOptions{
		// Note: except for waitForSync and hasStream, collection properties cannot be changed once a collection is created.
		HasStream:   true,
		WaitForSync: true,
	}

	res, err := c.UpdateCollectionProperties(fabric, collName, properties)
	assert.NoError(t, err)
	assert.NotNil(t, res)

	if verbose {
		println(res.String())
	}
}

func TestTruncateCollection(t *testing.T) {
	c := goC8.NewClient(nil)
	fabric := "SouthEastAsia"
	collName := "TestCollection"

	res, err := c.TruncateCollection(fabric, collName)
	assert.NoError(t, err)
	assert.NotNil(t, res)

	if verbose {
		println(res.String())
	}
}

func TestDeleteCollection(t *testing.T) {
	c := goC8.NewClient(nil)
	fabric := "SouthEastAsia"
	collName := "TestCollection"

	err := c.DeleteCollection(fabric, collName, false)
	assert.NoError(t, err)
}
