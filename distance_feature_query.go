package picker

type DistantFeatureQueryParams struct {
	Name string
	completeClause
}

func (DistantFeatureQueryParams) Kind() QueryKind {
	return QueryKindDistantFeature
}

func (p DistantFeatureQueryParams) Clause() (QueryClause, error) {
	return p.DistantFeature()
}
func (p DistantFeatureQueryParams) DistantFeature() (*DistantFeatureQuery, error) {
	q := &DistantFeatureQuery{}
	_ = q
	panic("not implemented")
	// return q, nil
}

type DistantFeatureQuery struct {
	nameParam
	completeClause
}

func (DistantFeatureQuery) Kind() QueryKind {
	return QueryKindDistantFeature
}
func (q *DistantFeatureQuery) Clause() (QueryClause, error) {
	return q, nil
}
func (q *DistantFeatureQuery) DistantFeature() (QueryClause, error) {
	return q, nil
}
func (q *DistantFeatureQuery) Clear() {
	if q == nil {
		return
	}
	*q = DistantFeatureQuery{}
}
func (q *DistantFeatureQuery) UnmarshalJSON(data []byte) error {
	panic("not implemented")
}
func (q DistantFeatureQuery) MarshalJSON() ([]byte, error) {
	panic("not implemented")
}
func (q *DistantFeatureQuery) IsEmpty() bool {
	panic("not implemented")
}

type distantFeatureQuery struct {
	Name string `json:"_name,omitempty"`
}
