package document

import (
	"github.com/stretchr/testify/assert"
	"scigraph/src/clients/gdn_client"
	"testing"
)

func TestUpdateManyDocument(t *testing.T) {
	c := gdn_client.NewClient(nil)
	fabric := "SouthEastAsia"
	collName := "TestCollection"

	key := "2"
	jsonDocument := getTestUpdateManyData(key)

	res, err := c.UpdateManyDocuments(fabric, collName, jsonDocument, nil)
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
	c := gdn_client.NewClient(nil)
	fabric := "SouthEastAsia"
	collName := "TestCollection"
	k1 := "1"
	k2 := "2"

	jsonDocument := getTestReplaceManyData(k1, k2)

	res, err := c.ReplaceManyDocuments(fabric, collName, jsonDocument, nil)
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

func TestDeleteManyDocuments(t *testing.T) {
	c := gdn_client.NewClient(nil)
	fabric := "SouthEastAsia"
	collName := "TestCollection"

	k1 := "1"
	k2 := "2"
	keysToDelete := getKeysToDelete(k1, k2)

	resDel, errDel := c.DeleteManyDocuments(fabric, collName, keysToDelete, nil)
	assert.NoError(t, errDel)
	assert.NotNil(t, resDel)
}
