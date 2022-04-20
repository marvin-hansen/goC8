package graph

import "github.com/marvin-hansen/goC8"

func printRes(res goC8.Responder, verbose bool) {
	if verbose {
		println(res.String())
	}
}

func printJsonRes(res goC8.JsonResponder, verbose bool) {
	if verbose {
		println(res.String())
	}
}
