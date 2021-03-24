package search

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
	MatchQuery   `json:",inline" bson:",inline"`
	ScriptQuery  `json:",inline" bson:",inline"`
	ExistsQuery  `json:",inline" bson:",inline"`
	BooleanQuery `json:",inline" bson:",inline"`
	TermQuery    `json:",inline" bson:",inline"`
}

func NewQuery() Query {
	return Query{
		MatchQuery: newMatchQuery(),
	}
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
