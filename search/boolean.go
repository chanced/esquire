package search

type Booler interface {
	Boolean() (BooleanQuery, error)
}

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
	Name               string
}

func (b Boolean) Clause() (Clause, error) {
	return b.Boolean()
}

func (b Boolean) Boolean() (BooleanQuery, error) {
	q := BooleanQuery{}
	err := q.SetMust(b.Must)
	if err != nil {
		return q, err
	}
	err = q.SetMustNot(b.MustNot)
	if err != nil {
		return q, err
	}
	err = q.SetShould(b.Should)
	if err != nil {
		return q, err
	}

	err = q.SetFilter(b.Filter)
	if err != nil {
		return q, err
	}

	q.SetName(b.Name)
	q.SetMinimumShouldMatch(b.MinimumShouldMatch)
	return q, nil
}

func (b Boolean) Type() Type {
	return TypeBoolean
}

// BooleanQuery is a query that matches documents matching boolean combinations
// of other queries. The bool query maps to Lucene BooleanQuery. It is built
// using one or more boolean clauses, each clause with a typed occurrence.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-bool-query.html
type BooleanQuery struct {
	must    Clauses
	filter  Clauses
	should  Clauses
	mustNot Clauses
	minimumShouldMatchParam
	nameParam
}

func (b BooleanQuery) Type() Type {
	return TypeBoolean
}

func (b *BooleanQuery) Set(v Booler) error {
	*b = BooleanQuery{}
	q, err := v.Boolean()
	if err != nil {
		return err
	}
	*b = q
	return nil
}

func (b *BooleanQuery) SetMust(clauses Clauses) error {
	must, err := unpackClauses(clauses)
	if err != nil {
		return err
	}
	b.must = must
	return nil
}

func (b *BooleanQuery) SetMustNot(clauses Clauses) error {
	mustNot, err := unpackClauses(clauses)
	if err != nil {
		return err
	}
	b.mustNot = mustNot
	return nil
}

func (b *BooleanQuery) SetShould(clauses Clauses) error {
	should, err := unpackClauses(clauses)
	if err != nil {
		return err
	}
	b.should = should
	return nil

}

func (b *BooleanQuery) SetFilter(clauses Clauses) error {
	filter, err := unpackClauses(clauses)
	if err != nil {
		return err
	}
	b.filter = filter
	return nil

}

func (b *BooleanQuery) AddMust(c Clause) (err error) {
	return b.must.Add(c)
}

func (b *BooleanQuery) AddShould(c Clause) (err error) {
	return b.should.Add(c)
}

func (b *BooleanQuery) AddMustNot(c Clause) (err error) {
	return b.mustNot.Add(c)
}

func (b *BooleanQuery) AddFilter(c Clause) (err error) {
	return b.filter.Add(c)
}

func (b *BooleanQuery) IsEmpty() bool {
	return !(len(b.must) > 0 || len(b.mustNot) > 0 || len(b.should) > 0 || len(b.filter) > 0)
}
