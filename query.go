package picker

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

type Querier interface {
	Query() (*QueryValues, error)
}

type Query struct {

	// Term returns documents that contain an exact term in a provided field.
	//
	// You can use the term query to find documents based on a precise value such as
	// a price, a product ID, or a username.
	//
	// Avoid using the term query for text fields.
	//
	// By default, Elasticsearch changes the values of text fields as part of
	// analysis. This can make finding exact matches for text field values
	// difficult.
	//
	// To search text field values, use the match query instead.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-term-query.html
	Term *TermQuery

	// Terms returns documents that contain one or more exact terms in a provided
	// field.
	//
	// The terms query is the same as the term query, except you can search for
	// multiple values.
	Terms Termser

	// Match returns documents that match a provided text, number, date or boolean
	// value. The provided text is analyzed before matching.
	//
	// The match query is the standard query for performing a full-text search,
	// including options for fuzzy matching.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-query.html
	Match *MatchQuery

	// Boolean is a query that matches documents matching boolean combinations
	// of other queries. The bool query maps to Lucene BooleanQuery. It is built
	// using one or more boolean clauses, each clause with a typed occurrence.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-bool-query.html
	Boolean *BooleanQuery

	// Fuzzy returns documents that contain terms similar to the search term,
	// as measured by a Levenshtein edit distance.
	//
	// An edit distance is the number of one-character changes needed to turn one
	// term into another. These changes can include:
	//
	//      - Changing a character (box → fox)
	//
	//      - Removing a character (black → lack)
	//
	//      - Inserting a character (sic → sick)
	//
	//      - Transposing two adjacent characters (act → cat)
	//
	// To find similar terms, the fuzzy query creates a set of all possible
	// variations, or expansions, of the search term within a specified edit
	// distance. The query then returns exact matches for each expansion.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-fuzzy-query.html
	Fuzzy *FuzzyQuery

	// Prefix returns documents that contain a specific prefix in a provided field.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-prefix-query.html
	Prefix *Prefix

	// FunctionScore  allows you to modify the score of documents that are retrieved
	// by a query. This can be useful if, for example, a score function is
	// computationally expensive and it is sufficient to compute the score on a
	// filtered set of documents.
	//
	// To use function_score, the user has to define a query and one or more
	// functions, that compute a new score for each document returned by the query.
	FunctionScore *FunctionScoreQuery

	// ScoreScript uses a script to provide a custom score for returned documents.
	//
	// The script_score query is useful if, for example, a scoring function is
	// expensive and you only need to calculate the score of a filtered set of
	// documents.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-script-score-query.html
	ScriptScore *ScriptScoreQuery

	// Range returns documents that contain terms within a provided range.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-range-query.html
	Range *Range

	// MatchAll matches all documents, giving them all a _score of 1.0.
	MatchAll *MatchAll

	// MatchNone is the inverse of the match_all query, which matches no documents.
	MatchNone *MatchNone

	// Exists returns documents that contain an indexed value for a field.
	//
	// An indexed value may not exist for a document’s field due to a variety of
	// reasons:
	//
	//      - The field in the source JSON is null or []
	//
	//      - The field has "index" : false set in the mapping
	//
	//      - The length of the field value exceeded an ignore_above setting in the
	// mapping
	//
	//      - The field value was malformed and ignore_malformed was defined in the
	// mapping
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-exists-query.html
	Exists *Exists
}

func (q *Query) fuzzyClause() (*FuzzyClause, error) {
	if q == nil || q.Fuzzy == nil {
		return nil, nil
	}
	return q.Fuzzy.Fuzzy()
}

func (q *Query) termClause() (*TermClause, error) {
	if q == nil || q.Term == nil {
		return nil, nil
	}
	return q.Term.Term()
}

func (q *Query) termsClause() (*TermsClause, error) {
	if q == nil || q.Terms == nil {
		return nil, nil
	}
	return q.Terms.Terms()
}

func (q *Query) rangeClause() (*RangeClause, error) {
	if q == nil || q.Range == nil {
		return nil, nil
	}
	return q.Range.Range()
}

