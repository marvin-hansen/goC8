package goC8

import (
	"github.com/marvin-hansen/goC8/src/requests/collection_req"
	"strings"
)

type CollectionManager struct {
	client *Client
}

func NewCollectionManager(client *Client) *CollectionManager {
	return &CollectionManager{client: client}
}

func (c CollectionManager) GetAllCollections(fabric string) (response *collection_req.ResponseForGetAllCollections, err error) {
	req := collection_req.NewRequestForGetAllCollections(fabric)
	response = collection_req.NewResponseForGetAllCollections()
	if err = c.client.Request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c CollectionManager) CreateNewCollection(fabric, collectionName string, allowUserKeys bool, collectionType collection_req.CollectionType) (err error) {
	req := collection_req.NewRequestForCreateNewCollection(fabric, collectionName, allowUserKeys, collectionType)
	response := collection_req.NewResponseForCreateNewCollection()
	if err = c.client.Request(req, response); err != nil {
		return err
	}
	return nil
}

func (c CollectionManager) GetCollectionInfo(fabric, collectionName string) (response *collection_req.ResponseForGetCollectionInfo, err error) {
	req := collection_req.NewRequestForGetCollectionInfo(fabric, collectionName)
	response = collection_req.NewResponseForGetCollectionInfo()
	if err = c.client.Request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c CollectionManager) CheckCollectionExists(fabric, collectionName string) (exists bool, err error) {
	req := collection_req.NewRequestForGetCollectionInfo(fabric, collectionName)
	response := collection_req.NewResponseForGetCollectionInfo()
	if err = c.client.Request(req, response); err != nil {
		if strings.Contains(err.Error(), "1203") { // Number=1203,  Error Message=collection or view not found
			return false, nil
		} else {
			return false, err // Any other error
		}
	}
	return true, nil
}

func (c CollectionManager) CountCollection(fabric, collectionName string) (response *collection_req.ResultFromCollection, err error) {
	req := collection_req.NewRequestForCountCollection(fabric, collectionName)
	response = collection_req.NewResultFromCollection()
	if err = c.client.Request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

// UpdateCollectionProperties updates collection properties.
// Note: except for waitForSync and hasStream, collection properties cannot be changed once a collection is created.
func (c CollectionManager) UpdateCollectionProperties(fabric, collectionName string, properties *collection_req.UpdateOptions) (response *collection_req.ResponseForUpdateCollection, err error) {
	req := collection_req.NewRequestForUpdateCollection(fabric, collectionName, properties)
	response = collection_req.NewResponseForUpdateCollection()
	if err = c.client.Request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c CollectionManager) TruncateCollection(fabric, collectionName string) (response *collection_req.ResponseForTruncateCollection, err error) {
	req := collection_req.NewRequestForTruncateCollection(fabric, collectionName)
	response = collection_req.NewResponseForTruncateCollection()
	if err = c.client.Request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c CollectionManager) DeleteCollection(fabric, collectionName string, isSystem bool) (err error) {
	req := collection_req.NewRequestForDeleteCollection(fabric, collectionName, isSystem)
	response := collection_req.NewResponseForDeleteCollection()
	if err = c.client.Request(req, response); err != nil {
		return err
	}
	// valid response: Code: 200, Error: false, ID: 159XXXXX
	return nil
}
