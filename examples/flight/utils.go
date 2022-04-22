package main

import "log"

func checkError(err error, msg string) {
	if err != nil {
		log.Println("error: " + err.Error())
		log.Fatalf(msg)
	}
}

func dbgPrint(msg string) {
	if verbose {
		println(msg)
	}
}
