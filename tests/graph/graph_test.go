package graph

import (
	"github.com/marvin-hansen/goC8"
	"github.com/stretchr/testify/assert"
	"testing"
)

const verbose = true

func TestGetAllGraphs(t *testing.T) {
	c := goC8.NewClient(nil)
	fabric := "_system"

	res, err := c.GetAllGraphs(fabric)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	printRes(res, verbose)
}

func TestGetGraph(t *testing.T) {
	c := goC8.NewClient(nil)
	fabric := "_system"
	graphName := "lectureteacher"

	res, err := c.GetGraph(fabric, graphName)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	printRes(res, verbose)
}
