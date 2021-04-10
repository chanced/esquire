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
	MatchPhrasePrefix MatchPhrasePrefixer
	GeoBoundingBox    GeoBoundingBoxer
	GeoDistance       GeoDistancer
	GeoShape          GeoShaper
	Shape             ShapeQuerier
	Nested            NestedQuerier
	// HasChild         HasChilder
	// HasParent        HasParenter
	ParentID ParentIDer
	// DistanceFeature   DistanceFeaturer
	MoreLikeThis MoreLikeThiser
	Percolate    Percolater
	// RankFeature      RankFeaturer
	// Wrapper          Wrapperer
	// Pinned           Pinneder
	// GeoPolygon       GeoPolygoner
	// SpanContaining   SpanContaininger
	// FieldMaskingSpan FieldMaskingSpanner
	// SpanFirst        SpanFirster
	// SpanMulti        SpanMultier
	// SpanNear         SpanNearer
	// SpanNot          SpanNoter
	// SpanOr           SpanOrer
	// SpanTerm         SpanTermer
	// SpanWithin       SpanWithiner
	// Common   Commoner
	// Regexp   Regexper
	// TermSet  TermSetter
	// Type     Typer
	Wildcard Wildcarder
}

// func (q QueryParams) common() (*CommonQuery, error) {
// 	if q.Common == nil {
// 		return nil, nil
// 	}
// 	return q.Common.Common()
// }
// func (q QueryParams) regexp() (*RegexpQuery, error) {
// 	if q.Regexp == nil {
// 		return nil, nil
// 	}
// 	return q.Regexp.Regexp()
// }
// func (q QueryParams) termSet() (*TermSetQuery, error) {
// 	if q.TermSet == nil {
// 		return nil, nil
// 	}
// 	return q.TermSet.TermSet()
// }
// func (q QueryParams) typ() (*TypeQuery, error) {
// 	if q.Type == nil {
// 		return nil, nil
// 	}
// 	return q.Type.Type()
// }
func (q QueryParams) wildcard() (*WildcardQuery, error) {
	if q.Wildcard == nil {
		return nil, nil
	}
	return q.Wildcard.Wildcard()
}
func (q QueryParams) matchPhrasePrefix() (*MatchPhrasePrefixQuery, error) {
	if q.MatchPhrasePrefix == nil {
		return nil, nil
	}
	return q.MatchPhrasePrefix.MatchPhrasePrefix()
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

func (q QueryParams) geoDistance() (*GeoDistanceQuery, error) {
	if q.GeoDistance == nil {
		return nil, nil
	}
	return q.GeoDistance.GeoDistance()
}

// func (q QueryParams) geoPolygon() (*GeoPolygonQuery, error) {
// 	if q.GeoPolygon == nil {
// 		return nil, nil
// 	}
// 	return q.GeoPolygon.GeoPolygon()
// }
func (q QueryParams) geoShape() (*GeoShapeQuery, error) {
	if q.GeoShape == nil {
		return nil, nil
	}
	return q.GeoShape.GeoShape()
}

func (q QueryParams) shape() (*ShapeQuery, error) {
	if q.Shape == nil {
		return nil, nil
	}
	return q.Shape.ShapeQuery()
}

func (q QueryParams) nested() (*NestedQuery, error) {
	if q.Nested == nil {
		return nil, nil
	}
	return q.Nested.Nested()
}

// func (q QueryParams) hasChild() (*HasChildQuery, error) {
// 	if q.HasChild == nil {
// 		return nil, nil
// 	}
// 	return q.HasChild.HasChild()
// }
// func (q QueryParams) hasParent() (*HasParentQuery, error) {
// 	if q.HasParent == nil {
// 		return nil, nil
// 	}
// 	return q.HasParent.HasParent()
// }
func (q QueryParams) parentID() (*ParentIDQuery, error) {
	if q.ParentID == nil {
		return nil, nil
	}
	return q.ParentID.ParentID()
}

// func (q QueryParams) distanceFeature() (*DistanceFeatureQuery, error) {
// 	if q.DistanceFeature == nil {
// 		return nil, nil
// 	}
// 	return q.DistanceFeature.DistanceFeature()
// }
func (q QueryParams) moreLikeThis() (*MoreLikeThisQuery, error) {
	if q.MoreLikeThis == nil {
		return nil, nil
	}
	return q.MoreLikeThis.MoreLikeThis()
}

func (q QueryParams) percolate() (*PercolateQuery, error) {
	if q.Percolate == nil {
		return nil, nil
	}
	return q.Percolate.Percolate()
}

// func (q QueryParams) rankFeature() (*RankFeatureQuery, error) {
// 	if q.RankFeature == nil {
// 		return nil, nil
// 	}
// 	return q.RankFeature.RankFeature()
// }
// func (q QueryParams) wrapper() (*WrapperQuery, error) {
// 	if q.Wrapper == nil {
// 		return nil, nil
// 	}
// 	return q.Wrapper.Wrapper()
// }
// func (q QueryParams) pinned() (*PinnedQuery, error) {
// 	if q.Pinned == nil {
// 		return nil, nil
// 	}
// 	return q.Pinned.Pinned()
// }
// func (q QueryParams) spanContaining() (*SpanContainingQuery, error) {
// 	if q.SpanContaining == nil {
// 		return nil, nil
// 	}
// 	return q.SpanContaining.SpanContaining()
// }
// func (q QueryParams) fieldMaskingSpan() (*FieldMaskingSpanQuery, error) {
// 	if q.FieldMaskingSpan == nil {
// 		return nil, nil
// 	}
// 	return q.FieldMaskingSpan.FieldMaskingSpan()
// }
// func (q QueryParams) spanFirst() (*SpanFirstQuery, error) {
// 	if q.SpanFirst == nil {
// 		return nil, nil
// 	}
// 	return q.SpanFirst.SpanFirst()
// }
// func (q QueryParams) spanMulti() (*SpanMultiQuery, error) {
// 	if q.SpanMulti == nil {
// 		return nil, nil
// 	}
// 	return q.SpanMulti.SpanMulti()
// }
// func (q QueryParams) spanNear() (*SpanNearQuery, error) {
// 	if q.SpanNear == nil {
// 		return nil, nil
// 	}
// 	return q.SpanNear.SpanNear()
// }
// func (q QueryParams) spanNot() (*SpanNotQuery, error) {
// 	if q.SpanNot == nil {
// 		return nil, nil
// 	}
// 	return q.SpanNot.SpanNot()
// }
// func (q QueryParams) spanOr() (*SpanOrQuery, error) {
// 	if q.SpanOr == nil {
// 		return nil, nil
// 	}
// 	return q.SpanOr.SpanOr()
// }
// func (q QueryParams) spanTerm() (*SpanTermQuery, error) {
// 	if q.SpanTerm == nil {
// 		return nil, nil
// 	}
// 	return q.SpanTerm.SpanTerm()
// }
// func (q QueryParams) spanWithin() (*SpanWithinQuery, error) {
// 	if q.SpanWithin == nil {
// 		return nil, nil
// 	}
// 	return q.SpanWithin.SpanWithin()
// }

func (q *QueryParams) Query() (*Query, error) {
	if q == nil {
		return &Query{}, nil
	}
	matchPhrasePrefix, err := q.matchPhrasePrefix()
	if err != nil {
		return nil, err
	}
	geoBoundingBox, err := q.geoBoundingBox()
	if err != nil {
		return nil, err
	}
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
	// common, err := q.common()
	// if err != nil {
	// 	return nil, err
	// }
	// regexp, err := q.regexp()
	// if err != nil {
	// 	return nil, err
	// }
	// termSet, err := q.termSet()
	// if err != nil {
	// 	return nil, err
	// }
	// typ, err := q.typ()
	// if err != nil {
	// 	return nil, err
	// }

	wildcard, err := q.wildcard()
	if err != nil {
		return nil, err
	}
	geoDistance, err := q.geoDistance()
	if err != nil {
		return nil, err
	}
	// geoPolygon, err := q.geoPolygon()
	// if err != nil {
	// 	return nil, err
	// }
	geoShape, err := q.geoShape()
	if err != nil {
		return nil, err
	}
	shape, err := q.shape()
	if err != nil {
		return nil, err
	}
	nested, err := q.nested()
	if err != nil {
		return nil, err
	}
	// hasChild, err := q.hasChild()
	// if err != nil {
	// 	return nil, err
	// }
	// hasParent, err := q.hasParent()
	// if err != nil {
	// 	return nil, err
	// }
	parentID, err := q.parentID()
	if err != nil {
		return nil, err
	}
	// distanceFeature, err := q.distanceFeature()
	// if err != nil {
	// 	return nil, err
	// }
	moreLikeThis, err := q.moreLikeThis()
	if err != nil {
		return nil, err
	}
	percolate, err := q.percolate()
	if err != nil {
		return nil, err
	}
	// rankFeature, err := q.rankFeature()
	// if err != nil {
	// 	return nil, err
	// }
	// wrapper, err := q.wrapper()
	// if err != nil {
	// 	return nil, err
	// }
	// pinned, err := q.pinned()
	// if err != nil {
	// 	return nil, err
	// }
	// spanContaining, err := q.spanContaining()
	// if err != nil {
	// 	return nil, err
	// }
	// fieldMaskingSpan, err := q.fieldMaskingSpan()
	// if err != nil {
	// 	return nil, err
	// }
	// spanFirst, err := q.spanFirst()
	// if err != nil {
	// 	return nil, err
	// }
	// spanMulti, err := q.spanMulti()
	// if err != nil {
	// 	return nil, err
	// }
	// spanNear, err := q.spanNear()
	// if err != nil {
	// 	return nil, err
	// }
	// spanNot, err := q.spanNot()
	// if err != nil {
	// 	return nil, err
	// }
	// spanOr, err := q.spanOr()
	// if err != nil {
	// 	return nil, err
	// }
	// spanTerm, err := q.spanTerm()
	// if err != nil {
	// 	return nil, err
	// }
	// spanWithin, err := q.spanWithin()
	// if err != nil {
	// 	return nil, err
	// }
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
		matchPhrasePrefix: matchPhrasePrefix,
		wildcard:          wildcard,
		// common:            common,
		// regexp:            regexp,
		// termSet:           termSet,
		// typ:               typ,
		geoDistance: geoDistance,
		// geoPolygon:        geoPolygon,
		geoShape: geoShape,
		shape:    shape,
		nested:   nested,
		// hasChild:          hasChild,
		// hasParent:         hasParent,
		parentID: parentID,
		// distanceFeature:    distanceFeature,
		moreLikeThis: moreLikeThis,
		percolate:    percolate,
		// rankFeature:       rankFeature,
		// wrapper:           wrapper,
		// pinned:            pinned,
		// spanContaining:    spanContaining,
		// fieldMaskingSpan:  fieldMaskingSpan,
		// spanFirst:         spanFirst,
		// spanMulti:         spanMulti,
		// spanNear:          spanNear,
		// spanNot:           spanNot,
		// spanOr:            spanOr,
		// spanTerm:          spanTerm,
		// spanWithin:        spanWithin,
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
	matchPhrasePrefix *MatchPhrasePrefixQuery
	multiMatch        *MultiMatchQuery
	queryString       *QueryStringQuery
	simpleQueryString *SimpleQueryStringQuery
	geoBoundingBox    *GeoBoundingBoxQuery
	wildcard          *WildcardQuery
	// common            *CommonQuery
	// regexp            *RegexpQuery
	// termSet           *TermSetQuery
	// typ               *TypeQuery
	geoDistance *GeoDistanceQuery
	// geoPolygon        *GeoPolygonQuery
	geoShape *GeoShapeQuery
	shape    *ShapeQuery
	nested   *NestedQuery
	// hasChild          *HasChildQuery
	// hasParent         *HasParentQuery
	parentID *ParentIDQuery
	// distanceFeature    *DistanceFeatureQuery
	moreLikeThis *MoreLikeThisQuery
	percolate    *PercolateQuery
	// rankFeature       *RankFeatureQuery
	// wrapper           *WrapperQuery
	// pinned            *PinnedQuery
	// spanContaining    *SpanContainingQuery
	// fieldMaskingSpan  *FieldMaskingSpanQuery
	// spanFirst         *SpanFirstQuery
	// spanMulti         *SpanMultiQuery
	// spanNear          *SpanNearQuery
	// spanNot           *SpanNotQuery
	// spanOr            *SpanOrQuery
	// spanTerm          *SpanTermQuery
	// spanWithin        *SpanWithinQuery
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

func (q *Query) MultiMatch() *MultiMatchQuery {
	if q.multiMatch == nil {
		q.multiMatch = &MultiMatchQuery{}
	}
	return q.multiMatch
}
func (q *Query) Prefix() *PrefixQuery {
	if q.prefix == nil {
		q.prefix = &PrefixQuery{}
	}
	return q.prefix
}
func (q *Query) Match() *MatchQuery {
	if q.match == nil {
		q.match = &MatchQuery{}
	}
	return q.match
}
func (q *Query) MatchAll() *MatchAllQuery {
	if q.matchAll == nil {
		q.matchAll = &MatchAllQuery{}
	}
	return q.matchAll
}
func (q *Query) MatchNone() *MatchNoneQuery {
	if q.matchNone == nil {
		q.matchNone = &MatchNoneQuery{}
	}
	return q.matchNone
}

// func (q *Query) Common() *CommonQuery {
// 	if q.common == nil {
// 		q.common = &CommonQuery{}
// 	}
// 	return q.common
// }
// func (q *Query) Regexp() *RegexpQuery {
// 	if q.regexp == nil {
// 		q.regexp = &RegexpQuery{}
// 	}
// 	return q.regexp
// }
// func (q *Query) TermSet() *TermSetQuery {
// 	if q.termSet == nil {
// 		q.termSet = &TermSetQuery{}
// 	}
// 	return q.termSet
// }
// func (q *Query) Type() *TypeQuery {
// 	if q.typ == nil {
// 		q.typ = &TypeQuery{}
// 	}
// 	return q.typ
// }

func (q *Query) Wildcard() *WildcardQuery {
	if q.wildcard == nil {
		q.wildcard = &WildcardQuery{}
	}
	return q.wildcard
}

// func (q *Query) AllOf() *AllOfQuery {
//     if q.allOf == nil {
//         q.allOf = &AllOfQuery{}
//     }
//     return q.allOf
// }

func (q *Query) MatchPhrasePrefix() *MatchPhrasePrefixQuery {
	if q.matchPhrasePrefix == nil {
		q.matchPhrasePrefix = &MatchPhrasePrefixQuery{}
	}
	return q.matchPhrasePrefix
}

func (q *Query) GeoDistance() *GeoDistanceQuery {
	if q.geoDistance == nil {
		q.geoDistance = &GeoDistanceQuery{}
	}
	return q.geoDistance
}

// func (q *Query) GeoPolygon() *GeoPolygonQuery {
//     if q.geoPolygon == nil {
//         q.geoPolygon = &GeoPolygonQuery{}
//     }
//     return q.geoPolygon
// }

func (q *Query) GeoShape() *GeoShapeQuery {
	if q.geoShape == nil {
		q.geoShape = &GeoShapeQuery{}
	}
	return q.geoShape
}

func (q *Query) Shape() *ShapeQuery {
	if q.shape == nil {
		q.shape = &ShapeQuery{}
	}
	return q.shape
}

func (q *Query) Nested() *NestedQuery {
	if q.nested == nil {
		q.nested = &NestedQuery{}
	}
	return q.nested
}

// func (q *Query) HasChild() *HasChildQuery {
//     if q.hasChild == nil {
//         q.hasChild = &HasChildQuery{}
//     }
//     return q.hasChild
// }
// func (q *Query) HasParent() *HasParentQuery {
//     if q.hasParent == nil {
//         q.hasParent = &HasParentQuery{}
//     }
//     return q.hasParent
// }

func (q *Query) ParentID() *ParentIDQuery {
	if q.parentID == nil {
		q.parentID = &ParentIDQuery{}
	}
	return q.parentID
}

// func (q *Query) DistanceFeature() *DistanceFeatureQuery {
//     if q.distanceFeature == nil {
//         q.distanceFeature = &DistanceFeatureQuery{}
//     }
//     return q.distanceFeature
// }

func (q *Query) MoreLikeThis() *MoreLikeThisQuery {
	if q.moreLikeThis == nil {
		q.moreLikeThis = &MoreLikeThisQuery{}
	}
	return q.moreLikeThis
}

func (q *Query) Percolate() *PercolateQuery {
	if q.percolate == nil {
		q.percolate = &PercolateQuery{}
	}
	return q.percolate
}

// func (q *Query) RankFeature() *RankFeatureQuery {
//     if q.rankFeature == nil {
//         q.rankFeature = &RankFeatureQuery{}
//     }
//     return q.rankFeature
// }
// func (q *Query) Wrapper() *WrapperQuery {
//     if q.wrapper == nil {
//         q.wrapper = &WrapperQuery{}
//     }
//     return q.wrapper
// }
// func (q *Query) Pinned() *PinnedQuery {
//     if q.pinned == nil {
//         q.pinned = &PinnedQuery{}
//     }
//     return q.pinned
// }
// func (q *Query) SpanContaining() *SpanContainingQuery {
//     if q.spanContaining == nil {
//         q.spanContaining = &SpanContainingQuery{}
//     }
//     return q.spanContaining
// }
// func (q *Query) FieldMaskingSpan() *FieldMaskingSpanQuery {
//     if q.fieldMaskingSpan == nil {
//         q.fieldMaskingSpan = &FieldMaskingSpanQuery{}
//     }
//     return q.fieldMaskingSpan
// }
// func (q *Query) SpanFirst() *SpanFirstQuery {
//     if q.spanFirst == nil {
//         q.spanFirst = &SpanFirstQuery{}
//     }
//     return q.spanFirst
// }
// func (q *Query) SpanMulti() *SpanMultiQuery {
//     if q.spanMulti == nil {
//         q.spanMulti = &SpanMultiQuery{}
//     }
//     return q.spanMulti
// }
// func (q *Query) SpanNear() *SpanNearQuery {
//     if q.spanNear == nil {
//         q.spanNear = &SpanNearQuery{}
//     }
//     return q.spanNear
// }
// func (q *Query) SpanNot() *SpanNotQuery {
//     if q.spanNot == nil {
//         q.spanNot = &SpanNotQuery{}
//     }
//     return q.spanNot
// }
// func (q *Query) SpanOr() *SpanOrQuery {
//     if q.spanOr == nil {
//         q.spanOr = &SpanOrQuery{}
//     }
//     return q.spanOr
// }
// func (q *Query) SpanTerm() *SpanTermQuery {
//     if q.spanTerm == nil {
//         q.spanTerm = &SpanTermQuery{}
//     }
//     return q.spanTerm
// }
// func (q *Query) SpanWithin() *SpanWithinQuery {
//     if q.spanWithin == nil {
//         q.spanWithin = &SpanWithinQuery{}
//     }
//     return q.spanWithin
// }

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

		QueryKindPrefix:            q.prefix,
		QueryKindMatch:             q.match,
		QueryKindMatchAll:          q.matchAll,
		QueryKindMatchNone:         q.matchNone,
		QueryKindTerm:              q.term,
		QueryKindExists:            q.exists,
		QueryKindTerms:             q.terms,
		QueryKindRange:             q.rng,
		QueryKindBoosting:          q.boosting,
		QueryKindBoolean:           q.boolean,
		QueryKindConstantScore:     q.constantScore,
		QueryKindFunctionScore:     q.functionScore,
		QueryKindDisjunctionMax:    q.disjunctionMax,
		QueryKindFuzzy:             q.fuzzy,
		QueryKindScriptScore:       q.scriptScore,
		QueryKindScript:            q.script,
		QueryKindIDs:               q.ids,
		QueryKindIntervals:         q.intervals,
		QueryKindMatchBoolPrefix:   q.matchBoolPrefix,
		QueryKindMatchPhrase:       q.matchPhrase,
		QueryKindMatchPhrasePrefix: q.matchPhrasePrefix,
		QueryKindMultiMatch:        q.multiMatch,
		QueryKindQueryString:       q.queryString,
		QueryKindSimpleQueryString: q.simpleQueryString,
		QueryKindGeoBoundingBox:    q.geoBoundingBox,
		QueryKindWildcard:          q.wildcard,
		// QueryKindCommon:            q.common,
		// QueryKindRegexp:            q.regexp,
		// QueryKindTermSet:           q.termSet,
		// QueryKindType:              q.typ,
		// QueryKindWildcard:          q.wildcard,
		// QueryKindAllOf:             q.allOf,

		QueryKindGeoDistance: q.geoDistance,
		// QueryKindGeoPolygon:       q.geoPolygon,
		QueryKindGeoShape: q.geoShape,
		QueryKindShape:    q.shape,
		QueryKindNested:   q.nested,
		// QueryKindHasChild:         q.hasChild,
		// QueryKindHasParent:        q.hasParent,
		QueryKindParentID: q.parentID,
		// QueryKindDistanceFeature:   q.distanceFeature,
		QueryKindMoreLikeThis: q.moreLikeThis,
		QueryKindPercolate:    q.percolate,
		// QueryKindRankFeature:      q.rankFeature,
		// QueryKindWrapper:          q.wrapper,
		// QueryKindPinned:           q.pinned,
		// QueryKindSpanContaining:   q.spanContaining,
		// QueryKindFieldMaskingSpan: q.fieldMaskingSpan,
		// QueryKindSpanFirst:        q.spanFirst,
		// QueryKindSpanMulti:        q.spanMulti,
		// QueryKindSpanNear:         q.spanNear,
		// QueryKindSpanNot:          q.spanNot,
		// QueryKindSpanOr:           q.spanOr,
		// QueryKindSpanTerm:         q.spanTerm,
		// QueryKindSpanWithin:       q.spanWithin,
	}
}

