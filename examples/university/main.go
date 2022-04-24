package main

import (
	"github.com/marvin-hansen/goC8"
)

const (
	verbose              = true
	fabric               = "SouthEastAsia"
	collectionTeachers   = "teachers"
	collectionLectures   = "lectures"
	collectionTutorials  = "tutorials"
	edgeCollectionTeach  = "teach"
	edgeCollectionTutors = "tutors"
	graph                = "lectureteacher"
)

func main() {
	c := goC8.NewClient(nil)

	setupTeachers(c)
	//setup(c)
}
