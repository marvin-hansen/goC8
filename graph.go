package goC8

import r "github.com/marvin-hansen/goC8/requests/graph_req"

func (c Client) GetAllGraphs(fabric string) (response *r.ResponseForGetAllGraphs, err error) {

	req := r.NewRequestForGetAllGraphs(fabric)
	response = r.NewResponseForGetAllGraphs()
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}
