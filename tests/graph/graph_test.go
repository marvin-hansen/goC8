package graph

import (
	"github.com/marvin-hansen/goC8"
	config "github.com/marvin-hansen/goC8/tests/conf"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	verbose   = true
	fabric    = "_system"
	graphName = "lectureteacher"
)

func TestGetAllGraphs(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())

	res, err := c.GetAllGraphs(fabric)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	goC8.PrintRes(res, verbose)
}

func TestCreateGraph(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	graphDef := getGraphDefinition()
	res, err := c.CreateGraph(fabric, graphDef)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	goC8.PrintRes(res, verbose)
}

func TestGetGraph(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())

	res, err := c.GetGraph(fabric, graphName)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	goC8.PrintRes(res, verbose)
}

func TestCheckGraphExists(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())

	exists, err := c.CheckGraphExists(fabric, graphName)
	assert.NoError(t, err)
	expected := true
	actual := exists
	assert.Equal(t, expected, actual, "Should exists")
	noneExistingGraphName := "noneExistingGraphName"
	exists, err = c.CheckGraphExists(fabric, noneExistingGraphName)
	assert.NoError(t, err)

	expected = false
	actual = exists
	assert.Equal(t, expected, actual, "Should exists")
}

func TestDeleteGraph(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	dropCollections := false

	res, err := c.DeleteGraph(fabric, graphName, dropCollections)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	goC8.PrintRes(res, verbose)
}
