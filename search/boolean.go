package search

// Boolean is a query that matches documents matching boolean combinations
// of other queries. The bool query maps to Lucene BooleanQuery. It is built
// using one or more boolean clauses, each clause with a typed occurrence.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-bool-query.html

type Boolean struct {
}

type BooleanRule struct {
}

func (b BooleanRule) Type() Type {
	return TypeBoolean
}

// BooleanQuery is a query that matches documents matching boolean combinations
// of other queries. The bool query maps to Lucene BooleanQuery. It is built
// using one or more boolean clauses, each clause with a typed occurrence.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-bool-query.html
type BooleanQuery struct {
}
