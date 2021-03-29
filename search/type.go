package search

type Kind string

func (t Kind) String() string {
	return string(t)
}

func (t Kind) IsValid() bool {
	_, ok := clauseHandlers[t]
	return ok
}

const (
	KindPrefix         Kind = "prefix"
	KindMatch          Kind = "match"
	KindMatchAll       Kind = "match_all"
	KindTerm           Kind = "term"
	KindExists         Kind = "exists"
	KindTerms          Kind = "terms"
	KindRange          Kind = "range"
	KindBoosting       Kind = "boosting"
	KindBoolean        Kind = "boolean"
	KindConstantScore  Kind = "constant_score"
	KindFunctionScore  Kind = "function_score"
	KindDisjunctionMax Kind = "dis_max"
	KindAllOf          Kind = "all_of"
	KindFuzzy          Kind = "fuzzy"
)

var clauseHandlers = map[Kind]func() Clause{
	KindPrefix:  func() Clause { return &PrefixQuery{} },
	KindMatch:   func() Clause { return &MatchQuery{} },
	KindTerm:    func() Clause { return &TermQuery{} },
	KindTerms:   func() Clause { return &TermsQuery{} },
	KindBoolean: func() Clause { return &BooleanQuery{} },
	KindExists:  func() Clause { return &ExistsQuery{} },
}
