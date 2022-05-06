package goC8

import (
	"github.com/marvin-hansen/goC8/requests/query_req"
	"github.com/marvin-hansen/goC8/utils"
	"time"
)

type QueryManager struct {
	client *Client
}

func NewQueryManagerManager(client *Client) *QueryManager {
	return &QueryManager{client: client}
}

func (c QueryManager) ExplainQuery(fabric, query string) (res *query_req.ResponseForExplainQuery, err error) {
	if benchmark {
		defer utils.TimeTrack(time.Now(), "ValidateQuery")
	}

	req := query_req.NewRequestForExplainQuery(fabric, query)
	res = query_req.NewResponseForExplainQuery()
	err = c.client.Request(req, res)
	return res, CheckReturnError(err)
}

// ValidateQuery
// validates a C8QL query. To actually query the database, see Query
// https://macrometa.com/docs/api#/operations/parseQuery
func (c QueryManager) ValidateQuery(fabric, query string) (res *query_req.ResponseForValidateQuery, err error) {
	if benchmark {
		defer utils.TimeTrack(time.Now(), "ValidateQuery")
	}

	req := query_req.NewRequestForValidateQuery(fabric, query)
	res = query_req.NewResponseForValidateQuery()
	err = c.client.Request(req, res)
	return res, CheckReturnError(err)
}

// Query
// Queries database / graph
func (c QueryManager) Query(fabric, query string, bindVars map[string]interface{}, options *query_req.CursorOptions) (res *query_req.Cursor, err error) {
	if benchmark {
		defer utils.TimeTrack(time.Now(), "Query")
	}

	if options == nil {
		options = query_req.NewDefaultCursorOptions()
	}
	ttl := c.client.getQueryTTL()
	req := query_req.NewRequestForCreateCursor(fabric, query, bindVars, options, ttl)

	// Create a cursor for the query onfrom the server
	response := query_req.NewResponseForCreateCursor()
	if err = c.client.Request(req, response); err != nil {
		return nil, err
	}

	// We cast into a cursor here to allow for updates from NextCursor
	res = query_req.NewCursorFromCreateCursor(response)

	// check for more to come
	if response.HasMore {
		for {
			// request update for the cursor
			reqNext := query_req.NewRequestForReadNextCursor(fabric, res.Id)
			responseNext := query_req.NewResponseForReadNextCursor()
			if err = c.client.Request(reqNext, responseNext); err != nil {
				return nil, err
			}

			// updated cursor with next result
			res.Update(responseNext)

			// stop when there are no more results, delete cursor and stop
			if responseNext.HasMore == false {

				// Delete the cursor from the server
				reqDel := query_req.NewRequestForDeleteCursor(fabric, response.Id)
				resDel := query_req.NewResponseForDeleteCursor()
				if err = c.client.Request(reqDel, resDel); err != nil {
					return nil, err
				}

				// stop the loop
				break
			}
		}
	}

	return res, nil
}
