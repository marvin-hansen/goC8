package utils

import (
	"log"
	"time"
)

// TimeTrack tracks execution time.
// put in the first line of method to track:
// 	defer TimeTrack(time.Now(), "methodName")
func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
	println()
}
