package search

// Boolean is a query that matches documents matching boolean combinations
// of other queries. The bool query maps to Lucene BooleanQuery. It is built
// using one or more boolean clauses, each clause with a typed occurrence.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-bool-query.html
type Boolean struct {
	//The clause (query) must appear in matching documents and will contribute
	//to the score.
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
	QueryName          string
}

func (b Boolean) Clause() (Clause, error) {
	return b.boolean()
}

func (b Boolean) boolean() (*booleanClause, error) {
	c := &booleanClause{
		Must:    b.Must,
		Filter:  b.Filter,
		Should:  b.Should,
		MustNot: b.MustNot,
	}
	c.SetName(b.Name())
	c.SetMinimumShouldMatch(b.MinimumShouldMatch)
	return c, nil
}

func (b Boolean) Name() string {
	return b.QueryName
}
func (b Boolean) Type() Type {
	return TypeBoolean
}

type booleanClause struct {
	//The clause (query) must appear in matching documents and will contribute
	//to the score.
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
	minimumShouldMatchParam
	nameParam
}

func (b booleanClause) Type() Type {
	return TypeBoolean
}

// BooleanQuery is a query that matches documents matching boolean combinations
// of other queries. The bool query maps to Lucene BooleanQuery. It is built
// using one or more boolean clauses, each clause with a typed occurrence.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-bool-query.html
type BooleanQuery struct {
	booleanClause
}

func (b *BooleanQuery) HasBooleanClause() bool {
	return len(b.Must) > 0 || len(b.MustNot) > 0 || len(b.Should) > 0 || len(b.Filter) > 0
}

func (b *BooleanQuery) SetBoolean(v *Boolean) error {
	c, err := v.boolean()
	if err != nil {
		return err
	}
	b.booleanClause = *c
	return nil
}

func (b *BooleanQuery) AddMustClause(c Clause) (err error) {
	return b.Must.Add(c)
}
