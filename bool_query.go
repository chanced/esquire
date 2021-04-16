package picker

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

type Booler interface {
	Bool() (*BoolQuery, error)
}

// BoolQueryParams is a query that matches documents matching boolean combinations
// of other queries. The bool query maps to Lucene BoolQueryParams. It is built
// using one or more boolean clauses, each clause with a typed occurrence.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-bool-query.html
type BoolQueryParams struct {
	// The clause (query) must appear in matching documents and will contribute
	// to the score.
	Must Clauses
	// The clause (query) must appear in matching documents. However unlike must
	// the score of the query will be ignored. Filter clauses are executed in
	// filter context, meaning that scoring is ignored and clauses are
	// considered for caching.
	Filter Clauses
	// The clause (query) should appear in the matching document.
	Should Clauses
	// The clause (query) must not appear in the matching documents. Clauses are
	// executed in filter context meaning that scoring is ignored and clauses are
	// considered for caching. Because scoring is ignored, a score of 0 for all
	// documents is returned.
	MustNot Clauses
	// You can use the minimum_should_match parameter to specify the number or
	// percentage of should clauses returned documents must match.
	//
	// If the bool query includes at least one should clause and no must or
	// filter clauses, the default value is 1. Otherwise, the default value is
	// 0.
	MinimumShouldMatch string
	Name               string
	completeClause
}

func (b BoolQueryParams) Clause() (QueryClause, error) {
	return b.Bool()
}

func (b BoolQueryParams) Bool() (*BoolQuery, error) {
	q := &BoolQuery{}
	err := q.SetMust(b.Must)
	if err != nil {
		return q, newQueryError(err, QueryKindBoolean)
	}
	err = q.SetMustNot(b.MustNot)
	if err != nil {
		return q, newQueryError(err, QueryKindBoolean)
	}
	err = q.SetShould(b.Should)
	if err != nil {
		return q, newQueryError(err, QueryKindBoolean)
	}

	err = q.SetFilter(b.Filter)
	if err != nil {
		return q, newQueryError(err, QueryKindBoolean)
	}

	q.SetName(b.Name)
	q.SetMinimumShouldMatch(b.MinimumShouldMatch)
	return q, nil
}

func (b BoolQueryParams) Kind() QueryKind {
	return QueryKindBoolean
}

// BoolQuery is a query that matches documents matching boolean combinations
// of other queries. The bool query maps to Lucene BoolQuery. It is built
// using one or more boolean clauses, each clause with a typed occurrence.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-bool-query.html
type BoolQuery struct {
	must    QueryClauses
	filter  QueryClauses
	should  QueryClauses
	mustNot QueryClauses
	minimumShouldMatchParam
	nameParam
	completeClause
}

func (b *BoolQuery) Clause() (QueryClause, error) {
	return b, nil
}

func (BoolQuery) Kind() QueryKind {
	return QueryKindBoolean
}

func (b *BoolQuery) Set(v Booler) error {
	q, err := v.Bool()
	if err != nil {
		return newQueryError(err, QueryKindBoolean)
	}
	*b = *q
	return nil
}

// Must clauses (query) must appear in matching documents and will contribute
// to the score.
func (b *BoolQuery) Must() *QueryClauses {
	if b == nil {
		return nil
	}
	return &b.must
}

// MustNot is a set of clauses (query) where each clause must not appear in the
// matching documents. Clauses are executed in filter context meaning that
// scoring is ignored and clauses are considered for caching. Because scoring is
// ignored, a score of 0 for all documents is returned.
func (b *BoolQuery) MustNot() *QueryClauses {
	if b == nil {
		return nil
	}
	return &b.mustNot
}

// Filter clauses (query) that must appear in matching documents. However unlike
// must the score of the query will be ignored. Filter clauses are executed in
// filter context, meaning that scoring is ignored and clauses are considered
// for caching.
func (b *BoolQuery) Filter() *QueryClauses {
	if b == nil {
		return nil
	}
	return &b.filter
}

// Should clauses (query) that should appear in the matching document.
func (b *BoolQuery) Should() *QueryClauses {
	if b == nil {
		return nil
	}
	return &b.should
}

func (b *BoolQuery) SetMust(clauses Clauses) error {
	if b == nil {
		*b = BoolQuery{}
	}
	return b.must.Set(clauses)
}

func (b *BoolQuery) SetMustNot(clauses Clauses) error {
	if b == nil {
		*b = BoolQuery{}
	}
	return b.mustNot.Set(clauses)
}

func (b *BoolQuery) SetShould(clauses Clauses) error {
	if b == nil {
		*b = BoolQuery{}
	}
	return b.should.Set(clauses)
}

func (b *BoolQuery) SetFilter(clauses Clauses) error {
	if b == nil {
		*b = BoolQuery{}
	}
	return b.filter.Set(clauses)
}

func (b *BoolQuery) UnmarshalBSON(data []byte) error {
	return b.UnmarshalJSON(data)
}

func (b *BoolQuery) UnmarshalJSON(data []byte) error {
	*b = BoolQuery{}
	obj, err := unmarshalClauseParams(data, b)
	if err != nil {
		return err
	}
	err = b.filter.UnmarshalJSON(obj["filter"])
	if err != nil {
		return err
	}
	err = b.should.UnmarshalJSON(obj["should"])
	if err != nil {
		return err
	}
	err = b.must.UnmarshalJSON(obj["must"])
	if err != nil {
		return err
	}
	err = b.mustNot.UnmarshalJSON(obj["must_not"])
	if err != nil {
		return err
	}
	return nil
}

func (b BoolQuery) MarshalBSON() ([]byte, error) {
	return b.MarshalJSON()
}

func (b BoolQuery) MarshalJSON() ([]byte, error) {
	if b.IsEmpty() {
		return dynamic.Null, nil
	}
	data, err := marshalClauseParams(&b)
	if err != nil {
		return nil, err
	}
	if !b.must.IsEmpty() {
		data["must"] = b.must
	}
	if !b.mustNot.IsEmpty() {
		data["must_not"] = b.mustNot
	}
	if !b.should.IsEmpty() {
		data["should"] = b.should
	}
	if !b.filter.IsEmpty() {
		data["filter"] = b.filter
	}
	return json.Marshal(data)
}

func (b *BoolQuery) IsEmpty() bool {
	return b == nil || !(!b.must.IsEmpty() || !b.mustNot.IsEmpty() || !b.should.IsEmpty() || !b.filter.IsEmpty())
}

func (b *BoolQuery) Clear() {
	*b = BoolQuery{}
}
