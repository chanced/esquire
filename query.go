package picker

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

type Queryset interface {
	Queries() (Queries, error)
}

type Queriers []Querier

func (r Queriers) Queries() (Queries, error) {
	res := make(Queries, len(r))
	for k, v := range r {
		qv, err := v.Query()
		if err != nil {
			return res, err
		}
		if qv.IsEmpty() {
			continue
		}
		res[k] = qv
	}
	return res, nil
}

type Queries []*Query

func (q Queries) IsEmpty() bool {
	if len(q) == 0 {
		return true
	}
	for _, v := range q {
		if !v.IsEmpty() {
			return false
		}
	}
	return true
}
func (q Queries) Queries() (Queries, error) {
	res := make(Queries, 0, len(q))
	for i, v := range q {
		if !v.IsEmpty() {
			res[i] = v
		}
	}
	return res, nil
}

func (q *Queries) Add(params Querier) (*Query, error) {
	qv, err := params.Query()
	if err != nil {
		return qv, err
	}
	if q == nil {
		*q = Queries{qv}
		return qv, nil
	}
	*q = append(*q, qv)
	return qv, nil
}

type Querier interface {
	Query() (*Query, error)
}

// TODO: Add specific clause functions so the actual query, like *TermQuery, can be used as a param

type QueryParams struct {

	// Term returns documents that contain an exact term in a provided field.
	//
	// You can use the term query to find documents based on a precise value
	// such as a price, a product ID, or a username.
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
	Term CompleteTermer

	// Terms returns documents that contain one or more exact terms in a provided
	// field.
	//
	// The terms query is the same as the term query, except you can search for
	// multiple values.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-terms-query.html
	Terms CompleteTermser

	// Match returns documents that match a provided text, number, date or
	// boolean value. The provided text is analyzed before matching.
	//
	// The match query is the standard query for performing a full-text search,
	// including options for fuzzy matching.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-query.html
	Match CompleteMatcher

	// Bool is a query that matches documents matching boolean combinations
	// of other queries. The bool query maps to Lucene BooleanQuery. It is built
	// using one or more boolean clauses, each clause with a typed occurrence.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-bool-query.html
	Bool Booler

	// Fuzzy returns documents that contain terms similar to the search term, as
	// measured by a Levenshtein edit distance.
	//
	// An edit distance is the number of one-character changes needed to turn
	// one term into another. These changes can include:
	//
	// - Changing a character (box → fox)
	//
	// - Removing a character (black → lack)
	//
	// - Inserting a character (sic → sick)
	//
	// - Transposing two adjacent characters (act → cat)
	//
	// To find similar terms, the fuzzy query creates a set of all possible
	// variations, or expansions, of the search term within a specified edit
	// distance. The query then returns exact matches for each expansion.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-fuzzy-query.html
	Fuzzy CompleteFuzzier

	// Prefix returns documents that contain a specific prefix in a provided
	// field.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-prefix-query.html
	Prefix CompletePrefixer

	// FunctionScore  allows you to modify the score of documents that are
	// retrieved by a query. This can be useful if, for example, a score
	// function is computationally expensive and it is sufficient to compute the
	// score on a filtered set of documents.
	//
	// To use function_score, the user has to define a query and one or more
	// functions, that compute a new score for each document returned by the
	// query.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-function-score-query.html
	FunctionScore FunctionScorer

	// ScoreScript uses a script to provide a custom score for returned
	// documents.
	//
	// The script_score query is useful if, for example, a scoring function is
	// expensive and you only need to calculate the score of a filtered set of
	// documents.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-script-score-query.html
	ScriptScore ScriptScorer

	// Filters documents based on a provided script. The script query is
	// typically used in a filter context.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-script-query.html
	Script Scripter

	// Range returns documents that contain terms within a provided range.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-range-query.html
	Range Ranger

	// MatchAll matches all documents, giving them all a _score of 1.0.
	MatchAll *MatchAllQueryParams

	// MatchNone is the inverse of the match_all query, which matches no documents.
	MatchNone *MatchNoneQueryParams

	// Exists returns documents that contain an indexed value for a field.
	//
	// An indexed value may not exist for a document’s field due to a variety of
	// reasons:
	//
	// - The field in the source JSON is null or []
	//
	// - The field has "index" : false set in the mapping
	//
	// - The length of the field value exceeded an ignore_above setting in the
	// mapping
	//
	// - The field value was malformed and ignore_malformed was defined in the
	// mapping
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-exists-query.html
	Exists Exister
	// Returns documents matching a positive query while reducing the relevance
	// score of documents that also match a negative query.
	//
	// You can use the boosting query to demote certain documents without
	// excluding them from the search results.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-boosting-query.html
	Boosting Boostinger

	// A query which wraps another query, but executes it in filter context. All
	// matching documents are given the same “constant” _score.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-constant-score-query.html
	ConstantScore ConstantScorer
	// Returns documents matching one or more wrapped queries, called query
	// clauses or clauses.
	//
	// If a returned document matches multiple query clauses, the dis_max query
	// assigns the document the highest relevance score from any matching
	// clause, plus a tie breaking increment for any additional matching
	// subqueries.
	//
	// You can use the dis_max to search for a term in fields mapped with
	// different boost factors.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-dis-max-query.html
	DisjunctionMax DisjunctionMaxer
	// Returns documents based on their IDs. This query uses document IDs stored
	// in the _id field.
	IDs IDser
	// Returns documents based on the order and proximity of matching terms.
	//
	// The intervals query uses matching rules, constructed from a small set of
	// definitions. These rules are then applied to terms from a specified
	// field.
	//
	// The definitions produce sequences of minimal intervals that span terms in
	// a body of text. These intervals can be further combined and filtered by
	// parent sources.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-intervals-query.html#intervals-all_of
	Intervals Intervalser
	// A match_bool_prefix query analyzes its input and constructs a bool query
	// from the terms. Each term except the last is used in a term query. The
	// last term is used in a prefix.
	MatchBoolPrefix MatchBoolPrefixer
	// The match_phrase query analyzes the text and creates a phrase query out
	// of the analyzed text.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-match-query-phrase.html
	MatchPhrase MatchPhraser
	// The multi-field version of the match query.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-multi-match-query.html#multi-match-types
	MultiMatch MultiMatcher
	// Returns documents based on a provided query string, using a parser with a strict syntax.
	//
	// This query uses a syntax to parse and split the provided query string based
	// on operators, such as AND or NOT. The query then analyzes each split text
	// independently before returning matching documents.
	//
	// You can use the query_string query to create a complex search that includes
	// wildcard characters, searches across multiple fields, and more. While
	// versatile, the query is strict and returns an error if the query string
	// includes any invalid syntax.
	QueryString QueryStringer
	// Returns documents based on a provided query string, using a parser with a
	// limited but fault-tolerant syntax.
	//
	// This query uses a simple syntax to parse and split the provided query
	// string into terms based on special operators. The query then analyzes
	// each term independently before returning matching documents.
	//
	// While its syntax is more limited than the query_string query, the
	// simple_query_string query does not return errors for invalid syntax.
	// Instead, it ignores any invalid parts of the query string.
	SimpleQueryString SimpleQueryStringer

	GeoBoundingBox GeoBoundingBoxer
}

