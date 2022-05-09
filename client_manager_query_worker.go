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

func (c QueryWorkerManager) CreateQueryWorker(fabric, workerName, queryString, parameterString string) (res *qw_req.ResponseForCreateQueryWorker, err error) {
	if benchmark {
		defer utils.TimeTrack(time.Now(), "CreateQueryWorker")
	}

	req := qw_req.NewRequestForCreateQueryWorker(fabric, workerName, queryString, parameterString)
	res = qw_req.NewResponseForCreateQueryWorker()
	err = c.client.Request(req, res)
	return res, CheckReturnError(err)
}

func (c QueryWorkerManager) RunQueryWorker(fabric, workerName string) (res *query_req.ResponseForExplainQuery, err error) {
	if benchmark {
		defer utils.TimeTrack(time.Now(), "RunQueryWorker")
	}

	return res, CheckReturnError(err)
}

func (c QueryWorkerManager) ReadAllQueryWorkers(fabric string) (res *query_req.ResponseForExplainQuery, err error) {
	if benchmark {
		defer utils.TimeTrack(time.Now(), "ReadAllQueryWorkers")
	}

	return res, CheckReturnError(err)
}

func (c QueryWorkerManager) UpdateQueryWorker(fabric, workerName string) (res *query_req.ResponseForExplainQuery, err error) {
	if benchmark {
		defer utils.TimeTrack(time.Now(), "UpdateQueryWorker")
	}

	return res, CheckReturnError(err)
}

func (c QueryWorkerManager) DeleteQueryWorker(fabric, workerName string) (res *query_req.ResponseForExplainQuery, err error) {
	if benchmark {
		defer utils.TimeTrack(time.Now(), "DeleteQueryWorker")
	}

	return res, CheckReturnError(err)
}
