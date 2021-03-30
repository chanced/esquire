package search

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

type Booler interface {
	Boolean() (BooleanQuery, error)
}

// Boolean is a query that matches documents matching boolean combinations
// of other queries. The bool query maps to Lucene BooleanQuery. It is built
// using one or more boolean clauses, each clause with a typed occurrence.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-bool-query.html
type Boolean struct {
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
	clause
}

func (b Boolean) Clause() (Clause, error) {
	return b.Boolean()
}

func (b Boolean) Boolean() (BooleanQuery, error) {
	q := BooleanQuery{}
	err := q.SetMust(b.Must)
	if err != nil {
		return q, NewQueryError(err, KindBoolean)
	}
	err = q.SetMustNot(b.MustNot)
	if err != nil {
		return q, NewQueryError(err, KindBoolean)
	}
	err = q.SetShould(b.Should)
	if err != nil {
		return q, NewQueryError(err, KindBoolean)
	}

	err = q.SetFilter(b.Filter)
	if err != nil {
		return q, NewQueryError(err, KindBoolean)
	}

	q.SetName(b.Name)
	q.SetMinimumShouldMatch(b.MinimumShouldMatch)
	return q, nil
}

func (b Boolean) Kind() Kind {
	return KindBoolean
}

// BooleanQuery is a query that matches documents matching boolean combinations
// of other queries. The bool query maps to Lucene BooleanQuery. It is built
// using one or more boolean clauses, each clause with a typed occurrence.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-bool-query.html
type BooleanQuery struct {
	must    QueryClauses
	filter  QueryClauses
	should  QueryClauses
	mustNot QueryClauses
	minimumShouldMatchParam
	nameParam
	clause
}

func (b BooleanQuery) Kind() Kind {
	return KindBoolean
}

func (b *BooleanQuery) Set(v Booler) error {
	q, err := v.Boolean()
	if err != nil {
		return NewQueryError(err, KindBoolean)
	}
	*b = q
	return nil
}

// Must clauses (query) must appear in matching documents and will contribute
// to the score.
func (b *BooleanQuery) Must() *QueryClauses {
	return &b.must
}

// MustNot is a set of clauses (query) where each clause must not appear in the
// matching documents. Clauses are executed in filter context meaning that
// scoring is ignored and clauses are considered for caching. Because scoring is
// ignored, a score of 0 for all documents is returned.
func (b *BooleanQuery) MustNot() *QueryClauses {
	return &b.mustNot
}

// Filter clauses (query) that must appear in matching documents. However unlike
// must the score of the query will be ignored. Filter clauses are executed in
// filter context, meaning that scoring is ignored and clauses are considered
// for caching.
func (b *BooleanQuery) Filter() *QueryClauses {
	return &b.filter
}

// Should clauses (query) that should appear in the matching document.
func (b *BooleanQuery) Should() *QueryClauses {
	return &b.should
}

func (b *BooleanQuery) SetMust(clauses Clauses) error {
	return b.must.Set(clauses)
}

func (b *BooleanQuery) SetMustNot(clauses Clauses) error {
	return b.mustNot.Set(clauses)
}

func (b *BooleanQuery) SetShould(clauses Clauses) error {
	return b.should.Set(clauses)
}

func (b *BooleanQuery) SetFilter(clauses Clauses) error {
	return b.filter.Set(clauses)
}

func (b *BooleanQuery) UnmarshalJSON(data []byte) error {
	*b = BooleanQuery{}
	obj, err := unmarshalParams(data, b)
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

func (b BooleanQuery) MarshalJSON() ([]byte, error) {
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

func (b *BooleanQuery) IsEmpty() bool {
	return b == nil || !(!b.must.IsEmpty() || !b.mustNot.IsEmpty() || !b.should.IsEmpty() || !b.filter.IsEmpty())
}

func (b *BooleanQuery) Clear() {
	*b = BooleanQuery{}
}
