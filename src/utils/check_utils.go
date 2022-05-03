package utils

import "log"

func CheckReturnError(err error) error {
	if err != nil {
		log.Println("error: " + err.Error())
		return err
	}
	return nil
}

func CheckError(err error, msg string) {
	if err != nil {
		log.Println("error: " + err.Error())
		log.Fatalf(msg)
	}
}

func CheckErrorLog(err error, msg string) {
	if err != nil {
		log.Println("error: " + err.Error())
		log.Fatalf(msg)
	}
}
