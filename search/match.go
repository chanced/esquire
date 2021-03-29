package search

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

type Matcher interface {
	Match() (*MatchQuery, error)
}

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
	Name string
	// The field which is being matched.
	//
	// If you are setting Match explicitly, this does not need to be set. It
	// does, however, if you are adding it to a set of Clauses.
	Field string
	// (Required) Text, number, boolean or date you wish to find in the
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
	MaxExpansions interface{}
	// Number of beginning characters left unchanged for fuzzy matching.
	// Defaults to 0.
	PrefixLength interface{}
	// Minimum number of clauses that must match for a document to be returned
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-minimum-should-match.html
	MinimumShouldMatchParam string
	// Indicates whether no documents are returned if the analyzer removes all
	// tokens, such as when using a stop filter.
	ZeroTermsQuery ZeroTerms

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

func (m Match) name() string {
	return m.Name
}

func (m Match) field() string {
	return m.Field
}

func (m Match) Type() Type {
	return TypeMatch
}
func (m Match) Clause() (Clause, error) {
	return m.Match()
}
func (m Match) Match() (*MatchQuery, error) {
	v := &MatchQuery{
		field: m.Field,
	}
	err := v.SetQuery(m.Query)
	if err != nil {
		return nil, NewQueryError(err, TypeMatch, m.Field)
	}
	v.SetAnalyzer(m.Analyzer)
	v.SetAutoGenerateSynonymsPhraseQuery(!m.NoAutoGenerateSynonymsPhraseQuery)
	v.SetFuzziness(m.Fuzziness)
	err = v.SetFuzzyRewrite(m.FuzzyRewrite)
	if err != nil {
		return nil, NewQueryError(err, TypeMatch, m.Field)
	}
	v.SetFuzzyTranspositions(!m.NoFuzzyTranspositions)
	v.SetLenient(m.Lenient)
	err = v.SetMaxExpansions(m.MaxExpansions)
	if err != nil {
		return nil, NewQueryError(err, TypeMatch, m.Field)
	}
	err = v.SetPrefixLength(m.PrefixLength)
	if err != nil {
		return nil, NewQueryError(err, TypeMatch, m.Field)
	}
	err = v.SetZeroTermsQuery(m.ZeroTermsQuery)
	if err != nil {
		return nil, NewQueryError(err, TypeMatch, m.Field)
	}
	v.cutoffFrequency = m.CutoffFrequency
	return v, nil
}

// MatchQuery returns documents that match a provided text, number, date or
// boolean value. The provided text is analyzed before matching.
//
// The match query is the standard query for performing a full-text search,
// including options for fuzzy matching.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-query.html
type MatchQuery struct {
	field string
	query dynamic.StringNumberBoolOrTime

	nameParam
	lenientParam
	operatorParam
	analyzerParam
	fuzzinessParam
	prefixLengthParam
	maxExpansionsParam
	zeroTermsQueryParam
	cutoffFrequencyParam
	minimumShouldMatchParam
	fuzzyTranspositionsParam
	autoGenerateSynonymsPhraseQueryParam
}

func (m MatchQuery) Field() string {
	return m.field
}

func (m MatchQuery) SetField(field string) {
	m.field = field
}

func (m MatchQuery) IsEmpty() bool {
	return len(m.field) == 0 || m.query.IsEmptyString()
}

// SetQuery sets the Match's query param. It returns an error if it is nil or
// empty. If you need to clear match, use Clear()
func (m *MatchQuery) SetQuery(query interface{}) error {
	if query == nil {
		return ErrQueryRequired
	}

	return nil
}

func (m *MatchQuery) Query() *dynamic.StringNumberBoolOrTime {
	return &m.query
}

func (m MatchQuery) MarshalJSON() ([]byte, error) {
	if m.IsEmpty() {
		return dynamic.Null, nil
	}
	data, err := m.marshalClauseJSON()
	if err != nil {
		return nil, err
	}
	return json.Marshal(dynamic.Map{m.field: data})
}

func (m MatchQuery) marshalClauseJSON() (dynamic.JSON, error) {
	params, err := marshalParams(&m)
	if err != nil {
		return nil, err
	}
	params["query"] = m.query
	return json.Marshal(params)
}

func (m *MatchQuery) UnmarshalJSON(data []byte) error {
	*m = MatchQuery{}

	d := map[string]dynamic.JSON{}
	err := json.Unmarshal(data, &d)
	if err != nil {
		return err
	}
	for k, v := range d {
		m.field = k
		return m.unmarshalClauseJSON(v)
	}
	return nil
}

func (m *MatchQuery) unmarshalClauseJSON(data dynamic.JSON) error {
	fields, err := unmarshalParams(data, m)
	if err != nil {
		return err
	}
	if v, ok := fields["query"]; ok {
		var q dynamic.StringNumberBoolOrTime
		err := json.Unmarshal(v, &q)
		if err != nil {
			return err
		}
		m.query = q
	}
	return nil
}

func (m MatchQuery) Type() Type {
	return TypeMatch
}

func (m *MatchQuery) Set(field string, match Matcher) error {
	if match == nil {
		m.RemoveMatch()
		return nil
	}
	if field == "" {
		return NewQueryError(ErrFieldRequired, TypeTerm)
	}
	r, err := match.Match()
	if err != nil {
		return err
	}
	r.field = field
	*m = *r
	return nil
}
func (m *MatchQuery) RemoveMatch() {
	*m = MatchQuery{}
}
