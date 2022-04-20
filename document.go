package goC8

import (
	r "github.com/marvin-hansen/goC8/requests/document_req"
)

// CreateNewDocument
// silent - If set to false, the primary key of the new doc is returned. If set to true, an empty object is returned as response. No meta-data is returned for the created document. This option can be used to save some network traffic. True by default
// parameters - additional query parameters for non-standard cases.
// jsonDocument the document to store in the collection
func (c Client) CreateNewDocument(
	fabric string, collectionName string, silent bool, jsonDocument []byte,
	parameters *r.CreateDocumentParameters) (response *r.ResponseForCreateDocument, err error) {

	if parameters == nil {
		parameters = r.GetDefaultCreateDocumentParameters()
	}

	req := r.NewRequestForCreateDocument(fabric, collectionName, silent, jsonDocument, parameters)
	response = r.NewResponseForCreateDocument()
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c Client) GetDocument(
	fabric string, collectionName string, key string) (response *r.ResponseForGetJsonDocument, err error) {

	req := r.NewRequestForGetDocument(fabric, collectionName, key)
	response = r.NewResponseForGetJsonDocument()
	if err = c.requestJsonResponse(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c Client) UpdateDocument(
	fabric string, collectionName string, key string,
	jsonDocument []byte,
	silent bool,
	parameters *r.UpdateDocumentParameters,
) (response *r.ResponseForUpdateDocument, err error) {

	if parameters == nil {
		parameters = r.GetDefaultUpdateDocumentParameters()
	}

	req := r.NewRequestForUpdateDocument(fabric, collectionName, key, jsonDocument, silent, parameters)
	response = r.NewResponseForUpdateDocument()
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c Client) ReplaceDocument(
	fabric string, collectionName, key string, jsonDocuments []byte,
	parameters *r.ReplaceDocumentParameters) (response *r.ResponseForReplaceDocument, err error) {

	if parameters == nil {
		parameters = r.GetDefaultReplaceDocumentParameters()
	}

	req := r.NewRequestForReplaceDocument(fabric, collectionName, key, jsonDocuments, parameters)
	response = r.NewResponseForReplaceDocument()
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c Client) DeleteDocument(
	fabric string, collectionName string, key string,
	parameters *r.DeleteDocumentParameters) (response *r.ResponseForDeleteDocument, err error) {

	if parameters == nil {
		parameters = r.GetDefaultDeleteDocumentParameters()
	}

	req := r.NewRequestForDeleteDocument(fabric, collectionName, key, parameters)
	response = r.NewResponseForDeleteDocument()
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}
