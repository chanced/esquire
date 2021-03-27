package search

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

type QueryClauses struct {

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
	Term *Term

	// Terms returns documents that contain one or more exact terms in a provided
	// field.
	//
	// The terms query is the same as the term query, except you can search for
	// multiple values.
	Terms *Terms

	// Match returns documents that match a provided text, number, date or boolean
	// value. The provided text is analyzed before matching.
	//
	// The match query is the standard query for performing a full-text search,
	// including options for fuzzy matching.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-query.html
	Match *Match

	// Boolean is a query that matches documents matching boolean combinations
	// of other queries. The bool query maps to Lucene BooleanQuery. It is built
	// using one or more boolean clauses, each clause with a typed occurrence.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-bool-query.html
	Boolean *Boolean
}

func NewQuery(params QueryClauses) (*Query, error) {
	panic("not impl")
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
	MatchQuery
	ScriptQuery
	ExistsQuery
	BooleanQuery
	TermQuery
	TermsQuery
}

func (q Query) HasClauses() bool {
	return q.HasMatchClause() || q.HasTermsClause() || q.HasTermClause() ||
		q.HasBooleanClause()
}

func (q Query) IsEmpty() bool {
	return !q.HasClauses()
}

func (q *Query) UnmarshalJSON(data []byte) error {
	m := map[string]json.RawMessage{}
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	if term, ok := m["term"]; ok {
		err = json.Unmarshal(term, &q.TermQuery)
		if err != nil {
			return err
		}
	}

	if terms, ok := m["terms"]; ok {
		err = json.Unmarshal(terms, &q.TermsQuery)
		if err != nil {
			return err
		}
	}
	if match, ok := m["match"]; ok {
		err = json.Unmarshal(match, &q.MatchQuery)
		if err != nil {
			return err
		}
	}
	return nil
}

func (q Query) MarshalJSON() ([]byte, error) {
	m := map[string]json.RawMessage{}
	terms, err := q.TermsQuery.MarshalJSON()
	if err != nil {
		return nil, err
	}
	if !dynamic.RawJSON(terms).IsNull() {
		m["terms"] = json.RawMessage(terms)
	}

	term, err := q.TermQuery.MarshalJSON()
	if err != nil {
		return nil, err
	}
	if !dynamic.RawJSON(term).IsNull() {
		m["term"] = json.RawMessage(term)
	}

	match, err := q.MatchQuery.MarshalJSON()
	if err != nil {
		return nil, err
	}
	if !dynamic.RawJSON(match).IsNull() {
		m["match"] = json.RawMessage(match)
	}
	return json.Marshal(m)
}

func (q *Query) Clone() *Query {
	if q == nil {
		return nil
	}
	// TODO: implement this
	return &Query{}
}
