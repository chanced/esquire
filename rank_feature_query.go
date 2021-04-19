package picker

type RankFeatureQueryParams struct {
	// (Required) rank_feature or rank_features field used to boost
	// relevance scores.
	Field string
	// (Optional, float) Floating point number used to decrease or increase
	// relevance scores. Defaults to 1.0.
	//
	// Boost values are relative to the default value of 1.0. A boost value
	// between 0 and 1.0 decreases the relevance score. A value greater than 1.0
	// increases the relevance score.
	Boost interface{}
	// (Optional, function object) Saturation function used to boost relevance
	// scores based on the value of the rank feature field. If no function is
	// provided, the rank_feature query defaults to the saturation function. See
	// Saturation for more information.
	//
	// Only one function saturation, log, sigmoid or linear can be provided.
	Saturation SaturationFunction
	Name       string
	completeClause
}

func (RankFeatureQueryParams) Kind() QueryKind {
	return QueryKindRankFeature
}

func (p RankFeatureQueryParams) Clause() (QueryClause, error) {
	return p.RankFeature()
}
func (p RankFeatureQueryParams) RankFeature() (*RankFeatureQuery, error) {
	q := &RankFeatureQuery{}
	_ = q
	panic("not implemented")
	// return q, nil
}

type RankFeatureQuery struct {
	nameParam
	completeClause
}

func (RankFeatureQuery) Kind() QueryKind {
	return QueryKindRankFeature
}
func (q *RankFeatureQuery) Clause() (QueryClause, error) {
	return q, nil
}
func (q *RankFeatureQuery) RankFeature() (*RankFeatureQuery, error) {
	return q, nil
}
func (q *RankFeatureQuery) Clear() {
	if q == nil {
		return
	}
	*q = RankFeatureQuery{}
}
func (q *RankFeatureQuery) UnmarshalBSON(data []byte) error {
	return q.UnmarshalJSON(data)
}

func (q *RankFeatureQuery) UnmarshalJSON(data []byte) error {
	panic("not implemented")
}
func (q RankFeatureQuery) MarshalBSON() ([]byte, error) {
	return q.MarshalJSON()
}

func (q RankFeatureQuery) MarshalJSON() ([]byte, error) {
	panic("not implemented")
}
func (q *RankFeatureQuery) IsEmpty() bool {
	panic("not implemented")
}

type rankFeatureQuery struct {
	Name string `json:"_name,omitempty"`
}

// SaturationFunction gives a score equal to S / (S + pivot), where S is the
// value of the rank feature field and pivot is a configurable pivot value so
// that the result will be less than 0.5 if S is less than pivot and greater
// than 0.5 otherwise. Scores are always (0,1).
//
// If the rank feature has a negative score impact then the function will be
// computed as pivot / (S + pivot), which decreases when S increases.
//
// If a pivot value is not provided, Elasticsearch computes a default value
// equal to the approximate geometric mean of all rank feature values in the
// index.
//easyjson:json
type SaturationFunction struct {
	pivot interface{} `json:"pivot"`
}

// LogFunction is a logarithmic function used to boost relevance scores based on
// the value of the rank feature field.
type LogFunction struct {
}
