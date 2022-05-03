package goC8

import (
	collection_req2 "github.com/marvin-hansen/goC8/src/requests/collection_req"
	"strings"
)

func (c Client) GetAllCollections(fabric string) (response *collection_req2.ResponseForGetAllCollections, err error) {
	req := collection_req2.NewRequestForGetAllCollections(fabric)
	response = collection_req2.NewResponseForGetAllCollections()
	if err = c.Request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c Client) CreateNewCollection(fabric, collectionName string, allowUserKeys bool, collectionType collection_req2.CollectionType) (err error) {
	req := collection_req2.NewRequestForCreateNewCollection(fabric, collectionName, allowUserKeys, collectionType)
	response := collection_req2.NewResponseForCreateNewCollection()
	if err = c.Request(req, response); err != nil {
		return err
	}
	return nil
}

func (c Client) GetCollectionInfo(fabric, collectionName string) (response *collection_req2.ResponseForGetCollectionInfo, err error) {
	req := collection_req2.NewRequestForGetCollectionInfo(fabric, collectionName)
	response = collection_req2.NewResponseForGetCollectionInfo()
	if err = c.Request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c Client) CheckCollectionExists(fabric, collectionName string) (exists bool, err error) {
	req := collection_req2.NewRequestForGetCollectionInfo(fabric, collectionName)
	response := collection_req2.NewResponseForGetCollectionInfo()
	if err = c.Request(req, response); err != nil {
		if strings.Contains(err.Error(), "1203") { // Number=1203,  Error Message=collection or view not found
			return false, nil
		} else {
			return false, err // Any other error
		}
	}
	return true, nil
}

// UpdateCollectionProperties updates collection properties.
// Note: except for waitForSync and hasStream, collection properties cannot be changed once a collection is created.
func (c Client) UpdateCollectionProperties(fabric, collectionName string, properties *collection_req2.UpdateOptions) (response *collection_req2.ResponseForUpdateCollection, err error) {
	req := collection_req2.NewRequestForUpdateCollection(fabric, collectionName, properties)
	response = collection_req2.NewResponseForUpdateCollection()
	if err = c.Request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c Client) TruncateCollection(fabric, collectionName string) (response *collection_req2.ResponseForTruncateCollection, err error) {
	req := collection_req2.NewRequestForTruncateCollection(fabric, collectionName)
	response = collection_req2.NewResponseForTruncateCollection()
	if err = c.Request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c Client) DeleteCollection(fabric, collectionName string, isSystem bool) (err error) {
	req := collection_req2.NewRequestForDeleteCollection(fabric, collectionName, isSystem)
	response := collection_req2.NewResponseForDeleteCollection()
	if err = c.Request(req, response); err != nil {
		return err
	}
	// valid response: Code: 200, Error: false, ID: 159XXXXX
	return nil
}
