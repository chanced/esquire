package search

type Type string

const (
	TypePrefix         Type = "prefix"
	TypeMatch          Type = "match"
	TypeMatchAll       Type = "match_all"
	TypeTerm           Type = "term"
	TypeTerms          Type = "terms"
	TypeRange          Type = "range"
	TypeBoosting       Type = "boosting"
	TypeBoolean        Type = "boolean"
	TypeConstantScore  Type = "constant_score"
	TypeFunctionScore  Type = "function_score"
	TypeDisjunctionMax Type = "dis_max"
)

var TypeHandlers = map[Type]func() Statement{
	TypePrefix: func() Statement { return &PrefixQueryValue{} },
	TypeMatch:  func() Statement { return &MatchQueryValue{} },
}

func (qt Type) String() string {
	return string(qt)
}
