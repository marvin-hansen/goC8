package index_req

// NewCreateIndexParameters
// fields: An array of attribute names as array of strings.
//to exactly one attribute.
//Type: Type of index.
//Sparse: True if the index is sparse type.
//Deduplicate: True then it controls whether inserting duplicate index values from the same document into a unique array index will lead to a unique constraint error or not. (The default is true, so only a single instance of each non-unique index value will be inserted into the index per document.)
//Unique: True if the index is unique.
//Note: Unique indexes on non-shard keys are not supported in a cluster.
//The following index types do not support uniqueness,
// and using the unique attribute with these types may lead to an error:
// - geo indexes
// - fulltext indexes
//
//Note:
// Hash, skiplist and persistent indexes can optionally be created in a sparse variant.
// Sparse indexes do not index documents for which any of the index attributes is either not set or is null.
func NewCreateIndexParameters(indexFields []string, sparse, unique, deduplicate bool) *CreateIndexParameters {
	return &CreateIndexParameters{
		Fields:      indexFields,
		Type:        "hash",
		Sparse:      sparse,
		Unique:      unique,
		Deduplicate: deduplicate,
	}
}

// CreateIndexParameters
// fields: An array of attribute names as array of strings.
//to exactly one attribute.
//Type: Type of index.
//Sparse: True if the index is sparse type.
//Deduplicate: True then it controls whether inserting duplicate index values from the same document into a unique array index will lead to a unique constraint error or not. (The default is true, so only a single instance of each non-unique index value will be inserted into the index per document.)
//Unique: True if the index is unique.
//Note: Unique indexes on non-shard keys are not supported in a cluster.
//The following index types do not support uniqueness,
// and using the unique attribute with these types may lead to an error:
// - geo indexes
// - fulltext indexes
//
//Note:
// Hash, skiplist and persistent indexes can optionally be created in a sparse variant.
// Sparse indexes do not index documents for which any of the index attributes is either not set or is null.
type CreateIndexParameters struct {
	Fields      []string `json:"fields,omitempty"`
	Type        string   `json:"type,omitempty"`
	Sparse      bool     `json:"sparse,omitempty"`
	Unique      bool     `json:"unique,omitempty"`
	Deduplicate bool     `json:"deduplicate,omitempty"`
}

func (c CreateIndexParameters) Json() []byte {

	return nil
}
