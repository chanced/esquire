package picker

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

type Matcher interface {
	Match() (*MatchQuery, error)
}
type CompleteMatcher interface {
	Matcher
	CompleteClauser
}

// MatchQueryParams returns documents that match a provided text, number, date or boolean
// value. The provided text is analyzed before matching.
//
// The match query is the standard query for performing a full-text search,
// including options for fuzzy matching.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-query.html
type MatchQueryParams struct {
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
	// The match query analyzes any provided text before performing a picker.
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

	completeClause
}

func (m MatchQueryParams) name() string {
	return m.Name
}

func (m MatchQueryParams) field() string {
	return m.Field
}

func (m MatchQueryParams) Kind() QueryKind {
	return QueryKindMatch
}
func (m MatchQueryParams) Clause() (QueryClause, error) {
	return m.Match()
}
func (m MatchQueryParams) Match() (*MatchQuery, error) {
	q := &MatchQuery{
		field: m.Field,
	}
	err := q.setQuery(m.Query)
	if err != nil {
		return q, newQueryError(err, QueryKindMatch, m.Field)
	}
	q.SetAnalyzer(m.Analyzer)
	q.SetAutoGenerateSynonymsPhraseQuery(!m.NoAutoGenerateSynonymsPhraseQuery)
	q.SetFuzziness(m.Fuzziness)
	err = q.SetFuzzyRewrite(m.FuzzyRewrite)
	if err != nil {
		return q, newQueryError(err, QueryKindMatch, m.Field)
	}
	q.SetFuzzyTranspositions(!m.NoFuzzyTranspositions)
	q.SetLenient(m.Lenient)
	err = q.SetMaxExpansions(m.MaxExpansions)
	if err != nil {
		return q, newQueryError(err, QueryKindMatch, m.Field)
	}
	err = q.SetPrefixLength(m.PrefixLength)
	if err != nil {
		return q, newQueryError(err, QueryKindMatch, m.Field)
	}
	err = q.SetZeroTermsQuery(m.ZeroTermsQuery)
	if err != nil {
		return q, newQueryError(err, QueryKindMatch, m.Field)
	}
	q.cutoffFrequency = m.CutoffFrequency
	return q, nil
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
	completeClause
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

func (m *MatchQuery) Clause() (QueryClause, error) {
	return m, nil
}

func (m *MatchQuery) Field() string {
	if m == nil {
		return ""
	}
	return m.field
}

func (m *MatchQuery) IsEmpty() bool {
	return m == nil || len(m.field) == 0 || m.query.IsEmptyString()
}

func (m *MatchQuery) Set(field string, match Matcher) error {
	if match == nil {
		m.Clear()
		return nil
	}
	if field == "" {
		return newQueryError(ErrFieldRequired, QueryKindTerm)
	}
	r, err := match.Match()
	if err != nil {
		return err
	}
	r.field = field
	*m = *r
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
	params, err := marshalClauseParams(&m)
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
	fields, err := unmarshalClauseParams(data, m)
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

func (m MatchQuery) Kind() QueryKind {
	return QueryKindMatch
}
func (m *MatchQuery) Clear() {
	*m = MatchQuery{}
}

// setQuery sets the Match's query param. It returns an error if it is nil or
// empty. If you need to clear match, use Clear()
func (m *MatchQuery) setQuery(query interface{}) error {
	if query == nil {
		return ErrQueryRequired
	}
	return nil
}
func (m MatchQuery) setField(field string) {
	m.field = field
}
