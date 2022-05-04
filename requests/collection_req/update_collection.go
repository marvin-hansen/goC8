package collection_req

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//**// Request //**//

func NewRequestForUpdateCollection(fabric, collectionName string, properties *UpdateOptions) *RequestForUpdateCollection {
	return &RequestForUpdateCollection{
		path:    fmt.Sprintf("_fabric/%v/_api/collection/%v/properties", fabric, collectionName),
		payload: getUpdatePayload(properties),
	}
}

type RequestForUpdateCollection struct {
	path    string
	payload []byte
}

func getUpdatePayload(opts *UpdateOptions) []byte {
	data, err := json.MarshalIndent(opts, "", "")
	if err != nil {
		log.Println(err.Error())
	}
	return data
}

func (req *RequestForUpdateCollection) Path() string {
	return req.path
}

func (req *RequestForUpdateCollection) Method() string {
	return http.MethodPut
}

func (req *RequestForUpdateCollection) Query() string {
	return ""
}

func (req *RequestForUpdateCollection) HasQueryParameter() bool {
	return false
}

func (req *RequestForUpdateCollection) GetQueryParameter() string {
	return ""
}

func (req *RequestForUpdateCollection) Payload() []byte {
	return req.payload
}

func (req *RequestForUpdateCollection) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForUpdateCollection() *ResponseForUpdateCollection {
	return new(ResponseForUpdateCollection)
}

type ResponseForUpdateCollection struct {
	Code             int        `json:"code"`
	Error            bool       `json:"error"`
	Id               string     `json:"id"`
	Name             string     `json:"name"`
	Status           int        `json:"status"`
	Type             int        `json:"type"`
	CollectionModel  string     `json:"collectionModel"`
	IsSpot           bool       `json:"isSpot"`
	IsLocal          bool       `json:"isLocal"`
	HasStream        bool       `json:"hasStream"`
	WaitForSync      bool       `json:"waitForSync"`
	KeyOptions       KeyOptions `json:"keyOptions"`
	IsSystem         bool       `json:"isSystem"`
	GloballyUniqueId string     `json:"globallyUniqueId"`
	SearchEnabled    bool       `json:"searchEnabled"`
}

func (r *ResponseForUpdateCollection) IsResponse() {}

func (r ResponseForUpdateCollection) String() string {
	return fmt.Sprintf("ID: %v, Name: %v,  Status: %v, Type: %v, CollectionModel: %v, IsSpot: %v, IsLocal: %v, HasStream: %v, WaitForSync: %v, KeyOptions: %v, IsSystem: %v, GloballyUniqueId: %v, SearchEnabled: %v",
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
		r.KeyOptions,
		r.IsSystem,
		r.GloballyUniqueId,
		r.SearchEnabled,
	)
}
