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
	KindPrefix:    func() QueryClause { return &PrefixClause{} },
	KindMatch:     func() QueryClause { return &MatchClause{} },
	KindTerm:      func() QueryClause { return &TermClause{} },
	KindTerms:     func() QueryClause { return &TermsClause{} },
	KindBoolean:   func() QueryClause { return &BooleanClause{} },
	KindExists:    func() QueryClause { return &ExistsClause{} },
	KindRange:     func() QueryClause { return &RangeClause{} },
	KindMatchAll:  func() QueryClause { return &MatchAllClause{} },
	KindMatchNone: func() QueryClause { return &MatchNoneClause{} },
	KindScript:    func() QueryClause { return &ScriptClause{} },
}
