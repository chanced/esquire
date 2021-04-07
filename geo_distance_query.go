package picker

type GeoDistanceQueryParams struct {
	Name string
	completeClause
}

func (GeoDistanceQueryParams) Kind() QueryKind {
	return QueryKindGeoDistance
}

func (p GeoDistanceQueryParams) Clause() (QueryClause, error) {
	return p.GeoDistance()
}
func (p GeoDistanceQueryParams) GeoDistance() (*GeoDistanceQuery, error) {
	q := &GeoDistanceQuery{}
	_ = q
	panic("not implemented")
	// return q, nil
}

type GeoDistanceQuery struct {
	nameParam
	completeClause
}

func (GeoDistanceQuery) Kind() QueryKind {
	return QueryKindGeoDistance
}
func (q *GeoDistanceQuery) Clause() (QueryClause, error) {
	return q, nil
}
func (q *GeoDistanceQuery) GeoDistance() (QueryClause, error) {
	return q, nil
}
func (q *GeoDistanceQuery) Clear() {
	if q == nil {
		return
	}
	*q = GeoDistanceQuery{}
}
func (q *GeoDistanceQuery) UnmarshalJSON(data []byte) error {
	panic("not implemented")
}
func (q GeoDistanceQuery) MarshalJSON() ([]byte, error) {
	panic("not implemented")
}
func (q *GeoDistanceQuery) IsEmpty() bool {
	panic("not implemented")
}

type geoDistanceQuery struct {
	Name string `json:"_name,omitempty"`
}
