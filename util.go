package goC8

import (
	"github.com/marvin-hansen/goC8/requests/query_req"
	"log"
	"time"
)

func CheckErrorLog(err error, msg string) {
	if err != nil {
		log.Println("error: " + err.Error())
		log.Fatalf(msg)
	}
}

func CheckError(err error) error {
	if err != nil {
		return err
	} else {
		return nil
	}
}

// TimeTrack tracks execution time.
// put in the first line of method to track:
// 	defer TimeTrack(time.Now(), "methodName")
func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
	println()
}

func PrintQuery(res *query_req.Cursor, verbose bool) {
	if verbose {
		println(res.Result.String())
	}
}

func PrintRes(res Responder, verbose bool) {
	if verbose {
		println(res.String())
	}
}

func PrintJsonRes(res JsonResponder, verbose bool) {
	if verbose {
		println(res.String())
	}
}
