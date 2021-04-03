package picker

type QueryKind string

func (t QueryKind) String() string {
	return string(t)
}

func (t QueryKind) IsValid() bool {
	_, ok := clauseHandlers[t]
	return ok
}

const (
	KindPrefix         QueryKind = "prefix"
	KindMatch          QueryKind = "match"
	KindMatchAll       QueryKind = "match_all"
	KindMatchNone      QueryKind = "match_none"
	KindTerm           QueryKind = "term"
	KindExists         QueryKind = "exists"
	KindTerms          QueryKind = "terms"
	KindRange          QueryKind = "range"
	KindBoosting       QueryKind = "boosting"
	KindBoolean        QueryKind = "boolean"
	KindConstantScore  QueryKind = "constant_score"
	KindFunctionScore  QueryKind = "function_score"
	KindDisjunctionMax QueryKind = "dis_max"
	KindAllOf          QueryKind = "all_of"
	KindFuzzy          QueryKind = "fuzzy"
	KindScriptScore    QueryKind = "script_score"
	KindScript         QueryKind = "script"
)

var clauseHandlers = map[QueryKind]func() QueryClause{
	KindPrefix:    func() QueryClause { return &PrefixQuery{} },
	KindMatch:     func() QueryClause { return &MatchQuery{} },
	KindTerm:      func() QueryClause { return &TermClauseQuery{} },
	KindTerms:     func() QueryClause { return &TermsQuery{} },
	KindBoolean:   func() QueryClause { return &BooleanQuery{} },
	KindExists:    func() QueryClause { return &ExistsQuery{} },
	KindRange:     func() QueryClause { return &RangeQuery{} },
	KindMatchAll:  func() QueryClause { return &MatchAllQuery{} },
	KindMatchNone: func() QueryClause { return &MatchNoneQuery{} },
	KindScript:    func() QueryClause { return &ScriptQuery{} },
}
