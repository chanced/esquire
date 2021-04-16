package picker

import (
	"encoding/json"
)

type MultiMatcher interface {
	MultiMatch() (*MultiMatchQuery, error)
}

type MultiMatchQueryParams struct {
	// Each query accepts a _name in its top level definition. You can use named
	// queries to track which queries matched returned documents. If named
	// queries are used, the response includes a matched_queries property for
	// each hit.
	Name string
	// The fields which are being matched.
	Fields []string
	// (Required) The query string
	Query string
	// Analyzer used to convert the text in the query value into tokens.
	// Defaults to the index-time analyzer mapped for the <field>. If no
	// analyzer is mapped, the index’s default analyzer is used.
	Analyzer string
	// If true, match phrase queries are NOT automatically created for
	// multi-term synonyms.
	//
	// If true, auto_generate_synonyms_phrase_query is set to false
	AutoGenerateSynonymsPhraseQuery interface{}
	// If true, edits for fuzzy matching DO NOT include transpositions of two
	// adjacent characters (ab → ba).
	FuzzyTranspositions interface{}
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
	CutoffFrequency interface{}
	completeClause
}

func (MultiMatchQueryParams) Kind() QueryKind {
	return QueryKindMultiMatch
}
func (p MultiMatchQueryParams) Clause() (QueryClause, error) {
	return p.MultiMatch()
}
func (p MultiMatchQueryParams) MultiMatch() (*MultiMatchQuery, error) {
	q := &MultiMatchQuery{}
	err := q.SetQuery(p.Query)
	if err != nil {
		return q, newQueryError(err, QueryKindMultiMatch, p.Fields...)
	}
	err = q.SetFields(p.Fields)
	if err != nil {
		return q, newQueryError(err, QueryKindMultiMatch, p.Fields...)
	}
	q.SetAnalyzer(p.Analyzer)
	err = q.SetAutoGenerateSynonymsPhraseQuery(p.AutoGenerateSynonymsPhraseQuery)
	if err != nil {
		return q, err
	}
	q.SetFuzziness(p.Fuzziness)

	err = q.SetFuzzyRewrite(p.FuzzyRewrite)
	if err != nil {
		return q, newQueryError(err, QueryKindMultiMatch, p.Fields...)
	}
	err = q.SetFuzzyTranspositions(p.FuzzyTranspositions)
	if err != nil {
		return q, err
	}
	q.SetLenient(p.Lenient)
	err = q.SetMaxExpansions(p.MaxExpansions)
	if err != nil {
		return q, newQueryError(err, QueryKindMultiMatch, p.Fields...)
	}
	err = q.SetPrefixLength(p.PrefixLength)
	if err != nil {
		return q, newQueryError(err, QueryKindMultiMatch, p.Fields...)
	}
	err = q.SetZeroTermsQuery(p.ZeroTermsQuery)
	if err != nil {
		return q, newQueryError(err, QueryKindMultiMatch, p.Fields...)
	}
	err = q.cutoffFrequency.Set(p.CutoffFrequency)
	if err != nil {
		return q, newQueryError(err, QueryKindMultiMatch, p.Fields...)
	}
	return q, nil
}

type MultiMatchQuery struct {
	query  string
	fields []string
	tieBreakerParam
	nameParam
	minimumShouldMatchParam
	analyzerParam
	boostParam
	fuzzinessParam
	lenientParam
	prefixLengthParam
	maxExpansionsParam
	zeroTermsQueryParam
	cutoffFrequencyParam
	operatorParam
	autoGenerateSynonymsPhraseQueryParam
	multiMatchTypeParam
	fuzzyTranspositionsParam
	completeClause
}

func (q *MultiMatchQuery) Clause() (QueryClause, error) {
	return q, nil
}
func (q *MultiMatchQuery) MultiMatch() (*MultiMatchQuery, error) {
	return q, nil
}

func (q MultiMatchQuery) Query() string {
	return q.query
}
func (q *MultiMatchQuery) SetQuery(query string) error {
	if len(query) == 0 {
		return ErrQueryRequired
	}
	q.query = query
	return nil
}

func (q *MultiMatchQuery) SetFields(fields []string) error {
	if len(fields) == 0 {
		return ErrFieldRequired
	}
	q.fields = fields
	return nil
}
func (q MultiMatchQuery) Fields() []string {
	return q.fields
}
func (q *MultiMatchQuery) Clear() {
	if q == nil {
		return
	}
	*q = MultiMatchQuery{}
}
func (q *MultiMatchQuery) IsEmpty() bool {
	return q == nil || len(q.fields) == 0 || len(q.query) == 0
}
func (q MultiMatchQuery) MarshalBSON() ([]byte, error) {
	return q.MarshalJSON()
}

func (q MultiMatchQuery) MarshalJSON() ([]byte, error) {
	params, err := marshalClauseParams(&q)
	if err != nil {
		return nil, err
	}
	params["query"] = q.query
	if q.typ != MultiMatchTypeUnspecified {
		params["type"] = q.typ
	}
	params["fields"] = q.fields
	return json.Marshal(params)
}

func (MultiMatchQuery) Kind() QueryKind {
	return QueryKindMultiMatch
}

func (q *MultiMatchQuery) UnmarshalBSON(data []byte) error {
	return q.UnmarshalJSON(data)
}

func (q *MultiMatchQuery) UnmarshalJSON(data []byte) error {
	*q = MultiMatchQuery{}

	params, err := unmarshalClauseParams(data, q)
	if err != nil {
		return err
	}
	if v, ok := params["query"]; ok {
		var query string
		err := json.Unmarshal(v, &query)
		if err != nil {
			return err
		}
		q.query = query
	}
	if v, ok := params["fields"]; ok {
		var fields []string
		err := json.Unmarshal(v, &fields)
		if err != nil {
			return err
		}
		q.fields = fields
	}
	if v, ok := params["type"]; ok {
		var typ MultiMatchType
		err := json.Unmarshal(v, &typ)
		if err != nil {
			return err
		}
	}
	return nil
}
