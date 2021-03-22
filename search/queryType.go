package search

type QueryType string

const (
	QueryTypePrefix         QueryType = "prefix"
	QueryTypeMatch          QueryType = "match"
	QueryTypeMatchAll       QueryType = "match_all"
	QueryTypeTerm           QueryType = "term"
	QueryTypeTerms          QueryType = "terms"
	QueryTypeRange          QueryType = "range"
	QueryTypeBoosting       QueryType = "boosting"
	QueryTypeBoolean        QueryType = "boolean"
	QueryTypeConstantScore  QueryType = "constant_score"
	QueryTypeFunctionScore  QueryType = "function_score"
	QueryTypeDisjunctionMax QueryType = "dis_max"
)

func (qt QueryType) String() string {
	return string(qt)
}
