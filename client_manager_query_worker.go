package goC8

import (
	"github.com/marvin-hansen/goC8/requests/query_req"
	"github.com/marvin-hansen/goC8/requests/qw_req"
	"github.com/marvin-hansen/goC8/utils"
	"time"
)

type QueryWorkerManager struct {
	client *Client
}

func NewQueryWorkerManager(client *Client) *QueryWorkerManager {
	return &QueryWorkerManager{client: client}
}

// CreateQueryWorker
// Save a query for a user for a fabric.
// bindVars: bindVars for the query
// name: Name for the query
// value: value of the query
// https://macrometa.com/docs/api#/operations/SaveRestqlByName
func (c QueryWorkerManager) CreateQueryWorker(fabric, workerName, queryString, bindVars string) (res *qw_req.ResponseForQueryWorker, err error) {
	if benchmark {
		defer utils.TimeTrack(time.Now(), "CreateQueryWorker")
	}

	req := qw_req.NewRequestForCreateQueryWorker(fabric, workerName, queryString, bindVars)
	res = qw_req.NewResponseForQueryWorker()
	err = c.client.Request(req, res)
	return res, CheckReturnError(err)
}

// RunQueryWorker
// Run a saved query for a given fabric. If there are more that 100 records, the hasMore flag is set to true.
// Note this client fetches all additional records, merges them, and returns just one combined result.
// https://macrometa.com/docs/api#/operations/ExecuteRestqlByName
func (c QueryWorkerManager) RunQueryWorker(fabric, workerName, bindVars string) (res *query_req.Cursor, err error) {
	if benchmark {
		defer utils.TimeTrack(time.Now(), "RunQueryWorker")
	}

	req := qw_req.NewRequestForRunQueryWorker(fabric, workerName, bindVars)
	res = qw_req.NewResponseForRunQueryWorker()
	if err = c.client.requestJsonResponse(req, res); err != nil {
		return nil, err
	}

	if res.HasMore {

		for {
			// request update for the cursor
			reqNext := qw_req.NewRequestForReadNextCursor(fabric, res.Id)
			responseNext := qw_req.NewResponseForReadNextCursor()
			if err = c.client.requestJsonResponse(reqNext, responseNext); err != nil {
				return nil, err
			}

			// updated cursor with next result
			res.Update(responseNext)

			if responseNext.HasMore == false {
				break
			}
		}

	}

	return res, CheckReturnError(err)
}

// ReadAllQueryWorkers
// Get list of saved queries for the fabric.
// https://macrometa.com/docs/api#/operations/ListRestqlAssociatedWithCurrentUser
func (c QueryWorkerManager) ReadAllQueryWorkers(fabric string) (res *qw_req.ResponseForReadAllQueryWorkers, err error) {
	if benchmark {
		defer utils.TimeTrack(time.Now(), "ReadAllQueryWorkers")
	}

	req := qw_req.NewRequestForReadAllQueryWorkers(fabric)
	res = qw_req.NewResponseForReadAllQueryWorkers()
	err = c.client.Request(req, res)

	return res, CheckReturnError(err)
}

// UpdateQueryWorker
// Update a saved query for a fabric.
// https://macrometa.com/docs/api#/operations/UpdateRestqlByName
func (c QueryWorkerManager) UpdateQueryWorker(fabric, workerName, queryString, bindVars string) (res *qw_req.ResponseForQueryWorker, err error) {
	if benchmark {
		defer utils.TimeTrack(time.Now(), "UpdateQueryWorker")
	}

	req := qw_req.NewRequestForUpdateQueryWorker(fabric, workerName, queryString, bindVars)
	res = qw_req.NewResponseForQueryWorker()
	err = c.client.Request(req, res)

	return res, CheckReturnError(err)
}

// DeleteQueryWorker
// Delete a query under the given fabric.
// https://macrometa.com/docs/api#/operations/DeleteRestqlByName
func (c QueryWorkerManager) DeleteQueryWorker(fabric, workerName string) (err error) {
	if benchmark {
		defer utils.TimeTrack(time.Now(), "DeleteQueryWorker")
	}

	req := qw_req.NewRequestForDeleteQueryWorker(fabric, workerName)
	res := qw_req.NewResponseForDeleteQueryWorker()
	err = c.client.Request(req, res)
	return CheckReturnError(err)
}
