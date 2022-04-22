package main

import (
	"github.com/marvin-hansen/goC8"
)

const (
	verbose          = true
	fabric           = "SouthEastAsia"
	graph            = "airline"
	collectionID     = "cities"
	edgeCollectionID = "flights"
)

func main() {
	c := goC8.NewClient(nil)
	setup(c)

	// query graph here

	//teardown(c)

	//setupGraph(c)
}
