package query_req

import "fmt"

func getQueryPayload(query string) []byte {
	return []byte(fmt.Sprintf(`{"query": "%v"}`, query))
}
