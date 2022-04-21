package goC8

import r "github.com/marvin-hansen/goC8/requests/index_req"

// GetIndexes
// Fetches the list of all indexes of a collection.
func (c Client) GetIndexes(fabric, collectionName string) (response *r.ResponseForGetAllIndices, err error) {
	req := r.NewRequestForGetAllIndices(fabric, collectionName)
	response = r.NewResponseForGetAllIndices()
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}
