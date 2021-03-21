package search

import "github.com/chanced/dynamic"

// Match returns documents that match a provided text, number, date or boolean
// value. The provided text is analyzed before matching.
//
// The match query is the standard query for performing a full-text search,
// including options for fuzzy matching.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-query.html
type Match struct {
	Query                             interface{}
	Analyzer                          string
	NoAutoGenerateSynonymsPhraseQuery bool
	NoFuzzyTranspositions             bool
	Fuzziness                         string
	FuzzyRewrite                      Rewrite
	Lenient                           bool
	Operator                          Operator
	MaxExpansions                     int
	PrefixLength                      int
	MinimumShouldMatchParam           string
	ZeroTermsQuery                    ZeroTermsQuery
}

func (m Match) MatchQueryValue() (MatchQueryValue, error) {
	v := MatchQueryValue{}
	err := v.SetQuery(v)
	if err != nil {
		return v, err
	}
	v.SetAnalyzer(m.Analyzer)
	v.SetAutoGenerateSynonymsPhraseQuery(!m.NoAutoGenerateSynonymsPhraseQuery)
	v.SetFuzziness(m.Fuzziness)
	err = v.SetFuzzyRewrite(m.FuzzyRewrite)
	if err != nil {
		return v, err
	}
	v.SetFuzzyTranspositions(!m.NoFuzzyTranspositions)
	v.SetLenient(m.Lenient)
	if m.MaxExpansions != 0 {
		v.SetMaxExpansions(m.MaxExpansions)
	}
	if m.PrefixLength != 0 {
		v.SetPrefixLength(m.PrefixLength)
	}
	if m.ZeroTermsQuery != "" {
		v.SetZeroTermsQuery(m.ZeroTermsQuery)
	}
	return v, nil
}

// MatchQueryValue returns documents that match a provided text, number, date or boolean
// value. The provided text is analyzed before matching.
//
// The match query is the standard query for performing a full-text search,
// including options for fuzzy matching.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-query.html
type MatchQueryValue struct {
	// (Required) Text, number, boolean value or date you wish to find in the
	// provided <field>.
	//
	// The match query analyzes any provided text before performing a search.
	// This means the match query can search text fields for analyzed tokens
	// rather than an exact term.
	Query dynamic.StringNumberBoolOrTime `json:"query" bson:"query"`

	// Analyzer used to convert the text in the query value into tokens.
	// Defaults to the index-time analyzer mapped for the <field>. If no
	// analyzer is mapped, the index’s default analyzer is used.
	AnalyzerParam `json:",inline" bson:",inline"`

	// If true, match phrase queries are automatically created for multi-term
	// synonyms. Defaults to true.
	AutoGenerateSynonymsPhraseQueryParam `json:",inline" bson:",inline"`

	// Maximum edit distance allowed for matching.
	FuzzinessParam `json:",inline" bson:",inline"`

	// Maximum number of terms to which the query will expand. Defaults to 50.
	MaxExpansionsParam `json:",inline" bson:",inline"`

	// Number of beginning characters left unchanged for fuzzy matching. Defaults to 0.
	PrefixLengthParam `json:",inline" bson:",inline"`

	// If true, edits for fuzzy matching include transpositions of two adjacent
	// characters (ab → ba). Defaults to true.
	FuzzyTranspositionsParam `json:",inline" bson:",inline"`

	//  If true, format-based errors, such as providing a text query value for a
	//  numeric field, are ignored. Defaults to false.
	LenientParam `json:",inline" bson:",inline"`

	// Boolean logic used to interpret text in the query value.
	OperatorParam `json:",inline" bson:",inline"`

	// Minimum number of clauses that must match for a document to be returned
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-minimum-should-match.html
	MinimumShouldMatchParam `json:",inline" bson:",inline"`
	// Indicates whether no documents are returned if the analyzer removes all
	// tokens, such as when using a stop filter.
	ZeroTermsQueryParam `json:",inline" bson:",inline"`
}

func (mq *MatchQueryValue) SetQuery(value interface{}) error {

	if snbt, ok := value.(dynamic.StringNumberBoolOrTime); ok {
		mq.Query = snbt
		return nil
	}
	if snbt, ok := value.(*dynamic.StringNumberBoolOrTime); ok {
		mq.Query = *snbt
		return nil
	}
	return mq.Query.Set(value)
}

func NewMatchQuery() MatchQuery {
	return MatchQuery{
		MatchQueryValue: map[string]MatchQueryValue{},
	}
}

// MatchQuery returns documents that match a provided text, number, date or
// boolean value. The provided text is analyzed before matching.
//
// The match query is the standard query for performing a full-text search,
// including options for fuzzy matching.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-query.html
type MatchQuery struct {
	MatchQueryValue map[string]MatchQueryValue `json:"match,omitempty" bson:"match,omitempty"`
}

// AddMatch returns an error if the field already exists. Use SetMatch to overwrite.
func (m *MatchQuery) AddMatch(field string, match Match) error {
	if m.MatchQueryValue == nil {
		m.MatchQueryValue = map[string]MatchQueryValue{}
	}
	if _, exists := m.MatchQueryValue[field]; exists {
		return QueryError{Err: ErrFieldExists, Field: field, QueryType: QueryTypeMatch}
	}
	return m.SetMatch(field, match)
}

func (m *MatchQuery) SetMatch(field string, match Match) error {
	if m.MatchQueryValue == nil {
		m.MatchQueryValue = map[string]MatchQueryValue{}
	}
	q, err := match.MatchQueryValue()
	if err != nil {
		return QueryError{Err: err, Field: field, QueryType: QueryTypeMatch}
	}
	m.MatchQueryValue[field] = q
	return nil
}

func (m *MatchQuery) RemoveMatch(field string) {
	delete(m.MatchQueryValue, field)
}
