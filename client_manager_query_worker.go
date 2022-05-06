package goC8

import "github.com/marvin-hansen/goC8/requests/query_req"

type QueryWorkerManager struct {
	client *Client
}

func NewQueryWorkerManager(client *Client) *QueryWorkerManager {
	return &QueryWorkerManager{client: client}
}

func (c QueryWorkerManager) CreateQueryWorker(fabric string) (res *query_req.ResponseForExplainQuery, err error) {

	return res, CheckReturnError(err)
}

func (c QueryWorkerManager) RunQueryWorker(fabric, workerName string) (res *query_req.ResponseForExplainQuery, err error) {

	return res, CheckReturnError(err)
}

func (c QueryWorkerManager) ReadAllQueryWorkers(fabric string) (res *query_req.ResponseForExplainQuery, err error) {

	return res, CheckReturnError(err)
}

func (c QueryWorkerManager) UpdateQueryWorker(fabric, workerName string) (res *query_req.ResponseForExplainQuery, err error) {

	return res, CheckReturnError(err)
}

func (c QueryWorkerManager) DeleteQueryWorker(fabric, workerName string) (res *query_req.ResponseForExplainQuery, err error) {

	return res, CheckReturnError(err)
}