func (q QueryParams) geoBoundingBox() (*GeoBoundingBoxQuery, error) {
	if q.GeoBoundingBox == nil {
		return nil, nil
	}
	return q.GeoBoundingBox.GeoBoundingBox()
}
func (q QueryParams) simpleQueryString() (*SimpleQueryStringQuery, error) {
	if q.SimpleQueryString == nil {
		return nil, nil
	}
	return q.SimpleQueryString.SimpleQueryString()
}
func (q QueryParams) queryString() (*QueryStringQuery, error) {
	if q.QueryString == nil {
		return nil, nil
	}
	return q.QueryString.QueryString()
}

func (q QueryParams) matchPhrase() (*MatchPhraseQuery, error) {
	if q.MatchPhrase == nil {
		return nil, nil
	}
	return q.MatchPhrase.MatchPhrase()
}
func (q QueryParams) boolean() (*BoolQuery, error) {
	if q.Bool == nil {
		return nil, nil
	}
	return q.Bool.Bool()
}
func (q QueryParams) matchBoolPrefix() (*MatchBoolPrefixQuery, error) {
	if q.MatchBoolPrefix == nil {
		return nil, nil
	}
	return q.MatchBoolPrefix.MatchBoolPrefix()
}
func (q QueryParams) fuzzy() (*FuzzyQuery, error) {
	if q.Fuzzy == nil {
		return nil, nil
	}
	return q.Fuzzy.Fuzzy()
}
func (q QueryParams) ids() (*IDsQuery, error) {
	if q.IDs == nil {
		return nil, nil
	}
	return q.IDs.IDs()
}
func (q QueryParams) term() (*TermQuery, error) {
	if q.Term == nil {
		return nil, nil
	}
	return q.Term.Term()
}
func (q QueryParams) script() (*ScriptQuery, error) {
	if q.Script == nil {
		return nil, nil
	}
	return q.Script.Script()
}

