package main

import (
	"github.com/marvin-hansen/goC8"
	"github.com/marvin-hansen/goC8/src/utils"
)

func teardown(c *goC8.Client) {
	exists, err := c.CheckGraphExists(fabric, graph)
	utils.CheckError(err, "Error CheckGraphExists: ")
	if exists {
		// if so delete graph
		_, createGraphErr := c.DeleteGraph(fabric, graph, true)
		utils.CheckError(createGraphErr, "Error CreateGraph")
	}
}
