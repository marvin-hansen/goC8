package document_req

type CreateDocumentParameters struct {
	returnNew   bool // If set to true, adds the new documents to the new attribute. False by default.
	returnOld   bool // If set to true, adds the old attribute which displays the previous version of the document. Only available if the overwrite option is set to true. False by default
	overwrite   bool // If set to true, the insert becomes a replace-insert. If a document with the same _key already exists the new document is not rejected with unique constraint violated but replaces the old document. False by default
	waitForSync bool // If set to true, returns only after data has been synced to disk. // False by default
}

// GetDefaultCreateDocumentParameters default values for createDocument paramters
//	returnNew   If set to true, adds the new documents to the new attribute. False by default.
//	returnOld   If set to true, adds the old attribute which displays the previous version of the document. Only available if the overwrite option is set to true. False by default
//	overwrite   If set to true, the insert becomes a replace-insert. If a document with the same _key already exists the new document is not rejected with unique constraint violated but replaces the old document. False by default
//	waitForSync If set to true, returns only after data has been synced to disk. // False by default
func GetDefaultCreateDocumentParameters() *CreateDocumentParameters {
	return &CreateDocumentParameters{
		returnNew:   false,
		returnOld:   false,
		overwrite:   false,
		waitForSync: false,
	}
}

// GetCustomCreateDocumentParameters default values for createDocument paramters
//	returnNew   If set to true, adds the new documents to the new attribute. False by default.
//	returnOld   If set to true, adds the old attribute which displays the previous version of the document. Only available if the overwrite option is set to true. False by default
//	overwrite   If set to true, the insert becomes a replace-insert. If a document with the same _key already exists the new document is not rejected with unique constraint violated but replaces the old document. False by default
//	waitForSync If set to true, returns only after data has been synced to disk. // False by default
func GetCustomCreateDocumentParameters(returnNew, returnOld, overwrite, waitForSync bool) *CreateDocumentParameters {
	return &CreateDocumentParameters{
		returnNew:   returnNew,
		returnOld:   returnOld,
		overwrite:   overwrite,
		waitForSync: waitForSync,
	}
}
