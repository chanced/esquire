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
	TypeAllOf          Type = "all_of"
)

var TypeHandlers = map[Type]func() Rule{
	TypePrefix: func() Rule { return &PrefixQueryValue{} },
	TypeMatch:  func() Rule { return &MatchRule{} },
}

func (qt Type) String() string {
	return string(qt)
}
