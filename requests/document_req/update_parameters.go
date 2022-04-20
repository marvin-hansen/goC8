package document_req

type UpdateDocumentParameters struct {
	keepNull     bool // If set to false, you can use the patch command to remove attributes from an existing document if the patch document contains the same attributes with a null value.
	mergeObjects bool // Controls whether objects (not arrays) are merged if present in both the existing and the patch document. If set to false, the value in the patch document overwrites the existing document's value. If set to true, objects are merged.
	ignoreRevs   bool // If this is set to true, the _rev attributes in the given documents are ignored. If this is set to false, then any _rev attribute given in a body document is taken as a precondition. The document is only updated if the current revision is the one specified.
	returnOld    bool // If set to true, adds the old attribute which displays the previous version of the document. Only available if the overwrite option is set to true.
	returnNew    bool // If set to true, adds the new documents to the new attribute.
	waitForSync  bool // If set to true, returns only after data has been synced to disk.
}

//GetDefaultUpdateDocumentParameters
//	keepNull     If set to false, you can use the patch command to remove attributes from an existing document if the patch document contains the same attributes with a null value. True by default
//	mergeObjects Controls whether objects (not arrays) are merged if present in both the existing and the patch document. If set to false, the value in the patch document overwrites the existing document's value. If set to true, objects are merged. True by default
//	ignoreRevs   If this is set to true, the _rev attributes in the given documents are ignored. If this is set to false, then any _rev attribute given in a body document is taken as a precondition. The document is only updated if the current revision is the one specified. True by default
//	returnOld    If set to true, adds the old attribute which displays the previous version of the document. Only available if the overwrite option is set to true. False by default
//	returnNew    If set to true, adds the new documents to the new attribute. False by default. False by default
//	waitForSync  If set to true, returns only after data has been synced to disk.  False by default
func GetDefaultUpdateDocumentParameters() *UpdateDocumentParameters {
	return &UpdateDocumentParameters{
		keepNull:     true,
		mergeObjects: true,
		ignoreRevs:   true,
		returnOld:    false,
		returnNew:    false,
		waitForSync:  false,
	}
}

// GetCustomUpdateDocumentParameters
//	keepNull     If set to false, you can use the patch command to remove attributes from an existing document if the patch document contains the same attributes with a null value. True by default
//	mergeObjects Controls whether objects (not arrays) are merged if present in both the existing and the patch document. If set to false, the value in the patch document overwrites the existing document's value. If set to true, objects are merged. True by default
//	ignoreRevs   If this is set to true, the _rev attributes in the given documents are ignored. If this is set to false, then any _rev attribute given in a body document is taken as a precondition. The document is only updated if the current revision is the one specified. True by default
//	returnOld    If set to true, adds the old attribute which displays the previous version of the document. Only available if the overwrite option is set to true. False by default
//	returnNew    If set to true, adds the new documents to the new attribute. False by default. False by default
//	waitForSync  If set to true, returns only after data has been synced to disk.  False by default
func GetCustomUpdateDocumentParameters(keepNull, mergeObjects, ignoreRevs, returnOld, returnNew, waitForSync bool) *UpdateDocumentParameters {
	return &UpdateDocumentParameters{
		keepNull:     keepNull,
		mergeObjects: mergeObjects,
		ignoreRevs:   ignoreRevs,
		returnOld:    returnOld,
		returnNew:    returnNew,
		waitForSync:  waitForSync,
	}
}
