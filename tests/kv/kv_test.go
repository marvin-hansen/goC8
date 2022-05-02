package kv

import (
	"github.com/marvin-hansen/goC8"
	"github.com/marvin-hansen/goC8/requests/kv_req"
	config "github.com/marvin-hansen/goC8/tests/conf"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	verbose        = true
	fabric         = "SouthEastAsia"
	collectionName = "TestKVCollection"
	expiration     = false
)

func TestCreateKVCollection(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	res, err := c.CreateNewKVCollection(fabric, collectionName, expiration, nil)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	goC8.PrintRes(res, verbose)
}

func TestGetAllKVCollections(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	res, err := c.GetAllKVCollections(fabric)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	goC8.PrintRes(res, verbose)
}

func TestSetKeyValuePairs(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())

	kvPair1 := kv_req.NewKVPair("key1", "value1", -1)
	kvPair2 := kv_req.NewKVPair("key2", "value2", -1)
	kvPair3 := kv_req.NewKVPair("key3", "value3", -1)
	kvCollection := kv_req.NewKVPairCollection(*kvPair1, *kvPair2, *kvPair3)

	res, err := c.SetKeyValuePairs(fabric, collectionName, *kvCollection)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	goC8.PrintRes(res, verbose)
}

func TestCountKVCollections(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	res, err := c.CountKVCollection(fabric, collectionName)

	expected := 3
	actual := res.Count
	assert.Equal(t, expected, actual, "Should be equal")
	assert.NoError(t, err)
	assert.NotNil(t, res)
	goC8.PrintRes(res, verbose)
}

func TestTruncateKVCollection(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	res, err := c.TruncateKVCollection(fabric, collectionName)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	goC8.PrintRes(res, verbose)
}

func TestDeleteKVCollection(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	res, err := c.DeleteKVCollection(fabric, collectionName)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	goC8.PrintRes(res, verbose)
}
