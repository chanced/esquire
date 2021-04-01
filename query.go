package picker

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

type Querier interface {
	Query() (*Query, error)
}

type QueryParams struct {

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
	Terms TermserComplete

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
	Fuzzy *FuzzyQueryParams

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

	Script *ScriptQuery

	// Range returns documents that contain terms within a provided range.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-range-query.html
	Range *Range

	// MatchAll matches all documents, giving them all a _score of 1.0.
	MatchAll *MatchAll

	// MatchNone is the inverse of the match_all query, which matches no documents.
	MatchNone *MatchNoneQuery

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

func (q *QueryParams) boolean() (*BooleanClause, error) {
	if q.Boolean == nil {
		return nil, nil
	}
	return q.Boolean.Boolean()
}

func (q *QueryParams) fuzzy() (*FuzzyQuery, error) {
	if q.Fuzzy == nil {
		return nil, nil
	}
	return q.Fuzzy.Fuzzy()
}

func (q *QueryParams) term() (*TermClause, error) {
	if q.Term == nil {
		return nil, nil
	}
	return q.Term.Term()
}
func (q *QueryParams) script() (*ScriptClause, error) {
	if q.Script == nil {
		return nil, nil
	}
	return q.Script.Script()
}

func (q *QueryParams) terms() (*TermsClause, error) {
	if q.Terms == nil {
		return nil, nil
	}
	return q.Terms.Terms()
}

func (q *QueryParams) rng() (*RangeClause, error) {
	if q.Range == nil {
		return nil, nil
	}
	return q.Range.Range()
}

func (q *QueryParams) prefix() (*PrefixClause, error) {
	if q.Prefix == nil {
		return nil, nil
	}
	return q.Prefix.Prefix()
}

func (q *QueryParams) match() (*MatchClause, error) {
	if q.Match == nil {
		return nil, nil
	}
	return q.Match.Match()

}

func (q *QueryParams) scriptScore() (*ScriptScoreClause, error) {
	if q.ScriptScore == nil {
		return nil, nil
	}
	return q.ScriptScore.ScriptScore()
}

func (q *QueryParams) functionScoreClause() (*FunctionScoreClause, error) {
	if q.FunctionScore == nil {
		return nil, nil
	}
	return q.FunctionScore.FunctionScore()
}

func (q *QueryParams) matchAll() (*MatchAllClause, error) {
	if q.MatchAll == nil {
		return nil, nil
	}
	return q.MatchAll.MatchAll()
}
func (q *QueryParams) matchNone() (*MatchNoneClause, error) {
	if q.MatchNone == nil {
		return nil, nil
	}
	return q.MatchNone.MatchNone()
}

func (q *QueryParams) exists() (*ExistsClause, error) {
	if q.Exists == nil {
		return nil, nil
	}
	return q.Exists.Exists()
}

func (q *QueryParams) Query() (*Query, error) {
	if q == nil {
		return &Query{}, nil
	}
	boolean, err := q.boolean()
	if err != nil {
		return nil, err
	}

	exists, err := q.exists()
	if err != nil {
		return nil, err
	}
	term, err := q.term()
	if err != nil {
		return nil, err
	}
	terms, err := q.terms()
	if err != nil {
		return nil, err
	}
	rng, err := q.rng()
	if err != nil {
		return nil, err
	}
	prefix, err := q.prefix()
	if err != nil {
		return nil, err
	}
	match, err := q.match()
	if err != nil {
		return nil, err
	}
	matchAll, err := q.matchAll()
	if err != nil {
		return nil, err
	}
	matchNone, err := q.matchNone()
	if err != nil {
		return nil, err
	}
	scriptScore, err := q.scriptScore()
	if err != nil {
		return nil, err
	}
	script, err := q.script()
	if err != nil {
		return nil, err
	}
	fuzzy, err := q.fuzzy()
	if err != nil {
		return nil, err
	}
	funcScore, err := q.functionScoreClause()
	if err != nil {
		return nil, err
	}

	qv := &Query{
		match:         match,
		exists:        exists,
		scriptScore:   scriptScore,
		script:        script,
		fuzzy:         fuzzy,
		boolean:       boolean,
		term:          term,
		terms:         terms,
		rng:           rng,
		prefix:        prefix,
		matchAll:      matchAll,
		matchNone:     matchNone,
		functionScore: funcScore,
	}
	return qv, nil
}

// Query defines the search definition using the ElasticSearch Query DSL
//
// Elasticsearch provides a full Query DSL (Domain Specific Language) based on
// JSON to define queries. Think of the Query DSL as an AST (Abstract Syntax
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
// Query clauses behave differently depending on whether they are used in query
// context or filter context.
type Query struct {
	match         *MatchClause
	scriptScore   *ScriptScoreClause
	exists        *ExistsClause
	boolean       *BooleanClause
	term          *TermClause
	terms         *TermsClause
	rng           *RangeClause
	prefix        *PrefixClause
	fuzzy         *FuzzyQuery
	functionScore *FunctionScoreClause
	matchAll      *MatchAllClause
	matchNone     *MatchNoneClause
	script        *ScriptClause
}

func (q Query) Match() *MatchClause {
	if q.match == nil {
		q.match = &MatchClause{}
	}
	return q.match
}

func (q *Query) ScriptScore() *ScriptScoreClause {
	if q.scriptScore == nil {
		q.scriptScore = &ScriptScoreClause{}
	}
	return q.scriptScore
}

func (q *Query) Script() *ScriptClause {
	if q.script == nil {
		q.script = &ScriptClause{}
	}
	return q.script
}

func (q *Query) FunctionScore() *FunctionScoreClause {
	if q.functionScore == nil {
		q.functionScore = &FunctionScoreClause{}
	}
	return q.functionScore
}
func (q *Query) Exists() *ExistsClause {
	if q.exists == nil {
		q.exists = &ExistsClause{}
	}
	return q.exists
}
func (q Query) Boolean() *BooleanClause {
	if q.boolean == nil {
		q.boolean = &BooleanClause{}
	}
	return q.boolean
}
func (q Query) Terms() *TermsClause {
	if q.terms == nil {
		q.terms = &TermsClause{}
	}
	return q.terms
}
func (q *Query) SetTerms(field string, t Termser) error {
	if q.terms == nil {
		q.terms = &TermsClause{}
	}
	return q.terms.Set(field, t)
}
func (q Query) Term() *TermClause {
	if q.term == nil {
		q.term = &TermClause{}
	}
	return q.term
}

func (q *Query) clauses() map[QueryKind]QueryClause {

	return map[QueryKind]QueryClause{
		KindMatch:         q.match,
		KindTerm:          q.term,
		KindTerms:         q.terms,
		KindBoolean:       q.boolean,
		KindExists:        q.exists,
		KindFuzzy:         q.fuzzy,
		KindMatchAll:      q.matchAll,
		KindMatchNone:     q.matchNone,
		KindPrefix:        q.prefix,
		KindRange:         q.rng,
		KindScriptScore:   q.scriptScore,
		KindScript:        q.script,
		KindFunctionScore: q.functionScore,

		// KindBoosting: q.boosting,
		// KindConstantScore: q.constantScore
		// KindDisjunctionMax: q.disjunctionMax
	}

}

func (q *Query) IsEmpty() bool {
	if q == nil {
		return true
	}
	for _, clause := range q.clauses() {
		if !clause.IsEmpty() {
			return false
		}
	}
	return true
}

func (q *Query) unmarshalTerm(data dynamic.JSONObject) error {
	if term, ok := data["term"]; ok {
		return q.Term().UnmarshalJSON(term)
	}
	return nil
}

func (q *Query) unmarshalTerms(data dynamic.JSONObject) error {
	if terms, ok := data["terms"]; ok {
		return q.terms.UnmarshalJSON(terms)
	}
	return nil
}

func (q *Query) unmarshalMatch(data dynamic.JSONObject) error {
	if match, ok := data["match"]; ok {
		return q.match.UnmarshalJSON(match)
	}
	return nil
}

func (q *Query) unmarshalBool(data dynamic.JSONObject) error {
	if boolean, ok := data["bool"]; ok {
		q.boolean = &BooleanClause{}
		return q.boolean.UnmarshalJSON(boolean)
	}
	return nil
}

func (q *Query) UnmarshalJSON(data []byte) error {
	*q = Query{}
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

func (q Query) marshalTerms() (dynamic.JSON, error) {
	if q.terms == nil {
		return nil, nil
	}
	terms, err := q.terms.MarshalJSON()
	return terms, err
}
func (q Query) marshalTerm() (dynamic.JSON, error) {
	if q.term == nil {
		return nil, nil
	}
	term, err := q.term.MarshalJSON()
	return term, err
}

func (q Query) MarshalJSON() ([]byte, error) {
	funcs := map[string]func() (dynamic.JSON, error){
		"terms": q.marshalTerms,
		"term":  q.marshalTerm,
	}

	obj := dynamic.JSONObject{}
	for key, fn := range funcs {
		val, err := fn()
		if err != nil {
			return nil, err
		}
		if val == nil || len(val) == 0 || val.IsNull() {
			continue
		}
		obj[key] = val
	}
	return json.Marshal(obj)
}

func checkField(field string, typ QueryKind) error {
	if len(field) == 0 {
		return NewQueryError(ErrFieldRequired, typ)
	}
	return nil
}

func checkValue(value string, typ QueryKind, field string) error {
	if len(value) == 0 {
		return NewQueryError(ErrValueRequired, typ, field)
	}
	return nil
}

func checkValues(values []string, typ QueryKind, field string) error {
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
