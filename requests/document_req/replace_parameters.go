package document_req

type ReplaceDocumentParameters struct {
	ignoreRevs  bool // If set to true, adds the old attribute which displays original document. Default value : true
	returnOld   bool // If set to true, adds the old attribute which displays the previous version of the document. Only available if the overwrite option is set to true. False by default
	returnNew   bool // If set to true, adds the new documents to the new attribute. False by default.
	waitForSync bool // If set to true, returns only after data has been synced to disk. False by default.
}

//GetDefaultReplaceDocumentParameters
//	ignoreRevs  If set to true, adds the old attribute which displays original document. Default value : true
//	returnOld   If set to true, adds the old attribute which displays the previous version of the document. Only available if the overwrite option is set to true. False by default
//	returnNew   If set to true, adds the new documents to the new attribute. False by default.
//	waitForSync If set to true, returns only after data has been synced to disk. False by default.
func GetDefaultReplaceDocumentParameters() *ReplaceDocumentParameters {
	return &ReplaceDocumentParameters{
		ignoreRevs:  true,
		returnOld:   false,
		returnNew:   false,
		waitForSync: false,
	}
}

//GetCustomReplaceDocumentParameters
//	ignoreRevs  If set to true, adds the old attribute which displays original document. Default value : true
//	returnOld   If set to true, adds the old attribute which displays the previous version of the document. Only available if the overwrite option is set to true. False by default
//	returnNew   If set to true, adds the new documents to the new attribute. False by default.
//	waitForSync If set to true, returns only after data has been synced to disk. False by default.
func GetCustomReplaceDocumentParameters(ignoreRevs, returnOld, returnNew, waitForSync bool) *ReplaceDocumentParameters {
	return &ReplaceDocumentParameters{
		ignoreRevs:  ignoreRevs,
		returnOld:   returnOld,
		returnNew:   returnNew,
		waitForSync: waitForSync,
	}
}
