package mapping

// WithPositiveScoreImpact is a mapping with the positive_score_impact parameter
// Rank features that correlate negatively with the score should set
// positive_score_impact to false (defaults to true). This will be used by the
// rank_feature query to modify the scoring formula in such a way that the score
// decreases with the value of the feature instead of increasing. For instance
// in web search, the url length is a commonly used feature which correlates
// negatively with scores.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/rank-feature.html
type WithPositiveScoreImpact interface {
	// PositiveScoreImpact is used by rank_feature queries to modify the scoring
	// formula in such a way that the score increases or decreases the value of
	// the feature
	PositiveScoreImpact() bool
	// SetPositiveScoreImpact sets the PositiveScoreImpact Value to v
	SetPositiveScoreImpact(v bool)
}

// FieldWithPositiveScoreImpact is a Field with the
// positive_score_impact param
type FieldWithPositiveScoreImpact interface {
	Field
	WithPositiveScoreImpact
}

// PositiveScoreImpactParam is a mixin that adds the
// positive_score_impact paramete
type PositiveScoreImpactParam struct {
	PositiveScoreValue *bool `bson:"positive_score_impact,omitempty" json:"positive_score_impact,omitempty"`
}

// PositiveScoreImpact is used by rank_feature queries to modify the scoring
// formula in such a way that the score increases or decreases the value of
// the feature
func (psi PositiveScoreImpactParam) PositiveScoreImpact() bool {
	if psi.PositiveScoreValue == nil {
		return true
	}
	return *psi.PositiveScoreValue
}

// SetPositiveScoreImpact sets the PositiveScoreImpact Value to v
func (psi *PositiveScoreImpactParam) SetPositiveScoreImpact(v bool) {
	psi.PositiveScoreValue = &v
}
