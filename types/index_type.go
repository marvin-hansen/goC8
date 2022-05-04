package types

type IndexType uint8 // 255 possible values

const (
	UnknownIndexType IndexType = iota
	FulltextIndex
	HashIndex
	GeoIndex
	PersistentIndex
	SkipListIndex
	TTLIndex
)

func (e IndexType) String() string {
	switch e {
	case FulltextIndex:
		return "FulltextIndex"
	case HashIndex:
		return "HashIndex"
	case GeoIndex:
		return "GeoIndex"
	case PersistentIndex:
		return "PersistentIndex"
	case SkipListIndex:
		return "SkipListIndex"
	case TTLIndex:
		return "TTLIndex"
	case UnknownIndexType:
		return "UnknownIndexType"
	default:
		return "UnknownIndexType"
	}
}
