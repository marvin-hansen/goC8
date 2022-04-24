package goC8

import (
	"github.com/marvin-hansen/goC8/requests/query_req"
)

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
