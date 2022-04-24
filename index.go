package goC8

import (
	r "github.com/marvin-hansen/goC8/requests/index_req"
	"time"
)

const benchmark = true

// GetIndexes
// Fetches the list of all indexes of a collection.
// https://macrometa.com/docs/api#/operations/getIndexes
func (c Client) GetIndexes(fabric, collectionName string) (response *r.ResponseForGetAllIndices, err error) {
	if benchmark {
		defer TimeTrack(time.Now(), "GetIndexes")
	}

	req := r.NewRequestForGetAllIndices(fabric, collectionName)
	response = r.NewResponseForGetAllIndices()
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetIndex
// Fetches the information about index.
// https://macrometa.com/docs/api#/operations/getIndexes:handle
func (c Client) GetIndex(fabric, collectionName, indexName string) (response *r.ResponseForGetIndex, err error) {
	if benchmark {
		defer TimeTrack(time.Now(), "GetIndex")
	}

	req := r.NewRequestForGetIndex(fabric, collectionName, indexName)
	response = r.NewResponseForGetIndex()
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

// CreateFulltextIndex
// Creates fulltext index, if it does not already exist.
// field: Attribute to index. Must be of text format.
// minLength: Minimum character length of words to index.
func (c Client) CreateFulltextIndex(fabric, collectionName, field string, minLength int) (response *r.ResponseForCreateFulltextIndex, err error) {
	if benchmark {
		defer TimeTrack(time.Now(), "CreateFulltextIndex")
	}

	req := r.NewRequestForCreateFulltextIndex(fabric, collectionName, field, minLength)
	response = r.NewResponseForCreateFulltextIndex()
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

// CreateGeoIndex
// Creates a geo index.
// Note: Geo indexes are always sparse, meaning that documents that do not contain the index attributes or have non-numeric values in the index attributes will not be indexed.
func (c Client) CreateGeoIndex(fabric, collectionName, field string, geoJson bool) (response *r.ResponseForCreateGeoIndex, err error) {
	if benchmark {
		defer TimeTrack(time.Now(), "CreateGeoIndex")
	}

	req := r.NewRequestForCreateGeoIndex(fabric, collectionName, field, geoJson)
	response = r.NewResponseForCreateGeoIndex()
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c Client) CreateHashIndex(fabric, collectionName, field string, deduplicate, sparse, unique bool) (response *r.ResponseForCreateHashIndex, err error) {
	if benchmark {
		defer TimeTrack(time.Now(), "CreateHashIndex")
	}

	req := r.NewRequestForCreateHashIndex(fabric, collectionName, field, deduplicate, sparse, unique)
	response = r.NewResponseForCreateHashIndex()
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

// DeleteIndex
// Remove an index.
// https://macrometa.com/docs/api#/operations/dropIndex
func (c Client) DeleteIndex(fabric, collectionName, indexName string) (response *r.ResponseForDeleteIndex, err error) {
	if benchmark {
		defer TimeTrack(time.Now(), "DeleteIndex")
	}

	req := r.NewRequestForDeleteIndex(fabric, collectionName, indexName)
	response = r.NewResponseForDeleteIndex()
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}
