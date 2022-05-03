package goC8

import (
	document_req2 "github.com/marvin-hansen/goC8/src/requests/document_req"
	"strings"
)

// CreateNewDocument
// silent - If set to false, the primary key of the new doc is returned. If set to true, an empty object is returned as response. No meta-data is returned for the created document. This option can be used to save some network traffic. True by default
// parameters - additional query parameters for non-standard cases.
// jsonDocument the document to store in the collection
func (c Client) CreateNewDocument(
	fabric string, collectionName string, silent bool, jsonDocument []byte,
	parameters *document_req2.CreateDocumentParameters) (response *document_req2.ResponseForCreateDocument, err error) {

	if parameters == nil {
		parameters = document_req2.GetDefaultCreateDocumentParameters()
	}

	req := document_req2.NewRequestForCreateDocument(fabric, collectionName, silent, jsonDocument, parameters)
	response = document_req2.NewResponseForCreateDocument()
	if err = c.Request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c Client) GetDocument(
	fabric string, collectionName string, key string) (response *document_req2.ResponseForGetJsonDocument, err error) {

	req := document_req2.NewRequestForGetDocument(fabric, collectionName, key)
	response = document_req2.NewResponseForGetJsonDocument()
	if err = c.requestJsonResponse(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c Client) CheckDocumentExists(
	fabric string, collectionName string, key string) (exists bool, err error) {

	req := document_req2.NewRequestForGetDocument(fabric, collectionName, key)
	response := document_req2.NewResponseForGetJsonDocument()
	if err = c.requestJsonResponse(req, response); err != nil {

		if strings.Contains(err.Error(), "1202") { // Number=1202,  Error Message=document not found
			return false, nil
		} else {
			return false, err // Any other error
		}
	}

	return true, nil
}

func (c Client) UpdateDocument(
	fabric string, collectionName string, key string,
	jsonDocument []byte,
	silent bool,
	parameters *document_req2.UpdateDocumentParameters,
) (response *document_req2.ResponseForUpdateDocument, err error) {

	if parameters == nil {
		parameters = document_req2.GetDefaultUpdateDocumentParameters()
	}

	req := document_req2.NewRequestForUpdateDocument(fabric, collectionName, key, jsonDocument, silent, parameters)
	response = document_req2.NewResponseForUpdateDocument()
	if err = c.Request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c Client) ReplaceDocument(
	fabric string, collectionName, key string, jsonDocuments []byte,
	parameters *document_req2.ReplaceDocumentParameters) (response *document_req2.ResponseForReplaceDocument, err error) {

	if parameters == nil {
		parameters = document_req2.GetDefaultReplaceDocumentParameters()
	}

	req := document_req2.NewRequestForReplaceDocument(fabric, collectionName, key, jsonDocuments, parameters)
	response = document_req2.NewResponseForReplaceDocument()
	if err = c.Request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c Client) DeleteDocument(
	fabric string, collectionName string, key string,
	parameters *document_req2.DeleteDocumentParameters) (response *document_req2.ResponseForDeleteDocument, err error) {

	if parameters == nil {
		parameters = document_req2.GetDefaultDeleteDocumentParameters()
	}

	req := document_req2.NewRequestForDeleteDocument(fabric, collectionName, key, parameters)
	response = document_req2.NewResponseForDeleteDocument()
	if err = c.Request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}
