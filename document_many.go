package goC8

import (
	r "github.com/marvin-hansen/goC8/requests/document_req"
)

func (c Client) UpdateManyDocuments(
	fabric string, collectionName string, jsonDocument []byte,
	parameters *r.UpdateDocumentParameters) (response *r.ResponseForUpdateManyDocument, err error) {

	if parameters == nil {
		parameters = r.GetDefaultUpdateDocumentParameters()
	}

	req := r.NewRequestForUpdateManyDocuments(fabric, collectionName, jsonDocument, parameters)
	response = r.NewResponseForUpdateManyDocuments()
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c Client) ReplaceManyDocuments(
	fabric string, collectionName string, jsonDocuments []byte,
	parameters *r.ReplaceDocumentParameters) (response *r.ResponseForReplaceManyDocument, err error) {

	if parameters == nil {
		parameters = r.GetDefaultReplaceDocumentParameters()
	}

	req := r.NewRequestForReplaceManyDocument(fabric, collectionName, jsonDocuments, parameters)
	response = r.NewResponseForReplaceManyDocument()
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c Client) DeleteManyDocuments(
	fabric string, collectionName string, keysToDelete []byte,
	parameters *r.DeleteDocumentParameters) (response *r.ResponseForDeleteManyDocuments, err error) {

	if parameters == nil {
		parameters = r.GetDefaultDeleteDocumentParameters()
	}

	req := r.NewRequestForDeleteManyDocuments(fabric, collectionName, keysToDelete, parameters)
	response = r.NewResponseForDeleteManyDocuments()
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}
