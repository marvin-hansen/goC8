package utils

import (
	"github.com/marvin-hansen/goC8/requests/query_req"
	"github.com/marvin-hansen/goC8/types"
)

func PrintQuery(res *query_req.Cursor, verbose bool) {
	if verbose {
		println(res.Result.String())
	}
}

func PrintRes(res types.Responder, verbose bool) {
	if verbose {
		println(res.String())
	}
}

func PrintJsonRes(res types.JsonResponder, verbose bool) {
	if verbose {
		println(res.String())
	}
}
