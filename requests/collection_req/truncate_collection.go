package collection_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForTruncateCollection(fabric, collectionName string) *RequestForTruncateCollection {
	return &RequestForTruncateCollection{
		path: fmt.Sprintf("_fabric/%v/_api/collection/%v/truncate", fabric, collectionName),
	}
}

type RequestForTruncateCollection struct {
	path string
}

func (req *RequestForTruncateCollection) Path() string {
	return req.path
}

func (req *RequestForTruncateCollection) Method() string {
	return http.MethodPut
}

func (req *RequestForTruncateCollection) Query() string {
	return ""
}

func (req *RequestForTruncateCollection) HasQueryParameter() bool {
	return false
}

func (req *RequestForTruncateCollection) GetQueryParameter() string {
	return ""
}

func (req *RequestForTruncateCollection) Payload() []byte {
	return nil
}

func (req *RequestForTruncateCollection) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForTruncateCollection() *ResponseForTruncateCollection {
	return new(ResponseForTruncateCollection)
}

type ResponseForTruncateCollection struct {
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

func (r *ResponseForTruncateCollection) IsResponse() {}

func (r ResponseForTruncateCollection) String() string {
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