func (q QueryParams) terms() (*TermsQuery, error) {
	if q.Terms == nil {
		return nil, nil
	}
	return q.Terms.Terms()
}

func (q QueryParams) rng() (*RangeQuery, error) {
	if q.Range == nil {
		return nil, nil
	}
	return q.Range.Range()
}

func (q QueryParams) prefix() (*PrefixQuery, error) {
	if q.Prefix == nil {
		return nil, nil
	}
	return q.Prefix.Prefix()
}

func (q QueryParams) match() (*MatchQuery, error) {
	if q.Match == nil {
		return nil, nil
	}
	return q.Match.Match()

}

func (q QueryParams) scriptScore() (*ScriptScoreQuery, error) {
	if q.ScriptScore == nil {
		return nil, nil
	}
	return q.ScriptScore.ScriptScore()
}

func (q QueryParams) functionScoreClause() (*FunctionScoreQuery, error) {
	if q.FunctionScore == nil {
		return nil, nil
	}
	return q.FunctionScore.FunctionScore()
}

func (q QueryParams) matchAll() (*MatchAllQuery, error) {
	if q.MatchAll == nil {
		return nil, nil
	}
	return q.MatchAll.MatchAll()
}
func (q QueryParams) matchNone() (*MatchNoneQuery, error) {
	if q.MatchNone == nil {
		return nil, nil
	}
	return q.MatchNone.MatchNone()
}

func (q QueryParams) exists() (*ExistsQuery, error) {
	if q.Exists == nil {
		return nil, nil
	}
	return q.Exists.Exists()
}

func (q QueryParams) boosting() (*BoostingQuery, error) {
	if q.Boosting == nil {
		return nil, nil
	}
	return q.Boosting.Boosting()
}

