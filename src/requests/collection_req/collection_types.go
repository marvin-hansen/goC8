package collection_req

type CollectionType uint8 // 255 possible values

const (
	UnknownCollectionType  CollectionType = 0
	DocumentCollectionType CollectionType = 2
	EdgeCollectionType     CollectionType = 3
)

func (e CollectionType) String() string {
	switch e {
	case DocumentCollectionType:
		return "DocumentCollectionType"
	case EdgeCollectionType:
		return "EdgeCollectionType"
	case UnknownCollectionType:
		return "UnknownCollectionType"
	default:
		return "UnknownCollectionType"
	}
}

func (e CollectionType) ToInt() int {
	switch e {
	case DocumentCollectionType:
		return 2
	case EdgeCollectionType:
		return 3
	case UnknownCollectionType:
		return 0
	default:
		return 0
	}
}

func (e CollectionType) FromString(str string) CollectionType {
	switch str {
	case "DocumentCollectionType":
		return DocumentCollectionType
	case "EdgeCollectionType":
		return EdgeCollectionType
	case "UnknownCollectionType":
		return UnknownCollectionType
	default:
		return UnknownCollectionType
	}
}

func (e CollectionType) FromInt(i int) CollectionType {
	switch i {
	case 2:
		return DocumentCollectionType
	case 3:
		return EdgeCollectionType
	case 0:
		return UnknownCollectionType
	default:
		return UnknownCollectionType
	}
}
