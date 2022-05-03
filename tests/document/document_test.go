package document

import (
	"github.com/marvin-hansen/goC8"
	"github.com/marvin-hansen/goC8/src/requests/collection_req"
	"github.com/marvin-hansen/goC8/src/requests/document_req"
	"github.com/marvin-hansen/goC8/src/utils"
	config "github.com/marvin-hansen/goC8/tests/conf"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	verbose  = true
	fabric   = "SouthEastAsia"
	collName = "TestCollection"
)

// we need a shared variable here to store the key of inserted doc
// to test deletion of the same doc later.
var SharedKey string = ""

func TestSetup(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	collType := collection_req.DocumentCollectionType
	err := c.Collection.CreateNewCollection(fabric, collName, false, collType)
	assert.NoError(t, err)
}

func TestCreateNewDocument(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	jsonDocument := getTestInsertData()
	silent := false

	res, err := c.Document.CreateNewDocument(fabric, collName, silent, jsonDocument, nil)
	assert.NoError(t, err)

	if verbose {
		if res != nil {
			assert.NotNil(t, res)
			for _, v := range *res {
				println(v.String())
				SharedKey = v.Key
			}
		}
	}

	res, err = c.Document.CreateNewDocument(fabric, collName, silent, jsonDocument, nil)
	assert.NoError(t, err)

}

func TestGetDocument(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	key := SharedKey

	res, err := c.Document.GetDocument(fabric, collName, key)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	utils.PrintJsonRes(res, verbose)
}

func TestCheckDocumentExists(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	key := SharedKey

	exists, err := c.Document.CheckDocumentExists(fabric, collName, key)
	assert.NoError(t, err)
	expected := true
	actual := exists
	assert.Equal(t, expected, actual, "Should exists")

	key = "99"
	exists, err = c.Document.CheckDocumentExists(fabric, collName, key)
	assert.NoError(t, err)
	expected = false
	actual = exists
	assert.Equal(t, expected, actual, "Should not exist")
}

func TestUpdateDocument(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	key := SharedKey
	jsonDocument := getTestUpdateSingleData(key)
	var silent = true

	res, err := c.Document.UpdateDocument(fabric, collName, key, jsonDocument, silent, nil)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	//PrintRes(res, verbose) Nothing to print as we do silent update

	silent = false
	res, err = c.Document.UpdateDocument(fabric, collName, key, jsonDocument, silent, nil)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	utils.PrintRes(res, verbose)

}

func TestReplaceDocuments(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	key := SharedKey

	jsonDocument := getTestReplaceSingleData(key)

	res, err := c.Document.ReplaceDocument(fabric, collName, key, jsonDocument, nil)
	assert.NoError(t, err)
	utils.PrintRes(res, verbose)

}

func TestUpdateManyDocument(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())

	key := SharedKey
	jsonDocument := getTestUpdateManyData(key)

	res, err := c.Document.UpdateManyDocuments(fabric, collName, jsonDocument, nil)
	assert.NoError(t, err)
	assert.NotNil(t, res)

	if verbose {
		if res != nil {
			assert.NotNil(t, res)
			for _, v := range *res {
				println(v.String())
			}
		}
	}
}

func TestReplaceManyDocuments(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	k1 := SharedKey

	jsonDocument := getTestReplaceManyData(k1)

	res, err := c.Document.ReplaceManyDocuments(fabric, collName, jsonDocument, nil)
	assert.NoError(t, err)

	if verbose {
		if res != nil {
			assert.NotNil(t, res)
			for _, v := range *res {
				println(v.String())
			}
		}
	}
}

func TestDeleteDocument(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	key := SharedKey

	para := document_req.GetCustomDeleteDocumentParameters(false, false, false, false)
	res, err := c.Document.DeleteDocument(fabric, collName, key, para)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestTeardown(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	err := c.Collection.DeleteCollection(fabric, collName, false)
	assert.NoError(t, err)
}
