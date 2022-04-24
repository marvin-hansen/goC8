package goC8

import (
	r "github.com/marvin-hansen/goC8/requests/query_req"
	"time"
)

func (c Client) Query(fabric, query string, bindVars map[string]interface{}, options *r.CursorOptions) (res *r.Cursor, err error) {
	if benchmark {
		defer TimeTrack(time.Now(), "Query")
	}

	if options == nil {
		options = r.NewDefaultCursorOptions()
	}
	ttl := c.getQueryTTL()
	req := r.NewRequestForCreateCursor(fabric, query, bindVars, options, ttl)

	// Create a cursor for the query onfrom the server
	response := r.NewResponseForCreateCursor()
	if err = c.request(req, response); err != nil {
		return nil, err
	}

	// We cast into a cursor here to allow for updates from NextCursor
	res = r.NewCursorFromCreateCursor(response)

	// check for more to come
	if response.HasMore {
		for {
			// request update for the cursor
			reqNext := r.NewRequestForReadNextCursor(fabric, res.Id)
			responseNext := r.NewResponseForReadNextCursor()
			if err = c.request(reqNext, responseNext); err != nil {
				return nil, err
			}

			// updated cursor with next result
			res.Update(responseNext)

			// stop when there are no more results, delete cursor and stop
			if responseNext.HasMore == false {

				// Delete the cursor from the server
				reqDel := r.NewRequestForDeleteCursor(fabric, response.Id)
				resDel := r.NewResponseForDeleteCursor()
				if err = c.request(reqDel, resDel); err != nil {
					return nil, err
				}

				// stop the loop
				break
			}
		}
	}

	return res, nil
}
