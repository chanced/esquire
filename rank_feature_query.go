package picker

type RankFeatureQueryParams struct {
	Name string
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
