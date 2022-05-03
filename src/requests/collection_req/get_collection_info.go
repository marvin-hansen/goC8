package collection_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForGetCollectionInfo(fabric, collectionName string) *RequestForGetCollectionInfo {
	return &RequestForGetCollectionInfo{
		path: fmt.Sprintf("_fabric/%v/_api/collection/%v", fabric, collectionName),
	}
}

type RequestForGetCollectionInfo struct {
	path string
}

func (req *RequestForGetCollectionInfo) Path() string {
	return req.path
}

func (req *RequestForGetCollectionInfo) Method() string {
	return http.MethodGet
}

func (req *RequestForGetCollectionInfo) Query() string {
	return ""
}

func (req *RequestForGetCollectionInfo) HasQueryParameter() bool {
	return false
}

func (req *RequestForGetCollectionInfo) GetQueryParameter() string {
	return ""
}

func (req *RequestForGetCollectionInfo) Payload() []byte {
	return nil
}

func (req *RequestForGetCollectionInfo) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForGetCollectionInfo() *ResponseForGetCollectionInfo {
	return new(ResponseForGetCollectionInfo)
}

type ResponseForGetCollectionInfo struct {
	Code             int    `json:"code"`
	Error            bool   `json:"error"`
	Id               string `json:"id"`
	Name             string `json:"name"`
	Status           int    `json:"status"`
	Type             int    `json:"type"`
	CollectionModel  string `json:"collectionModel"`
	IsSpot           bool   `json:"isSpot"`
	IsLocal          bool   `json:"isLocal"`
	HasStream        bool   `json:"hasStream"`
	WaitForSync      bool   `json:"waitForSync"`
	IsSystem         bool   `json:"isSystem"`
	GloballyUniqueId string `json:"globallyUniqueId"`
	SearchEnabled    bool   `json:"searchEnabled"`
}

func (r *ResponseForGetCollectionInfo) IsResponse() {}

func (r ResponseForGetCollectionInfo) String() string {
	return fmt.Sprintf("ID: %v, Name: %v,  Status: %v, Type: %v, CollectionModel: %v, IsSpot: %v, IsLocal: %v, HasStream: %v, WaitForSync: %v, IsSystem: %v, GloballyUniqueId: %v, SearchEnabled: %v",
		//Code: %v, Error: %v,
		//r.Code,
		//r.Error,
		r.Id,
		r.Name,
		r.Status,
		r.Type,
		r.CollectionModel,
		r.IsSpot,
		r.IsLocal,
		r.HasStream,
		r.WaitForSync,
		r.IsSystem,
		r.GloballyUniqueId,
		r.SearchEnabled,
	)
}
