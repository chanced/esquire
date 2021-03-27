package search

type Type string

func (t Type) String() string {
	return string(t)
}

func (t Type) IsValid() bool {
	_, ok := clauseHandlers[t]
	return ok
}

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

var clauseHandlers = map[Type]func() Clause{
	TypePrefix:  func() Clause { return &PrefixQuery{} },
	TypeMatch:   func() Clause { return &MatchQuery{} },
	TypeTerm:    func() Clause { return &TermQuery{} },
	TypeTerms:   func() Clause { return &TermsQuery{} },
	TypeBoolean: func() Clause { return &BooleanQuery{} },
}