func (q *Query) setClause(qc QueryClause) {
	switch qc.Kind() {

	case QueryKindPrefix:
		q.prefix = qc.(*PrefixQuery)
	case QueryKindMatch:
		q.match = qc.(*MatchQuery)
	case QueryKindMatchAll:
		q.matchAll = qc.(*MatchAllQuery)
	case QueryKindMatchNone:
		q.matchNone = qc.(*MatchNoneQuery)
	case QueryKindTerm:
		q.term = qc.(*TermQuery)
	case QueryKindExists:
		q.exists = qc.(*ExistsQuery)
	case QueryKindTerms:
		q.terms = qc.(*TermsQuery)
	case QueryKindRange:
		q.rng = qc.(*RangeQuery)
	case QueryKindBoosting:
		q.boosting = qc.(*BoostingQuery)
	case QueryKindBoolean:
		q.boolean = qc.(*BoolQuery)
	case QueryKindConstantScore:
		q.constantScore = qc.(*ConstantScoreQuery)
	case QueryKindFunctionScore:
		q.functionScore = qc.(*FunctionScoreQuery)
	case QueryKindDisjunctionMax:
		q.disjunctionMax = qc.(*DisjunctionMaxQuery)
	// case QueryKindAllOf:
	// 	q.allOf = qc.(*AllOfQuery)
	case QueryKindFuzzy:
		q.fuzzy = qc.(*FuzzyQuery)
	case QueryKindScriptScore:
		q.scriptScore = qc.(*ScriptScoreQuery)
	case QueryKindScript:
		q.script = qc.(*ScriptQuery)
	case QueryKindIDs:
		q.ids = qc.(*IDsQuery)
	case QueryKindIntervals:
		q.intervals = qc.(*IntervalsQuery)
	case QueryKindMatchBoolPrefix:
		q.matchBoolPrefix = qc.(*MatchBoolPrefixQuery)
	case QueryKindMatchPhrase:
		q.matchPhrase = qc.(*MatchPhraseQuery)
	case QueryKindMatchPhrasePrefix:
		q.matchPhrasePrefix = qc.(*MatchPhrasePrefixQuery)
	case QueryKindMultiMatch:
		q.multiMatch = qc.(*MultiMatchQuery)
	case QueryKindQueryString:
		q.queryString = qc.(*QueryStringQuery)
	case QueryKindSimpleQueryString:
		q.simpleQueryString = qc.(*SimpleQueryStringQuery)
	case QueryKindGeoBoundingBox:
		q.geoBoundingBox = qc.(*GeoBoundingBoxQuery)
	case QueryKindWildcard:
		q.wildcard = qc.(*WildcardQuery)

	// case QueryKindCommon:
	// 	q.common= qc.(*CommonQuery)
	// case QueryKindRegexp:
	// 	q.regexp= qc.(*RegexpQuery)
	// case QueryKindTermSet:
	// 	q.termSet= qc.(*TermSetQuery)
	// case QueryKindType:
	// 	q.typ= qc.(*TypeQuery)
	case QueryKindGeoDistance:
		q.geoDistance = qc.(*GeoDistanceQuery)
	// case QueryKindGeoPolygon:
	// 	q.geoPolygon = qc.(*GeoPolygonQuery)
	case QueryKindGeoShape:
		q.geoShape = qc.(*GeoShapeQuery)
	case QueryKindShape:
		q.shape = qc.(*ShapeQuery)
	case QueryKindNested:
		q.nested = qc.(*NestedQuery)
	// case QueryKindHasChild:
	// 	q.hasChild = qc.(*HasChildQuery)
	// case QueryKindHasParent:
	// 	q.hasParent = qc.(*HasParentQuery)
	case QueryKindParentID:
		q.parentID = qc.(*ParentIDQuery)
	// case QueryKindDistanceFeature:
	// 	q.distanceFeature = qc.(*DistanceFeatureQuery)
	case QueryKindMoreLikeThis:
		q.moreLikeThis = qc.(*MoreLikeThisQuery)
	case QueryKindPercolate:
		q.percolate = qc.(*PercolateQuery)
		// case QueryKindRankFeature:
		// 	q.rankFeature = qc.(*RankFeatureQuery)
		// case QueryKindWrapper:
		// 	q.wrapper = qc.(*WrapperQuery)
		// case QueryKindPinned:
		// 	q.pinned = qc.(*PinnedQuery)
		// case QueryKindSpanContaining:
		// 	q.spanContaining = qc.(*SpanContainingQuery)
		// case QueryKindFieldMaskingSpan:
		// 	q.fieldMaskingSpan = qc.(*FieldMaskingSpanQuery)
		// case QueryKindSpanFirst:
		// 	q.spanFirst = qc.(*SpanFirstQuery)
		// case QueryKindSpanMulti:
		// 	q.spanMulti = qc.(*SpanMultiQuery)
		// case QueryKindSpanNear:
		// 	q.spanNear = qc.(*SpanNearQuery)
		// case QueryKindSpanNot:
		// 	q.spanNot = qc.(*SpanNotQuery)
		// case QueryKindSpanOr:
		// 	q.spanOr = qc.(*SpanOrQuery)
		// case QueryKindSpanTerm:
		// 	q.spanTerm = qc.(*SpanTermQuery)
		// case QueryKindSpanWithin:
		// 	q.spanWithin = qc.(*SpanWithinQuery)
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