func (q *Query) prefixClause() (*PrefixClause, error) {
	if q == nil || q.Prefix == nil {
		return nil, nil
	}
	return q.Prefix.Prefix()
}

func (q *Query) matchClause() (*MatchClause, error) {
	if q == nil || q.Match == nil {
		return nil, nil
	}
	return q.Match.Match()

}

func (q *Query) scriptScoreClause() (*ScriptScoreClause, error) {
	if q == nil || q.ScriptScore == nil {
		return nil, nil
	}
	return q.ScriptScore.ScriptScore()
}

func (q *Query) functionScoreClause() (*FunctionScoreClause, error) {
	if q == nil || q.FunctionScore == nil {
		return nil, nil
	}
	return q.FunctionScore.FunctionScore()
}

func (q *Query) matchAllClause() (*MatchAllClause, error) {
	if q == nil || q.MatchAll == nil {
		return nil, nil
	}
	return q.MatchAll.MatchAll()
}
func (q *Query) matchNoneClause() (*MatchNoneClause, error) {
	if q == nil || q.MatchNone == nil {
		return nil, nil
	}
	return q.MatchNone.MatchNone()
}

func (q *Query) existsClause() (*ExistsClause, error) {
	if q == nil || q.Exists == nil {
		return nil, nil
	}
	return q.Exists.Exists()
}

func (q *Query) Query() (*QueryValues, error) {
	boolean, err := q.Boolean.Boolean()
	if err != nil {
		return nil, err
	}
	exists, err := q.existsClause()
	if err != nil {
		return nil, err
	}
	term, err := q.termClause()
	if err != nil {
		return nil, err
	}
	terms, err := q.termsClause()
	if err != nil {
		return nil, err
	}
	rng, err := q.rangeClause()
	if err != nil {
		return nil, err
	}
	prefix, err := q.prefixClause()
	if err != nil {
		return nil, err
	}
	match, err := q.matchClause()
	if err != nil {
		return nil, err
	}
	matchAll, err := q.matchAllClause()
	if err != nil {
		return nil, err
	}
	matchNone, err := q.matchNoneClause()
	if err != nil {
		return nil, err
	}
	scriptScore, err := q.scriptScoreClause()
	if err != nil {
		return nil, err
	}
	fuzzy, err := q.fuzzyClause()
	if err != nil {
		return nil, err
	}
	funcScore, err := q.functionScoreClause()
	if err != nil {
		return nil, err
	}

	qv := &QueryValues{
		matchClause:         match,
		existsClause:        exists,
		scriptScoreClause:   scriptScore,
		fuzzyClause:         fuzzy,
		booleanClause:       boolean,
		termClause:          term,
		termsClause:         terms,
		rangeClause:         rng,
		prefixClause:        prefix,
		matchAllClause:      matchAll,
		matchNoneClause:     matchNone,
		functionScoreClause: funcScore,
	}
	return qv, nil
}

// QueryValues defines the search definition using the ElasticSearch QueryValues DSL
//
// Elasticsearch provides a full QueryValues DSL (Domain Specific Language) based on
// JSON to define queries. Think of the QueryValues DSL as an AST (Abstract Syntax
// Tree) of queries, consisting of two types of clauses:
//
// Leaf query clauses
//
// Leaf query clauses look for a particular value in a particular field, such as
// the match, term or range queries. These queries can be used by themselves.
//
// Compound query clauses
//
// Compound query clauses wrap other leaf or compound queries and are used to
// combine multiple queries in a logical fashion (such as the bool or dis_max
// query), or to alter their behaviour (such as the constant_score query).
//
// QueryValues clauses behave differently depending on whether they are used in query
// context or filter context.
type QueryValues struct {
	matchClause         *MatchClause
	scriptScoreClause   *ScriptScoreClause
	existsClause        *ExistsClause
	booleanClause       *BooleanClause
	termClause          *TermClause
	termsClause         *TermsClause
	rangeClause         *RangeClause
	prefixClause        *PrefixClause
	fuzzyClause         *FuzzyClause
	functionScoreClause *FunctionScoreClause
	matchAllClause      *MatchAllClause
	matchNoneClause     *MatchNoneClause
}

