package goC8

import (
	"github.com/marvin-hansen/goC8/requests/document_req"
	"strings"
)

type DocumentManager struct {
	client *Client
}

func NewDocumentManager(client *Client) *DocumentManager {
	return &DocumentManager{client: client}
}

// CreateNewDocument
// silent - If set to false, the primary key of the new doc is returned. If set to true, an empty object is returned as response. No meta-data is returned for the created document. This option can be used to save some network traffic. True by default
// parameters - additional query parameters for non-standard cases.
// jsonDocument the document to store in the collection
func (c DocumentManager) CreateNewDocument(
	fabric string, collectionName string, silent bool, jsonDocument []byte,
	parameters *document_req.CreateDocumentParameters) (response *document_req.ResponseForCreateDocument, err error) {

	if parameters == nil {
		parameters = document_req.GetDefaultCreateDocumentParameters()
	}

	req := document_req.NewRequestForCreateDocument(fabric, collectionName, silent, jsonDocument, parameters)
	response = document_req.NewResponseForCreateDocument()
	if err = c.client.Request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c DocumentManager) GetDocument(
	fabric string, collectionName string, key string) (response *document_req.ResponseForGetJsonDocument, err error) {

	req := document_req.NewRequestForGetDocument(fabric, collectionName, key)
	response = document_req.NewResponseForGetJsonDocument()
	if err = c.client.requestJsonResponse(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c DocumentManager) CheckDocumentExists(
	fabric string, collectionName string, key string) (exists bool, err error) {

	req := document_req.NewRequestForGetDocument(fabric, collectionName, key)
	response := document_req.NewResponseForGetJsonDocument()
	if err = c.client.requestJsonResponse(req, response); err != nil {

		if strings.Contains(err.Error(), "1202") { // Number=1202,  Error Message=document not found
			return false, nil
		} else {
			return false, err // Any other error
		}
	}

	return true, nil
}

func (c DocumentManager) UpdateDocument(
	fabric string, collectionName string, key string,
	jsonDocument []byte,
	silent bool,
	parameters *document_req.UpdateDocumentParameters,
) (response *document_req.ResponseForUpdateDocument, err error) {

	if parameters == nil {
		parameters = document_req.GetDefaultUpdateDocumentParameters()
	}

	req := document_req.NewRequestForUpdateDocument(fabric, collectionName, key, jsonDocument, silent, parameters)
	response = document_req.NewResponseForUpdateDocument()
	if err = c.client.Request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c DocumentManager) ReplaceDocument(
	fabric string, collectionName, key string, jsonDocuments []byte,
	parameters *document_req.ReplaceDocumentParameters) (response *document_req.ResponseForReplaceDocument, err error) {

	if parameters == nil {
		parameters = document_req.GetDefaultReplaceDocumentParameters()
	}

	req := document_req.NewRequestForReplaceDocument(fabric, collectionName, key, jsonDocuments, parameters)
	response = document_req.NewResponseForReplaceDocument()
	if err = c.client.Request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c DocumentManager) DeleteDocument(
	fabric string, collectionName string, key string,
	parameters *document_req.DeleteDocumentParameters) (response *document_req.ResponseForDeleteDocument, err error) {

	if parameters == nil {
		parameters = document_req.GetDefaultDeleteDocumentParameters()
	}

	req := document_req.NewRequestForDeleteDocument(fabric, collectionName, key, parameters)
	response = document_req.NewResponseForDeleteDocument()
	if err = c.client.Request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

/////////// API to handle many documents ///////////

func (c DocumentManager) UpdateManyDocuments(
	fabric string, collectionName string, jsonDocument []byte,
	parameters *document_req.UpdateDocumentParameters) (response *document_req.ResponseForUpdateManyDocument, err error) {

	if parameters == nil {
		parameters = document_req.GetDefaultUpdateDocumentParameters()
	}

	req := document_req.NewRequestForUpdateManyDocuments(fabric, collectionName, jsonDocument, parameters)
	response = document_req.NewResponseForUpdateManyDocuments()
	if err = c.client.Request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c DocumentManager) ReplaceManyDocuments(
	fabric string, collectionName string, jsonDocuments []byte,
	parameters *document_req.ReplaceDocumentParameters) (response *document_req.ResponseForReplaceManyDocument, err error) {

	if parameters == nil {
		parameters = document_req.GetDefaultReplaceDocumentParameters()
	}

	req := document_req.NewRequestForReplaceManyDocument(fabric, collectionName, jsonDocuments, parameters)
	response = document_req.NewResponseForReplaceManyDocument()
	if err = c.client.Request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c DocumentManager) DeleteManyDocuments(
	fabric string, collectionName string, keysToDelete []byte,
	parameters *document_req.DeleteDocumentParameters) (response *document_req.ResponseForDeleteManyDocuments, err error) {

	if parameters == nil {
		parameters = document_req.GetDefaultDeleteDocumentParameters()
	}

	req := document_req.NewRequestForDeleteManyDocuments(fabric, collectionName, keysToDelete, parameters)
	response = document_req.NewResponseForDeleteManyDocuments()
	if err = c.client.Request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}
