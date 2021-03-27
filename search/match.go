package search

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

// Match returns documents that match a provided text, number, date or boolean
// value. The provided text is analyzed before matching.
//
// The match query is the standard query for performing a full-text search,
// including options for fuzzy matching.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-query.html
type Match struct {
	// Each query accepts a _name in its top level definition. You can use named
	// queries to track which queries matched returned documents. If named
	// queries are used, the response includes a matched_queries property for
	// each hit.
	QueryName string
	// The field which is being matched.
	//
	// If you are setting Match explicitly, this does not need to be set. It
	// does, however, if you are adding it to a set of Clauses.
	FieldName string
	// (Required) Text, number, boolean value or date you wish to find in the
	// provided <field>.
	//
	// The match query analyzes any provided text before performing a search.
	// This means the match query can search text fields for analyzed tokens
	// rather than an exact term.
	Query interface{}
	// Analyzer used to convert the text in the query value into tokens.
	// Defaults to the index-time analyzer mapped for the <field>. If no
	// analyzer is mapped, the index’s default analyzer is used.
	Analyzer string
	// If true, match phrase queries are NOT automatically created for
	// multi-term synonyms.
	//
	// If true, auto_generate_synonyms_phrase_query is set to false
	NoAutoGenerateSynonymsPhraseQuery bool
	// If true, edits for fuzzy matching DO NOT include transpositions of two
	// adjacent characters (ab → ba).
	//
	// if true, fuzzy_transpositions is set to false
	NoFuzzyTranspositions bool
	// Maximum edit distance allowed for matching.
	Fuzziness    string
	FuzzyRewrite Rewrite
	//  If true, format-based errors, such as providing a text query value for a
	//  numeric field, are ignored. Defaults to false.
	Lenient bool
	// Boolean logic used to interpret text in the query value. Defaults to OR
	Operator Operator
	// Maximum number of terms to which the query will expand. Defaults to 50.
	MaxExpansions dynamic.Number
	// Number of beginning characters left unchanged for fuzzy matching.
	// Defaults to 0.
	PrefixLength dynamic.Number
	// Minimum number of clauses that must match for a document to be returned
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-minimum-should-match.html
	MinimumShouldMatchParam string
	// Indicates whether no documents are returned if the analyzer removes all
	// tokens, such as when using a stop filter.
	ZeroTermsQuery ZeroTermsQuery

	//
	// The match query supports a cutoff_frequency that allows specifying an
	// absolute or relative document frequency where high frequency terms are
	// moved into an optional subquery and are only scored if one of the low
	// frequency (below the cutoff) terms in the case of an or operator or all
	// of the low frequency terms in the case of an and operator match.
	//
	// DEPRECATED in 7.3.0
	//
	// This option can be omitted as the Match can skip blocks of documents
	// efficiently, without any configuration, provided that the total number of
	// hits is not tracked.
	CutoffFrequency dynamic.Number
}

func (m Match) Name() string {
	return m.QueryName
}

func (m Match) SetName(name string) {
	m.QueryName = name
}

func (m Match) Type() Type {
	return TypeMatch
}
func (m Match) Clause() (Clause, error) {
	return m.Match()
}
func (m Match) Match() (*matchClause, error) {
	v := &matchClause{}
	err := v.SetQuery(m.Query)
	if err != nil {
		return nil, err
	}
	v.SetAnalyzer(m.Analyzer)
	v.SetAutoGenerateSynonymsPhraseQuery(!m.NoAutoGenerateSynonymsPhraseQuery)
	v.SetFuzziness(m.Fuzziness)
	err = v.SetFuzzyRewrite(m.FuzzyRewrite)
	if err != nil {
		return nil, err
	}
	v.SetFuzzyTranspositions(!m.NoFuzzyTranspositions)
	v.SetLenient(m.Lenient)
	if n, ok := m.MaxExpansions.Int(); ok {
		v.SetMaxExpansions(n)
	}
	if n, ok := m.PrefixLength.Int(); ok {
		v.SetPrefixLength(n)
	}
	if m.ZeroTermsQuery != "" {
		v.SetZeroTermsQuery(m.ZeroTermsQuery)
	}
	v.cutoffFrequency = m.CutoffFrequency
	return v, nil
}

