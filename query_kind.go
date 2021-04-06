package picker

type QueryKind string

func (t QueryKind) String() string {
	return string(t)
}

func (t QueryKind) IsValid() bool {
	_, ok := queryKindHandlers[t]
	return ok
}

const (
	QueryKindPrefix            QueryKind = "prefix"
	QueryKindMatch             QueryKind = "match"
	QueryKindMatchAll          QueryKind = "match_all"
	QueryKindMatchNone         QueryKind = "match_none"
	QueryKindTerm              QueryKind = "term"
	QueryKindExists            QueryKind = "exists"
	QueryKindTerms             QueryKind = "terms"
	QueryKindRange             QueryKind = "range"
	QueryKindBoosting          QueryKind = "boosting"
	QueryKindBoolean           QueryKind = "bool"
	QueryKindConstantScore     QueryKind = "constant_score"
	QueryKindFunctionScore     QueryKind = "function_score"
	QueryKindDisjunctionMax    QueryKind = "dis_max"
	QueryKindAllOf             QueryKind = "all_of"
	QueryKindFuzzy             QueryKind = "fuzzy"
	QueryKindScriptScore       QueryKind = "script_score"
	QueryKindScript            QueryKind = "script"
	QueryKindIDs               QueryKind = "ids"
	QueryKindIntervals         QueryKind = "intervals"
	QueryKindMatchBoolPrefix   QueryKind = "match_bool_prefix"
	QueryKindMatchPhrase       QueryKind = "match_phrase"
	QueryKindMatchPhrasePrefix QueryKind = "match_phrase_prefix"
)

var queryKindHandlers = map[QueryKind]func() QueryClause{
	QueryKindPrefix:            func() QueryClause { return &PrefixQuery{} },
	QueryKindMatch:             func() QueryClause { return &MatchQuery{} },
	QueryKindTerm:              func() QueryClause { return &TermQuery{} },
	QueryKindTerms:             func() QueryClause { return &TermsQuery{} },
	QueryKindBoolean:           func() QueryClause { return &BoolQuery{} },
	QueryKindExists:            func() QueryClause { return &ExistsQuery{} },
	QueryKindRange:             func() QueryClause { return &RangeQuery{} },
	QueryKindMatchAll:          func() QueryClause { return &MatchAllQuery{} },
	QueryKindMatchNone:         func() QueryClause { return &MatchNoneQuery{} },
	QueryKindScript:            func() QueryClause { return &ScriptQuery{} },
	QueryKindBoosting:          func() QueryClause { return &BoostingQuery{} },
	QueryKindConstantScore:     func() QueryClause { return &ConstantScoreQuery{} },
	QueryKindIDs:               func() QueryClause { return &IDsQuery{} },
	QueryKindIntervals:         func() QueryClause { return &IntervalsQuery{} },
	QueryKindMatchPhrase:       func() QueryClause { return &MatchPhraseQuery{} },
	QueryKindMatchPhrasePrefix: func() QueryClause { return &MatchPhrasePrefixQuery{} },
}
