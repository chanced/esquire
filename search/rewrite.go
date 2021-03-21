package search

// Rewrite as defined by:
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-multi-term-rewrite.html
type Rewrite string

const (
	// RewriteTypeConstantScore uses the constant_score_boolean method for fewer matching
	// terms. Otherwise, this method finds all matching terms in sequence and
	// returns matching documents using a bit set.
	RewriteConstantScore Rewrite = "constant_score"

	// RewriteTypeConstantScoreBoolean assigns each document a relevance score equal to the
	// boost parameter.
	//
	// This method changes the original query to a bool query. This bool query
	// contains a should clause and term query for each matching term.
	//
	// This method can cause the final bool query to exceed the clause limit in the
	// indices.query.bool.max_clause_count setting. If the query exceeds this limit,
	// Elasticsearch returns an error.
	RewriteConstantScoreBoolean Rewrite = "constant_score_boolean"

	// RewriteTypeScoringBoolean calculates a relevance score for each matching document.
	//
	// This method changes the original query to a bool query. This bool query
	// contains a should clause and term query for each matching term.
	//
	// This method can cause the final bool query to exceed the clause limit in the
	// indices.query.bool.max_clause_count setting. If the query exceeds this limit,
	// Elasticsearch returns an error.
	RewriteScoringBoolean Rewrite = "scoring_boolean"

	// RewriteTypeTopTermsBlendedFreqsN calculates a relevance score for each
	// matching document as if all terms had the same frequency. This frequency
	// is the maximum frequency of all matching terms.
	//
	// This method changes the original query to a bool query. This bool query
	// contains a should clause and term query for each matching term.
	//
	// The final bool query only includes term queries for the top N scoring
	// terms.
	//
	// You can use this method to avoid exceeding the clause limit in the
	// indices.query.bool.max_clause_count setting.
	//
	RewriteTopTermsBlendedFreqsN Rewrite = "top_terms_blended_freqs_N"

	// RewriteTypeTopTermsBoostN Assigns each matching document a relevance
	// score equal to the boost parameter.
	//
	// This method changes the original query to a bool query. This bool query
	// contains a should clause and term query for each matching term.
	//
	// The final bool query only includes term queries for the top N terms.
	//
	// You can use this method to avoid exceeding the clause limit in the
	// indices.query.bool.max_clause_count setting.
	RewriteTopTermsBoostN Rewrite = "top_terms_boost_N"

	// RewriteTopTermsN calculates a relevance score for each matching document.
	//
	// This method changes the original query to a bool query. This bool query
	// contains a should clause and term query for each matching term.
	//
	// The final bool query only includes term queries for the top N scoring terms.
	//
	// You can use this method to avoid exceeding the clause limit in the
	// indices.query.bool.max_clause_count setting.
	RewriteTopTermsN Rewrite = "top_terms_N"
)

// WithRewrite is a query with the rewrite param
type WithRewrite interface {
	Rewrite() Rewrite
	SetRewrite(v Rewrite)
}

// RewriteParam is a mixin that adds the rewrite param
//
// Method used to rewrite the query. For valid values and more information, see
// the rewrite parameter. (Optional)
type RewriteParam struct {
	RewriteValue *Rewrite `json:"rewrite,omitempty" bson:"rewrite,omitempty"`
}

func (r RewriteParam) Default() Rewrite {
	return RewriteConstantScore
}

func (r RewriteParam) Rewrite() Rewrite {
	if r.RewriteValue == nil {
		r.Default()
	}
	return *r.RewriteValue
}
func (r *RewriteParam) SetRewrite(v Rewrite) {
	if v != "" && v != r.Rewrite() {
		r.RewriteValue = &v
	}
}
