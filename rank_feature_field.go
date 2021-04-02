package picker

import "encoding/json"

type RankFeatureFieldParams struct {
	// PositiveScoreImpact is used by rank_feature queries to modify the scoring
	// formula in such a way that the score increases or decreases the value of
	// the feature
	//
	// (Bool, optional)
	PositiveScoreImpact interface{} `json:"positive_score_impact,omitempty"`
}

func (RankFeatureFieldParams) Type() FieldType {
	return FieldTypeRankFeature
}
func (p RankFeatureFieldParams) Field() (Field, error) {
	return p.RankFeature()
}
func (p RankFeatureFieldParams) RankFeature() (*RankFeatureField, error) {
	f := &RankFeatureField{}
	err := f.SetPositiveScoreImpact(p.PositiveScoreImpact)
	return f, err
}

func NewRankFeatureField(params RankFeatureFieldParams) (*RankFeatureField, error) {
	return params.RankFeature()
}

// A RankFeatureField can index numbers so that they can later be used to boost documents in queries with a rank_feature query.
//
// Rank features that correlate negatively with the score should set positive_score_impact to false (defaults to true). This will be used by the rank_feature query to modify the scoring formula in such a way that the score decreases with the value of the feature instead of increasing. For instance in web search, the url length is a commonly used feature which correlates negatively with scores.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/rank-feature.html
type RankFeatureField struct {
	positiveScoreImpactParam `json:",inline" bson:",inline"`
}

func (RankFeatureField) Type() FieldType {
	return FieldTypeRankFeature
}
func (f RankFeatureField) MarshalJSON() ([]byte, error) {
	return json.Marshal(RankFeatureFieldParams{
		PositiveScoreImpact: f.positiveScoreImpact.Value(),
	})
}
func (f *RankFeatureField) UnmarshalJSON(data []byte) error {
	var p RankFeatureFieldParams

	err := json.Unmarshal(data, &p)
	if err != nil {
		return err
	}
	n, err := p.RankFeature()
	if err != nil {
		return err
	}
	*f = *n
	return nil
}

type RankFeaturesFieldParams struct {
	// PositiveScoreImpact is used by rank_feature queries to modify the scoring
	// formula in such a way that the score increases or decreases the value of
	// the feature
	//
	// (Bool, optional)
	PositiveScoreImpact interface{} `json:"positive_score_impact,omitempty"`
}

func (RankFeaturesFieldParams) Type() FieldType {
	return FieldTypeRankFeatures
}
func (p RankFeaturesFieldParams) Field() (Field, error) {
	return p.RankFeatures()
}
func (p RankFeaturesFieldParams) RankFeatures() (*RankFeaturesField, error) {
	f := &RankFeaturesField{}
	err := f.SetPositiveScoreImpact(p.PositiveScoreImpact)
	return f, err
}

func NewRankFeaturesField(params RankFeaturesFieldParams) (*RankFeaturesField, error) {
	return params.RankFeatures()
}

// A RankFeaturesField can index numeric feature vectors, so that they can later be used to boost documents in queries with a rank_feature query.
//
// It is analogous to the RankFeature data type but is better suited when the list of features is sparse so that it wouldn’t be reasonable to add one field to the mappings for each of them.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/rank-features.html
type RankFeaturesField struct {
	positiveScoreImpactParam `json:",inline" bson:",inline"`
}

func (RankFeaturesField) Type() FieldType {
	return FieldTypeRankFeatures
}
func (f RankFeaturesField) MarshalJSON() ([]byte, error) {
	return json.Marshal(RankFeaturesFieldParams{
		PositiveScoreImpact: f.positiveScoreImpact.Value(),
	})
}
func (f *RankFeaturesField) UnmarshalJSON(data []byte) error {
	var p RankFeaturesFieldParams

	err := json.Unmarshal(data, &p)
	if err != nil {
		return err
	}
	n, err := p.RankFeatures()
	if err != nil {
		return err
	}
	*f = *n
	return nil
}