func (q QueryParams) constantScore() (*ConstantScoreQuery, error) {
	if q.ConstantScore == nil {
		return nil, nil
	}
	return q.ConstantScore.ConstantScore()
}
func (q QueryParams) disjunectionMax() (*DisjunctionMaxQuery, error) {
	if q.DisjunctionMax == nil {
		return nil, nil
	}
	return q.DisjunctionMax.DisjunctionMax()
}
func (q QueryParams) intervals() (*IntervalsQuery, error) {
	if q.Intervals == nil {
		return nil, nil
	}
	return q.Intervals.Intervals()
}
func (q QueryParams) multiMatch() (*MultiMatchQuery, error) {
	if q.MultiMatch == nil {
		return nil, nil
	}
	return q.MultiMatch.MultiMatch()
}
func (q *QueryParams) Query() (*Query, error) {
	if q == nil {
		return &Query{}, nil
	}

	geoBoundingBox, err := q.geoBoundingBox()

	simpleQueryString, err := q.simpleQueryString()
	if err != nil {
		return nil, err
	}
	queryString, err := q.queryString()
	if err != nil {
		return nil, err
	}
	boolean, err := q.boolean()
	if err != nil {
		return nil, err
	}
	matchBoolPrefix, err := q.matchBoolPrefix()
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
	boosting, err := q.boosting()
	if err != nil {
		return nil, err
	}
	constantScore, err := q.constantScore()
	if err != nil {
		return nil, err
	}
	disjunctionMax, err := q.disjunectionMax()
	if err != nil {
		return nil, err
	}
	ids, err := q.ids()
	if err != nil {
		return nil, err
	}
	intervals, err := q.intervals()
	if err != nil {
		return nil, err
	}
	matchPhrase, err := q.matchPhrase()
	if err != nil {
		return nil, err
	}
	multiMatch, err := q.multiMatch()
	if err != nil {
		return nil, err
	}
	qv := &Query{
		match:             match,
		exists:            exists,
		scriptScore:       scriptScore,
		script:            script,
		fuzzy:             fuzzy,
		boolean:           boolean,
		term:              term,
		terms:             terms,
		rng:               rng,
		prefix:            prefix,
		matchAll:          matchAll,
		matchNone:         matchNone,
		functionScore:     funcScore,
		boosting:          boosting,
		constantScore:     constantScore,
		disjunctionMax:    disjunctionMax,
		ids:               ids,
		intervals:         intervals,
		matchBoolPrefix:   matchBoolPrefix,
		matchPhrase:       matchPhrase,
		multiMatch:        multiMatch,
		queryString:       queryString,
		simpleQueryString: simpleQueryString,
		geoBoundingBox:    geoBoundingBox,
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
	match             *MatchQuery
	scriptScore       *ScriptScoreQuery
	exists            *ExistsQuery
	boolean           *BoolQuery
	term              *TermQuery
	terms             *TermsQuery
	rng               *RangeQuery
	prefix            *PrefixQuery
	fuzzy             *FuzzyQuery
	functionScore     *FunctionScoreQuery
	matchAll          *MatchAllQuery
	matchNone         *MatchNoneQuery
	script            *ScriptQuery
	boosting          *BoostingQuery
	constantScore     *ConstantScoreQuery
	disjunctionMax    *DisjunctionMaxQuery
	ids               *IDsQuery
	intervals         *IntervalsQuery
	matchBoolPrefix   *MatchBoolPrefixQuery
	matchPhrase       *MatchPhraseQuery
	multiMatch        *MultiMatchQuery
	queryString       *QueryStringQuery
	simpleQueryString *SimpleQueryStringQuery
	geoBoundingBox    *GeoBoundingBoxQuery
}

func (q *Query) Query() (*Query, error) {
	return q, nil
}

func (q *Query) GeoBoundingBox() *GeoBoundingBoxQuery {
	if q.geoBoundingBox == nil {
		q.geoBoundingBox = &GeoBoundingBoxQuery{}
	}
	return q.geoBoundingBox
}

func (q *Query) QueryString() *QueryStringQuery {
	if q.queryString == nil {
		q.queryString = &QueryStringQuery{}
	}
	return q.queryString
}
func (q *Query) SimpleQueryString() *SimpleQueryStringQuery {
	if q.simpleQueryString == nil {
		q.simpleQueryString = &SimpleQueryStringQuery{}
	}
	return q.simpleQueryString
}

func (q *Query) MultiMatch() *MultiMatchQuery {
	if q.multiMatch == nil {
		q.multiMatch = &MultiMatchQuery{}
	}
	return q.multiMatch
}
func (q *Query) MatchBoolPrefix() *MatchBoolPrefixQuery {
	if q.matchBoolPrefix == nil {
		q.matchBoolPrefix = &MatchBoolPrefixQuery{}
	}
	return q.matchBoolPrefix
}
func (q *Query) MatchPhrase() *MatchPhraseQuery {
	if q.matchPhrase == nil {
		q.matchPhrase = &MatchPhraseQuery{}
	}
	return q.matchPhrase
}
func (q *Query) Range() *RangeQuery {
	if q.rng == nil {
		q.rng = &RangeQuery{}
	}
	return q.rng
}
func (q *Query) Prefix() *PrefixQuery {
	if q.prefix == nil {
		q.prefix = &PrefixQuery{}
	}
	return q.prefix
}
func (q *Query) Fuzzy() *FuzzyQuery {
	if q.fuzzy == nil {
		q.fuzzy = &FuzzyQuery{}
	}
	return q.fuzzy
}
func (q *Query) Intervals() *IntervalsQuery {
	if q.intervals == nil {
		q.intervals = &IntervalsQuery{}
	}
	return q.intervals
}
func (q *Query) DisjunctionMax() *DisjunctionMaxQuery {
	if q.boosting == nil {
		q.disjunctionMax = &DisjunctionMaxQuery{}
	}
	return q.disjunctionMax
}
func (q *Query) IDs() *IDsQuery {
	if q.boosting == nil {
		q.ids = &IDsQuery{}
	}
	return q.ids
}
func (q *Query) ConstantScore() *ConstantScoreQuery {
	if q.boosting == nil {
		q.constantScore = &ConstantScoreQuery{}
	}
	return q.constantScore
}
func (q *Query) Boosting() *BoostingQuery {
	if q.boosting == nil {
		q.boosting = &BoostingQuery{}
	}
	return q.boosting
}
func (q *Query) Match() *MatchQuery {
	if q.match == nil {
		q.match = &MatchQuery{}
	}
	return q.match
}

func (q *Query) ScriptScore() *ScriptScoreQuery {
	if q.scriptScore == nil {
		q.scriptScore = &ScriptScoreQuery{}
	}
	return q.scriptScore
}

func (q *Query) Script() *ScriptQuery {
	if q.script == nil {
		q.script = &ScriptQuery{}
	}
	return q.script
}

func (q *Query) FunctionScore() *FunctionScoreQuery {
	if q.functionScore == nil {
		q.functionScore = &FunctionScoreQuery{}
	}
	return q.functionScore
}
func (q *Query) Exists() *ExistsQuery {
	if q.exists == nil {
		q.exists = &ExistsQuery{}
	}
	return q.exists
}
func (q *Query) Bool() *BoolQuery {
	if q.boolean == nil {
		q.boolean = &BoolQuery{}
	}
	return q.boolean
}
func (q *Query) Terms() *TermsQuery {
	if q.terms == nil {
		q.terms = &TermsQuery{}
	}
	return q.terms
}

// func (q *Query) SetTerms(field string, t Termser) error {
// 	if q.terms == nil {
// 		q.terms = &TermsQuery{}
// 	}
// 	return q.terms.Set(field, t)
// }

func (q *Query) Term() *TermQuery {
	if q.term == nil {
		q.term = &TermQuery{}
	}
	return q.term
}

func (q *Query) clauses() map[QueryKind]QueryClause {

	return map[QueryKind]QueryClause{
		QueryKindMatch:             q.match,
		QueryKindTerm:              q.term,
		QueryKindTerms:             q.terms,
		QueryKindBoolean:           q.boolean,
		QueryKindExists:            q.exists,
		QueryKindFuzzy:             q.fuzzy,
		QueryKindMatchAll:          q.matchAll,
		QueryKindMatchNone:         q.matchNone,
		QueryKindPrefix:            q.prefix,
		QueryKindRange:             q.rng,
		QueryKindScriptScore:       q.scriptScore,
		QueryKindScript:            q.script,
		QueryKindFunctionScore:     q.functionScore,
		QueryKindBoosting:          q.boosting,
		QueryKindConstantScore:     q.constantScore,
		QueryKindDisjunctionMax:    q.disjunctionMax,
		QueryKindIDs:               q.ids,
		QueryKindIntervals:         q.intervals,
		QueryKindMatchBoolPrefix:   q.matchBoolPrefix,
		QueryKindMatchPhrase:       q.matchPhrase,
		QueryKindMultiMatch:        q.multiMatch,
		QueryKindQueryString:       q.queryString,
		QueryKindSimpleQueryString: q.simpleQueryString,
		QueryKindGeoBoundingBox:    q.geoBoundingBox,
	}
}

func (q *Query) setClause(qc QueryClause) {
	switch qc.Kind() {
	case QueryKindSimpleQueryString:
		q.simpleQueryString = qc.(*SimpleQueryStringQuery)
	case QueryKindQueryString:
		q.queryString = qc.(*QueryStringQuery)
	case QueryKindMatch:
		q.match = qc.(*MatchQuery)
	case QueryKindTerm:
		q.term = qc.(*TermQuery)
	case QueryKindTerms:
		q.terms = qc.(*TermsQuery)
	case QueryKindBoolean:
		q.boolean = qc.(*BoolQuery)
	case QueryKindExists:
		q.exists = qc.(*ExistsQuery)
	case QueryKindFuzzy:
		q.fuzzy = qc.(*FuzzyQuery)
	case QueryKindMatchAll:
		q.matchAll = qc.(*MatchAllQuery)
	case QueryKindMatchNone:
		q.matchNone = qc.(*MatchNoneQuery)
	case QueryKindPrefix:
		q.prefix = qc.(*PrefixQuery)
	case QueryKindRange:
		q.rng = qc.(*RangeQuery)
	case QueryKindScriptScore:
		q.scriptScore = qc.(*ScriptScoreQuery)
	case QueryKindScript:
		q.script = qc.(*ScriptQuery)
	case QueryKindFunctionScore:
		q.functionScore = qc.(*FunctionScoreQuery)
	case QueryKindBoosting:
		q.boosting = qc.(*BoostingQuery)
	case QueryKindConstantScore:
		q.constantScore = qc.(*ConstantScoreQuery)
	case QueryKindDisjunctionMax:
		q.disjunctionMax = qc.(*DisjunctionMaxQuery)
	case QueryKindIDs:
		q.ids = qc.(*IDsQuery)
	case QueryKindIntervals:
		q.intervals = qc.(*IntervalsQuery)
	case QueryKindMatchPhrase:
		q.matchPhrase = qc.(*MatchPhraseQuery)
	case QueryKindMultiMatch:
		q.multiMatch = qc.(*MultiMatchQuery)
	case QueryKindGeoBoundingBox:
		q.geoBoundingBox = qc.(*GeoBoundingBoxQuery)
	}
}
func (q *Query) Set(params Querier) error {
	qv, err := params.Query()
	if err != nil {
		return err
	}
	*q = *qv
	return nil
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

func (q *Query) UnmarshalJSON(data []byte) error {
	*q = Query{}
	if len(data) == 0 || dynamic.JSON(data).IsNull() {
		return nil
	}
	obj := dynamic.JSONObject{}
	err := json.Unmarshal(data, &obj)
	if err != nil {
		return err
	}
	for k, v := range obj {
		handler, ok := queryKindHandlers[QueryKind(k)]
		if !ok {
			continue
		}
		c := handler()
		err := c.UnmarshalJSON(v)
		if err != nil {
			return err
		}
		q.setClause(c)
	}
	return nil
}

func (q Query) MarshalJSON() ([]byte, error) {

	obj := dynamic.JSONObject{}
	for key, clause := range q.clauses() {
		if clause.IsEmpty() {
			continue
		}
		val, err := clause.MarshalJSON()
		if err != nil {
			return nil, err
		}
		if len(val) == 0 || dynamic.JSON(val).IsNull() {
			continue
		}
		obj[key.String()] = val
	}
	return json.Marshal(obj)
}

func checkField(field string, typ QueryKind) error {
	if len(field) == 0 {
		return newQueryError(ErrFieldRequired, typ)
	}
	return nil
}

func checkValue(value string, typ QueryKind, field string) error {
	if len(value) == 0 {
		return newQueryError(ErrValueRequired, typ, field)
	}
	return nil
}

func checkValues(values []string, typ QueryKind, field string) error {
	if len(values) == 0 {
		return newQueryError(ErrValueRequired, typ, field)
	}
	return nil
}
