package goC8

import (
	"github.com/marvin-hansen/goC8/src/requests/document_req"
)

func (c Client) UpdateManyDocuments(
	fabric string, collectionName string, jsonDocument []byte,
	parameters *document_req.UpdateDocumentParameters) (response *document_req.ResponseForUpdateManyDocument, err error) {

	if parameters == nil {
		parameters = document_req.GetDefaultUpdateDocumentParameters()
	}

	req := document_req.NewRequestForUpdateManyDocuments(fabric, collectionName, jsonDocument, parameters)
	response = document_req.NewResponseForUpdateManyDocuments()
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c Client) ReplaceManyDocuments(
	fabric string, collectionName string, jsonDocuments []byte,
	parameters *document_req.ReplaceDocumentParameters) (response *document_req.ResponseForReplaceManyDocument, err error) {

	if parameters == nil {
		parameters = document_req.GetDefaultReplaceDocumentParameters()
	}

	req := document_req.NewRequestForReplaceManyDocument(fabric, collectionName, jsonDocuments, parameters)
	response = document_req.NewResponseForReplaceManyDocument()
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c Client) DeleteManyDocuments(
	fabric string, collectionName string, keysToDelete []byte,
	parameters *document_req.DeleteDocumentParameters) (response *document_req.ResponseForDeleteManyDocuments, err error) {

	if parameters == nil {
		parameters = document_req.GetDefaultDeleteDocumentParameters()
	}

	req := document_req.NewRequestForDeleteManyDocuments(fabric, collectionName, keysToDelete, parameters)
	response = document_req.NewResponseForDeleteManyDocuments()
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}
