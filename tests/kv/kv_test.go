package kv

import (
	"github.com/marvin-hansen/goC8"
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

func TestCountKVCollections(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	res, err := c.CountKVCollection(fabric, collectionName)
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
