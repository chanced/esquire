package search

type QueryType string

const (
	QueryTypePrefix QueryType = "prefix"
	QueryTypeMatch  QueryType = "match"
)

func (qt QueryType) String() string {
	return string(qt)
}
