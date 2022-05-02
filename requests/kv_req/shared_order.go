package kv_req

type Order string

const (
	Ascending  Order = "asc"
	Descending Order = "desc"
)

func (o Order) String() string {
	return string(o)
}
