package picker

type GeoShapeQueryParams struct {
	Name string
	completeClause
}

func (GeoShapeQueryParams) Kind() QueryKind {
	return QueryKindGeoShape
}

func (p GeoShapeQueryParams) Clause() (QueryClause, error) {
	return p.GeoShape()
}
func (p GeoShapeQueryParams) GeoShape() (*GeoShapeQuery, error) {
	q := &GeoShapeQuery{}
	_ = q
	panic("not implemented")
	// return q, nil
}

type GeoShapeQuery struct {
	nameParam
	completeClause
}

func (GeoShapeQuery) Kind() QueryKind {
	return QueryKindGeoShape
}
func (q *GeoShapeQuery) Clause() (QueryClause, error) {
	return q, nil
}
func (q *GeoShapeQuery) GeoShape() (QueryClause, error) {
	return q, nil
}
func (q *GeoShapeQuery) Clear() {
	if q == nil {
		return
	}
	*q = GeoShapeQuery{}
}
func (q *GeoShapeQuery) UnmarshalJSON(data []byte) error {
	panic("not implemented")
}
func (q GeoShapeQuery) MarshalJSON() ([]byte, error) {
	panic("not implemented")
}
func (q *GeoShapeQuery) IsEmpty() bool {
	panic("not implemented")
}

type geoShapeQuery struct {
	Name string `json:"_name,omitempty"`
}
