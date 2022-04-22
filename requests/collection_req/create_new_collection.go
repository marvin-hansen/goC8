package collection_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForCreateNewCollection(fabric, collectionName string, allowUserKeys bool, collectionType CollectionType) *RequestForCreateNewCollection {
	return &RequestForCreateNewCollection{
		path:    fmt.Sprintf("_fabric/%v/_api/collection", fabric),
		payload: getCreatePayload(collectionName, allowUserKeys, collectionType),
	}
}

type RequestForCreateNewCollection struct {
	path    string
	payload []byte
}

func getCreatePayload(collectionName string, allowUserKeys bool, collectionType CollectionType) []byte {
	// https://macrometa.com/docs/collections/documents/tutorials/using_rest_api/
	s := fmt.Sprintf(`{
      "name": "%v",
      "type": %v
	}`,
		collectionName, collectionType.ToInt(),
	)
	return []byte(s)
}

func (req *RequestForCreateNewCollection) Path() string {
	return req.path
}

func (req *RequestForCreateNewCollection) Method() string {
	return http.MethodPost
}

func (req *RequestForCreateNewCollection) Query() string {
	return ""
}

func (req *RequestForCreateNewCollection) HasQueryParameter() bool {
	return false
}

func (req *RequestForCreateNewCollection) GetQueryParameter() string {
	return ""
}

func (req *RequestForCreateNewCollection) Payload() []byte {
	return req.payload
}

func (req *RequestForCreateNewCollection) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForCreateNewCollection() *ResponseForCreateNewCollection {
	return new(ResponseForCreateNewCollection)
}

func (r ResponseForCreateNewCollection) String() string {
	return fmt.Sprintf("Result: %v", r.Result)
}

type ResponseForCreateNewCollection struct {
	Result []ResulFromCollections `json:"result"`
}

func (r *ResponseForCreateNewCollection) IsResponse() {}
