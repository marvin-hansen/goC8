package index

import "github.com/marvin-hansen/goC8"

func printRes(res goC8.Responder, verbose bool) {
	if verbose {
		println(res.String())
	}
}
