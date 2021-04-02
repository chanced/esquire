package picker

// A RankFeatureField can index numbers so that they can later be used to boost documents in queries with a rank_feature query.
//
// Rank features that correlate negatively with the score should set positive_score_impact to false (defaults to true). This will be used by the rank_feature query to modify the scoring formula in such a way that the score decreases with the value of the feature instead of increasing. For instance in web search, the url length is a commonly used feature which correlates negatively with scores.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/rank-feature.html
type RankFeatureField struct {
	BaseField                `json:",inline" bson:",inline"`
	PositiveScoreImpactParam `json:",inline" bson:",inline"`
}

func (f RankFeatureField) Clone() Field {
	n := NewRankFeatureField()
	n.SetPositiveScoreImpact(n.PositiveScoreImpact())
	return n
}

func NewRankFeatureField() *RankFeatureField {
	return &RankFeatureField{BaseField: BaseField{MappingType: FieldTypeRankFeature}}
}

// A RankFeaturesField can index numeric feature vectors, so that they can later be used to boost documents in queries with a rank_feature query.
//
// It is analogous to the RankFeature data type but is better suited when the list of features is sparse so that it wouldn’t be reasonable to add one field to the mappings for each of them.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/rank-features.html
type RankFeaturesField struct {
	BaseField                `json:",inline" bson:",inline"`
	PositiveScoreImpactParam `json:",inline" bson:",inline"`
}

func (f RankFeaturesField) Clone() Field {
	n := NewRankFeaturesField()
	n.SetPositiveScoreImpact(n.PositiveScoreImpact())
	return n
}

func NewRankFeaturesField() *RankFeaturesField {
	return &RankFeaturesField{BaseField: BaseField{MappingType: FieldTypeRankFeatures}}
}
