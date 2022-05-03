package utils

import (
	"github.com/marvin-hansen/goC8"
	"github.com/marvin-hansen/goC8/src/requests/query_req"
)

func PrintQuery(res *query_req.Cursor, verbose bool) {
	if verbose {
		println(res.Result.String())
	}
}

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
