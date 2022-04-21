package index

import (
	"github.com/marvin-hansen/goC8"
	"github.com/marvin-hansen/goC8/requests/collection_req"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	verbose      = true
	fabric       = "_system"
	collectionID = "teachers"
)

func TestSetup(t *testing.T) {
	c := goC8.NewClient(nil)
	collType := collection_req.DocumentCollectionType
	collName := "TestCollection"

	// test if collection exists
	exists, err := c.CheckCollectionExists(fabric, collName)
	assert.NoError(t, err)

	if !exists {
		// if not create collection
		err := c.CreateNewCollection(fabric, collName, false, collType)
		assert.NoError(t, err)
		assert.NoError(t, err)
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

}