type matchClause struct {
	MatchQueryValue dynamic.StringNumberBoolOrTime
	analyzerParam
	nameParam
	autoGenerateSynonymsPhraseQueryParam
	fuzzinessParam
	maxExpansionsParam
	prefixLengthParam
	fuzzyTranspositionsParam
	lenientParam
	operatorParam
	minimumShouldMatchParam
	zeroTermsQueryParam
	cutoffFrequencyParam
}

func (mr *matchClause) Type() Type {
	return TypeMatch
}
func (mr matchClause) HasMatchRule() bool {
	return !mr.MatchQueryValue.IsEmptyString()
}
func (mr *matchClause) SetQuery(value interface{}) error {
	return mr.MatchQueryValue.Set(value)
}

func (mr matchClause) MarshalJSON() ([]byte, error) {
	if !mr.HasMatchRule() {
		return dynamic.Null, nil
	}
	m, err := marshalParams(&mr)
	if err != nil {
		return nil, err
	}
	m["query"] = mr.MatchQueryValue
	return json.Marshal(m)
}

func (mr *matchClause) UnmarshalJSON(data []byte) error {
	mr.MatchQueryValue = dynamic.NewStringNumberBoolOrTime()
	mr.analyzerParam = analyzerParam{}
	mr.autoGenerateSynonymsPhraseQueryParam = autoGenerateSynonymsPhraseQueryParam{}
	mr.fuzzinessParam = fuzzinessParam{}
	mr.fuzzyTranspositionsParam = fuzzyTranspositionsParam{}
	mr.lenientParam = lenientParam{}
	mr.prefixLengthParam = prefixLengthParam{}
	mr.minimumShouldMatchParam = minimumShouldMatchParam{}
	mr.operatorParam = operatorParam{}
	mr.maxExpansionsParam = maxExpansionsParam{}
	mr.zeroTermsQueryParam = zeroTermsQueryParam{}
	mr.cutoffFrequencyParam = cutoffFrequencyParam{}
	fields, err := unmarshalParams(data, mr)
	if err != nil {
		return err
	}

	if v, ok := fields["query"]; ok {
		mr.MatchQueryValue = dynamic.NewStringNumberBoolOrTime(v.UnquotedString())

	} else {
		mr.MatchQueryValue = dynamic.StringNumberBoolOrTime{}
	}
	return nil
}

// MatchQuery returns documents that match a provided text, number, date or
// boolean value. The provided text is analyzed before matching.
//
// The match query is the standard query for performing a full-text search,
// including options for fuzzy matching.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-query.html
type MatchQuery struct {
	MatchField string
	matchClause
}

func (mq MatchQuery) MarshalJSON() ([]byte, error) {
	if !mq.HasMatchRule() {
		return dynamic.Null, nil
	}
	m, err := marshalParams(&mq)
	if err != nil {
		return nil, err
	}
	m["query"] = mq.MatchQueryValue

	return json.Marshal(dynamic.Map{mq.MatchField: m})
}

func (mq *MatchQuery) UnmarshalJSON(data []byte) error {
	mq.MatchField = ""
	mq.matchClause = matchClause{}

	m := map[string]json.RawMessage{}
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		mq.MatchField = k
		err := json.Unmarshal(v, &mq.matchClause)
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

func (mq MatchQuery) Type() Type {
	return TypeMatch
}

func (mq *MatchQuery) SetMatch(field string, match *Match) error {
	if match == nil {
		mq.RemoveMatch()
		return nil
	}
	if field == "" {
		return NewQueryError(ErrFieldRequired, TypeTerm)
	}
	r, err := match.Match()
	if err != nil {
		return err
	}
	mq.MatchField = field
	mq.matchClause = *r
	return nil
}
func (mq *MatchQuery) RemoveMatch() {
	mq.MatchField = ""
	mq.matchClause = matchClause{}
}
