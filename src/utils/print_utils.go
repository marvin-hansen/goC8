package utils

import (
	"github.com/marvin-hansen/goC8/src"
	"github.com/marvin-hansen/goC8/src/requests/query_req"
)

func PrintQuery(res *query_req.Cursor, verbose bool) {
	if verbose {
		println(res.Result.String())
	}
}

func PrintRes(res src.Responder, verbose bool) {
	if verbose {
		println(res.String())
	}
}

func PrintJsonRes(res src.JsonResponder, verbose bool) {
	if verbose {
		println(res.String())
	}
}
