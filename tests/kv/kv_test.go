package kv

import (
	"github.com/marvin-hansen/goC8"
	"github.com/marvin-hansen/goC8/requests/kv_req"
	config "github.com/marvin-hansen/goC8/tests/conf"
	"github.com/marvin-hansen/goC8/utils"
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
	res, err := c.KV.CreateNewKVCollection(fabric, collectionName, expiration, nil)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	utils.PrintRes(res, verbose)
}

func TestGetAllKVCollections(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	res, err := c.KV.GetAllKVCollections(fabric)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	utils.PrintRes(res, verbose)
}

func TestSetKeyValuePairs(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())

	kvPair1 := kv_req.NewKVPair("key1", "value1", -1)
	kvPair2 := kv_req.NewKVPair("key2", "value2", -1)
	kvPair3 := kv_req.NewKVPair("key3", "value3", -1)
	kvPair4 := kv_req.NewKVPair("key4", "value4", -1)
	kvPair5 := kv_req.NewKVPair("key5", "value5", -1)

	kvCollection := kv_req.NewKVPairCollection(*kvPair1, *kvPair2, *kvPair3, *kvPair4, *kvPair5)

	res, err := c.KV.SetKeyValuePairs(fabric, collectionName, *kvCollection)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	utils.PrintRes(res, verbose)
}

func TestGetAllKeys(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	offset := 0
	limit := 20
	order := kv_req.Ascending

	res, err := c.KV.GetAllKeys(fabric, collectionName, offset, limit, order)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	utils.PrintRes(res, verbose)
}

func TestGetAllValues(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	offset := 0
	limit := 20

	res, err := c.KV.GetAllValues(fabric, collectionName, offset, limit, nil)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	utils.PrintRes(res, verbose)
}

func TestGetValue(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	key := "key1"
	res, err := c.KV.GetValue(fabric, collectionName, key)

	expected := "value1"
	actual := res.Value
	assert.Equal(t, expected, actual, "Should be equal")
	assert.NoError(t, err)
	assert.NotNil(t, res)
	utils.PrintRes(res, verbose)
}

func TestCountKVCollections(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	res, err := c.KV.CountKVCollection(fabric, collectionName)

	expected := 5
	actual := res.Count
	assert.Equal(t, expected, actual, "Should be equal")
	assert.NoError(t, err)
	assert.NotNil(t, res)
	utils.PrintRes(res, verbose)
}

func TestDeleteValue(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	key := "key5"
	res, err := c.KV.DeleteValue(fabric, collectionName, key)

	expected := key
	actual := res.Key
	assert.Equal(t, expected, actual, "Should be equal")
	assert.NoError(t, err)
	assert.NotNil(t, res)

	resC, errC := c.KV.CountKVCollection(fabric, collectionName)
	assert.NoError(t, errC)

	expectedCount := 4
	actualCount := resC.Count
	assert.Equal(t, expectedCount, actualCount, "Should be equal")

	utils.PrintRes(res, verbose)
}

func TestDeleteKeyValuePairs(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	keyPairs := kv_req.KeyCollection{"key1", "key2", "key3"}

	res, err := c.KV.DeleteKeyValuePairs(fabric, collectionName, keyPairs)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	utils.PrintRes(res, verbose)
}

func TestTruncateKVCollection(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	res, err := c.KV.TruncateKVCollection(fabric, collectionName)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	utils.PrintRes(res, verbose)
}

func TestDeleteKVCollection(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	res, err := c.KV.DeleteKVCollection(fabric, collectionName)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	utils.PrintRes(res, verbose)
}
