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
	KindMatchNone      Kind = "match_none"
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
	KindScriptScore    Kind = "script_score"
)

var clauseHandlers = map[Kind]func() QueryClause{
	KindPrefix:  func() QueryClause { return &PrefixQuery{} },
	KindMatch:   func() QueryClause { return &MatchQuery{} },
	KindTerm:    func() QueryClause { return &TermQuery{} },
	KindTerms:   func() QueryClause { return &TermsQuery{} },
	KindBoolean: func() QueryClause { return &BooleanQuery{} },
	KindExists:  func() QueryClause { return &ExistsQuery{} },
	KindRange:   func() QueryClause { return &RangeQuery{} },
}
