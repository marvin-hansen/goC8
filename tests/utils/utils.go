package utils

import "github.com/marvin-hansen/goC8"

func PrintRes(res goC8.Responder, verbose bool) {
	if verbose {
		println(res.String())
	}
}

func PrintJsonRes(res goC8.JsonResponder, verbose bool) {
	if verbose {
		println(res.String())
	}
}
