package search

import (
	"encoding/json"

	"github.com/chanced/picker/internal/jsonutil"
)

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

	return nil
}

func (q Query) MarshalJSON() ([]byte, error) {
	m := map[string]json.RawMessage{}
	terms, err := q.TermsQuery.MarshalJSON()
	if err != nil {
		return nil, err
	}
	if jsonutil.IsNotNil(terms) {
		m["terms"] = json.RawMessage(terms)
	}

	term, err := q.TermQuery.MarshalJSON()
	if err != nil {
		return nil, err
	}
	if jsonutil.IsNotNil(term) {
		m["term"] = json.RawMessage(term)
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

type QueryParam struct {
	QueryValue *Query `json:"query,omitempty" bson:"query,omitempty"`
}
