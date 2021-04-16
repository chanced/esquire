package picker

type DistanceFeatureQueryParams struct {
	Name string
	completeClause
}

func (DistanceFeatureQueryParams) Kind() QueryKind {
	return QueryKindDistanceFeature
}

func (p DistanceFeatureQueryParams) Clause() (QueryClause, error) {
	return p.DistanceFeature()
}
func (p DistanceFeatureQueryParams) DistanceFeature() (*DistanceFeatureQuery, error) {
	q := &DistanceFeatureQuery{}
	_ = q
	panic("not implemented")
	// return q, nil
}

type DistanceFeatureQuery struct {
	nameParam
	completeClause
}

func (DistanceFeatureQuery) Kind() QueryKind {
	return QueryKindDistanceFeature
}
func (q *DistanceFeatureQuery) Clause() (QueryClause, error) {
	return q, nil
}
func (q *DistanceFeatureQuery) DistanceFeature() (*DistanceFeatureQuery, error) {
	return q, nil
}
func (q *DistanceFeatureQuery) Clear() {
	if q == nil {
		return
	}
	*q = DistanceFeatureQuery{}
}
func (q *DistanceFeatureQuery) UnmarshalBSON(data []byte) error {
	return q.UnmarshalJSON(data)
}

func (q *DistanceFeatureQuery) UnmarshalJSON(data []byte) error {
	panic("not implemented")
}
func (q DistanceFeatureQuery) MarshalBSON() ([]byte, error) {
	return q.MarshalJSON()
}

func (q DistanceFeatureQuery) MarshalJSON() ([]byte, error) {
	panic("not implemented")
}
func (q *DistanceFeatureQuery) IsEmpty() bool {
	panic("not implemented")
}

type distanceFeatureQuery struct {
	Name string `json:"_name,omitempty"`
}