func (q QueryValues) Match() *MatchClause {
	return q.matchClause
}
func (q *QueryValues) SetMatch(field string, matcher Matcher) error {
	return q.matchClause.Set(field, matcher)
}
func (q QueryValues) Script() *ScriptScoreClause {
	return q.scriptScoreClause
}
func (q QueryValues) Exists() *ExistsClause {
	return q.existsClause
}
func (q QueryValues) Boolean() *BooleanClause {
	return q.booleanClause
}
func (q QueryValues) Terms() *TermsClause {
	return q.termsClause
}
func (q *QueryValues) SetTerms(field string, t Termser) error {
	return q.termsClause.Set(field, t)
}
func (q QueryValues) Term() *TermClause {
	return q.termClause
}
func (q *QueryValues) SetTerm(field string, t Termer) error {
	return q.termClause.Set(field, t)
}
func (q QueryValues) IsEmpty() bool {
	return !q.matchClause.IsEmpty() || !q.termsClause.IsEmpty() || !q.termClause.IsEmpty() || !q.booleanClause.IsEmpty()
}

func (q *QueryValues) unmarshalTerm(data dynamic.JSONObject) error {
	if term, ok := data["term"]; ok {
		return q.termClause.UnmarshalJSON(term)
	}
	return nil
}

func (q *QueryValues) unmarshalTerms(data dynamic.JSONObject) error {
	if terms, ok := data["terms"]; ok {
		return q.termsClause.UnmarshalJSON(terms)
	}
	return nil
}

func (q *QueryValues) unmarshalMatch(data dynamic.JSONObject) error {
	if match, ok := data["match"]; ok {
		return q.matchClause.UnmarshalJSON(match)
	}
	return nil
}

func (q *QueryValues) unmarshalBool(data dynamic.JSONObject) error {
	if boolean, ok := data["bool"]; ok {
		return q.booleanClause.UnmarshalJSON(boolean)
	}
	return nil
}

func (q *QueryValues) UnmarshalJSON(data []byte) error {
	*q = QueryValues{}
	if len(data) == 0 || dynamic.JSON(data).IsNull() {
		return nil
	}
	m := dynamic.JSONObject{}
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	funcs := []func(dynamic.JSONObject) error{
		q.unmarshalBool,
		q.unmarshalMatch,
		q.unmarshalTerms,
		q.unmarshalTerm,
	}

	for _, fn := range funcs {
		err = fn(m)
		if err != nil {
			return err
		}
	}
	return nil
}

func (q QueryValues) marshalTerms() (string, dynamic.JSON, error) {
	terms, err := q.termsClause.MarshalJSON()
	return "terms", terms, err
}
func (q QueryValues) marshalTerm() (string, dynamic.JSON, error) {
	term, err := q.termClause.MarshalJSON()
	return "term", term, err
}

func (q QueryValues) MarshalJSON() ([]byte, error) {
	funcs := []func() (string, dynamic.JSON, error){
		q.marshalTerms,
		q.marshalTerm,
	}
	obj := dynamic.JSONObject{}
	for _, fn := range funcs {
		key, val, err := fn()
		if err != nil {
			return nil, err
		}
		if len(val) == 0 || val.IsNull() {
			continue
		}
		obj[key] = val
	}
	return json.Marshal(obj)
}

func checkField(field string, typ Kind) error {
	if len(field) == 0 {
		return NewQueryError(ErrFieldRequired, typ)
	}
	return nil
}

func checkValue(value string, typ Kind, field string) error {
	if len(value) == 0 {
		return NewQueryError(ErrValueRequired, typ, field)
	}
	return nil
}

func checkValues(values []string, typ Kind, field string) error {
	if len(values) == 0 {
		return NewQueryError(ErrValueRequired, typ, field)
	}
	return nil
}

func getField(q1 WithField, q2 WithField) string {
	var field string
	if q1 != nil {
		field = q1.Field()
	}
	if len(field) > 0 {
		return field
	}
	if q2 != nil {
		field = q2.Field()
	}
	return field

}
