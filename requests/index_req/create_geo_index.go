package index_req

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForCreateGeoIndex(fabric, collectionName, field string, geoJson bool) *RequestForCreateGeoIndex {
	indexType := "geo"
	return &RequestForCreateGeoIndex{
		path:       fmt.Sprintf("_fabric/%v/_api/index/%v", fabric, indexType),
		parameters: fmt.Sprintf("?collection=%v", collectionName),
		payload:    getGeoIndexPayload(field, geoJson),
	}
}

// getGeoIndexPayload
// Constructs a JSON object with these properties:
// * fields (string): An array with one or two attribute paths. If it is an array with one attribute path location, then a geo-spatial index on all documents is created using location as path to the coordinates. The value of the attribute must be an array with at least two double values. The array must contain the latitude (first value) and the longitude (second value).
//Note: All documents, which do not have the attribute path or with value that are not suitable, are ignored.If it is an array with two attribute paths latitude and longitude, then a geo-spatial index on all documents is created using latitude and longitude as paths the latitude and the longitude. The value of the attribute latitude and of the attribute longitude must be a double. All documents, which do not have the attribute paths or which values are not suitable, are ignored.
// * type: Must be equal to "geo".
// * geoJson: If a geo-spatial index on a location is constructed and geoJson is true, then the order within the array is longitude followed by latitude. This corresponds to the format described in http://geojson.org/geojson-spec.html#positions
// Note: Geo indexes are always sparse, meaning that documents that do not contain the index attributes or have non-numeric values in the index attributes will not be indexed.
func getGeoIndexPayload(field string, geoJson bool) []byte {
	str := fmt.Sprintf(`
{
  "fields": [
    "%v"
  ],
  "geoJson": %v,
  "type": "geo"
}
`, field, geoJson)
	return []byte(str)
}

type RequestForCreateGeoIndex struct {
	path       string
	parameters string
	payload    []byte
}

func (req *RequestForCreateGeoIndex) Path() string {
	return req.path
}

func (req *RequestForCreateGeoIndex) Method() string {
	return http.MethodPost
}

func (req *RequestForCreateGeoIndex) Query() string {
	return ""
}

func (req *RequestForCreateGeoIndex) HasQueryParameter() bool {
	return true
}

func (req *RequestForCreateGeoIndex) GetQueryParameter() string {
	return req.parameters
}

func (req *RequestForCreateGeoIndex) Payload() []byte {
	return req.payload
}

func (req *RequestForCreateGeoIndex) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForCreateGeoIndex() *IndexEntry {
	return new(IndexEntry)
}
