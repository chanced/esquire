package search

import (
	"github.com/chanced/dynamic"
)

const DefaultRewrite = RewriteConstantScore

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

var Rewrites = []Rewrite{
	RewriteConstantScore, RewriteConstantScoreBoolean, RewriteConstantScoreBoolean,
	RewriteScoringBoolean, RewriteTopTermsBlendedFreqsN, RewriteTopTermsBoostN,
	RewriteTopTermsN,
}

func (r Rewrite) IsValid() bool {
	if r == "" {
		return true
	}
	for _, v := range Rewrites {
		if v == r {
			return true
		}
	}
	return false
}

// WithRewrite is a query with the rewrite param
type WithRewrite interface {
	Rewrite() Rewrite
	SetRewrite(v Rewrite)
}

// rewriteParam is a mixin that adds the rewrite param
//
// Method used to rewrite the query. For valid values and more information, see
// the rewrite parameter. (Optional)
type rewriteParam struct {
	rewrite *Rewrite
}

func (r rewriteParam) Rewrite() Rewrite {
	if r.rewrite == nil {
		return DefaultRewrite
	}
	return *r.rewrite
}
func (r *rewriteParam) SetRewrite(v Rewrite) {
	if v != "" && v != r.Rewrite() {
		r.rewrite = &v
	}
}
func unmarshalRewriteParam(data dynamic.RawJSON, target interface{}) error {
	if a, ok := target.(WithRewrite); ok {
		a.SetRewrite(Rewrite(data.UnquotedString()))
	}
	return nil
}
func marshalRewriteParam(data M, source interface{}) (M, error) {
	if b, ok := source.(WithRewrite); ok {
		if b.Rewrite() != DefaultRewrite {
			data[paramRewrite] = b.Rewrite()
		}
	}
	return data, nil
}
