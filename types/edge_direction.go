package types

type EdgeDirection string

const (
	ANY EdgeDirection = "" // If not set, all edges are returned.
	IN  EdgeDirection = "in"
	OUT EdgeDirection = "out"
)

func (e EdgeDirection) String() string {
	return string(e)
}
