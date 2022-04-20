package document_req

type DeleteDocumentParameters struct {
	ignoreRevs  bool // If set to true, adds the old attribute which displays original document. Default value : false
	returnOld   bool // If set to true, adds the old attribute which displays the previous version of the document. Only available if the overwrite option is set to true. False by default
	silent      bool // If set to true, an empty object will be returned as response. No meta-data will be returned for the removed document. This option can be used to save some network traffic.
	waitForSync bool // If set to true, returns only after data has been synced to disk. // False by default
}

// GetDefaultDeleteDocumentParameters
//  ignoreRevs  If set to true, adds the old attribute which displays original document. Default value : false
//	returnOld   If set to true, adds the old attribute which displays the previous version of the document. Only available if the overwrite option is set to true. False by default
//  silent If set to true, an empty object will be returned as response. No meta-data will be returned for the removed document. This option can be used to save some network traffic. True by default
//	waitForSync If set to true, returns only after data has been synced to disk. // False by default
func GetDefaultDeleteDocumentParameters() *DeleteDocumentParameters {
	return &DeleteDocumentParameters{
		ignoreRevs:  false,
		returnOld:   false,
		silent:      true,
		waitForSync: false,
	}
}

//GetCustomDeleteDocumentParameters
//  ignoreRevs  If set to true, adds the old attribute which displays original document. Default value : false
//	returnOld   If set to true, adds the old attribute which displays the previous version of the document. Only available if the overwrite option is set to true. False by default
//  silent If set to true, an empty object will be returned as response. No meta-data will be returned for the removed document. This option can be used to save some network traffic.
//	waitForSync If set to true, returns only after data has been synced to disk. // False by default
func GetCustomDeleteDocumentParameters(ignoreRevs, returnOld, silent, waitForSync bool) *DeleteDocumentParameters {
	return &DeleteDocumentParameters{
		ignoreRevs:  ignoreRevs,
		returnOld:   returnOld,
		silent:      silent,
		waitForSync: waitForSync,
	}
}
